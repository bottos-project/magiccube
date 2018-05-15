package main

import (
	log "github.com/cihub/seelog"
	"github.com/micro/go-micro"
	user_proto "github.com/bottos-project/bottos/service/user/proto"
	"golang.org/x/net/context"
	"github.com/bottos-project/bottos/tools/db/mongodb"
	"gopkg.in/mgo.v2/bson"
	"github.com/bottos-project/bottos/config"
	"github.com/bottos-project/bottos/service/common/data"
	sign "github.com/bottos-project/bottos/service/common/signature/proto"
	"encoding/hex"
	"github.com/bottos-project/bottos/crypto"
	"github.com/protobuf/proto"
	"github.com/bottos-project/bottos/service/common/util"
	pack "github.com/bottos-project/bottos/core/contract/msgpack"
	"github.com/bottos-project/bottos/service/common/bean"
)
type User struct{}

func (u *User) GetBlockHeader(ctx context.Context, req *user_proto.GetBlockHeaderRequest, rsp *user_proto.GetBlockHeaderResponse) error {
	block_header, err:= data.BlockHeader()
	if block_header != nil {
		rsp.Data = block_header
	} else {
		rsp.Code = 1003
		rsp.Msg = err.Error()
	}
	return nil
}

func (u *User) Register(ctx context.Context, req *user_proto.RegisterRequest, rsp *user_proto.RegisterResponse) error {
	log.Info("req:", req);
	block_header, err:= data.BlockHeader()
	if err != nil {
		rsp.Code = 1003
		rsp.Msg = err.Error()
		return nil
	}
	//注册账号
	rsp.Code = 1004
	account_buf,err := pack.Marshal(req.Account)
	if err != nil {
		rsp.Msg = err.Error()
		return nil
	}
	tx_account_sign := &sign.BasicTransaction{
		Version:1,
		CursorNum: block_header.HeadBlockNum,
		CursorLabel: block_header.CursorLabel,
		Lifetime: block_header.HeadBlockTime + 20,
		Sender: "bottos",
		Contract: "bottos",
		Method: "newaccount",
		Param: account_buf,
		SigAlg: 1,
	}


	msg, err := proto.Marshal(tx_account_sign)
	if err != nil {
		rsp.Msg = err.Error()
		return nil
	}
	//配对的pubkey   0401787e34de40f3aeb4c28259637e8c9e84b5a58f57b3c23f010f4dc7230dffced4976238196bd32cd90569d66f747525b194ca83146965df092d2585b975d0d3
	seckey, err := hex.DecodeString("81407d25285450184d29247b5f06408a763f3057cba6db467ff999710aeecf8e")
	if err != nil {
		rsp.Msg = err.Error()
		return nil
	}

	signature, err := crypto.Sign(util.Sha256(msg), seckey)
	if err != nil {
		rsp.Msg = err.Error()
		return nil
	}

	tx_account := &sign.Transaction{
		Version:1,
		CursorNum: block_header.HeadBlockNum,
		CursorLabel: block_header.CursorLabel,
		Lifetime: block_header.HeadBlockTime + 20,
		Sender: "bottos",
		Contract: "bottos",
		Method: "newaccount",
		Param: hex.EncodeToString(account_buf),
		SigAlg: 1,
		Signature: hex.EncodeToString(signature),
	}

	ret, err := data.PushTransaction(tx_account)
	if err != nil {
		rsp.Msg = err.Error()
		return nil
	}

	log.Info("ret-account:", ret.Result.TrxHash)
	//time.Sleep(time.Duration(3)*time.Second)
	//注册用户
	rsp.Code = 1005
	ret_user, err := data.PushTransaction(&req.User)
	var did bean.Did
	buf, _ := hex.DecodeString(req.User.Param)
	pack.Unmarshal(buf, &did)
	log.Info(did.Didid)
	log.Info(did.Didinfo)
	if err != nil {
		rsp.Msg = err.Error()
		return nil
	}
	log.Info("ret-user:", ret_user)
	rsp.Code = 1
	return nil
}


func (u *User) GetAccountInfo(ctx context.Context, req *user_proto.GetAccountInfoRequest, rsp *user_proto.GetAccountInfoResponse) error {
	account_info, err:= data.AccountInfo(req.AccountName)
	if account_info != nil {
		rsp.Data = account_info
	} else {
		rsp.Code = 1006
		rsp.Msg = err.Error()
	}
	return nil
}

func (u *User) Login(ctx context.Context, req *user_proto.LoginRequest, rsp *user_proto.LoginResponse) error {
	//is_login, account := UserLogin(req.Body)
	//log.Info(account)
	//if is_login {
	//	token := create_token()
	//	is_save_token := save_token(account, token)
	//	if is_save_token {
	//		rsp.Code = 0
	//		rsp.Msg = "OK"
	//		rsp.Token = token
	//	} else {
	//		rsp.Code = 1002
	//		rsp.Msg = "Write token failure"
	//	}
	//} else {
	//	rsp.Code = 1001
	//	rsp.Msg = "Access to account information failure"
	//}
	return nil
}

func (u *User) Logout(ctx context.Context, req *user_proto.LogoutRequest, rsp *user_proto.LogoutResponse) error {
	is_logout := UserLogout(req.Token)
	log.Info(req.Token);
	if is_logout {
		rsp.Code = 0
		rsp.Msg = "OK"
	} else {
		rsp.Code = -1
		rsp.Msg = "Error"
	}
	return nil
}

//func (u *User) VerifyToken(ctx context.Context, req *user_proto.VerifyTokenRequest, rsp *user_proto.VerifyTokenResponse) error {
//	log.Info(req.Token)
//	if req.Token == "" {
//		rsp.Code = 1999
//		rsp.Msg = "Token is nil"
//		return nil
//	}
//
//	checkToken := CheckToken(req.Token)
//	if checkToken {
//		rsp.Code = 0
//		rsp.Msg = "OK"
//	} else {
//		rsp.Code = 1999
//		rsp.Msg = "Invalid Token"
//	}
//	return nil
//}

//func CheckToken(token string) bool {
//	var mgo = mgo.Session()
//	defer mgo.Close()
//	var ret UserTokenBean
//	err := mgo.DB(config.DB_NAME).C("user_token").Find(&bson.M{"token": token}).One(&ret)
//	if err != nil {
//		log.Error(err)
//		return false
//	}
//
//	if ret.Token == token {
//		if (ret.Ctime + config.TOKEN_EXPIRE_TIME > time.Now().Unix()) {
//			return true
//		}
//	}
//	return false
//}

func UserLogout(token string) bool {
	var mgo = mgo.Session()
	defer mgo.Close()
	err := mgo.DB(config.DB_NAME).C("user_token").Remove(&bson.M{"token": token})
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

func init() {
	logger, err := log.LoggerFromConfigAsFile("./config/user-log.xml")
	if err != nil{
		log.Error(err)
		panic(err)
	}
	defer logger.Flush()
	log.ReplaceLogger(logger)
}

func main() {

	service := micro.NewService(
		micro.Name("bottos.srv.user"),
		micro.Version("3.0.0"),
	)

	service.Init()

	user_proto.RegisterUserHandler(service.Server(), new(User))

	if err := service.Run(); err != nil {
		log.Critical(err)
	}

}


