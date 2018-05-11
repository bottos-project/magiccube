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

	//ref_block_num := GetBolckNum()
	//log.Info("ref_block_num:", ref_block_num)
	//time.Sleep(1)
	//if ref_block_num < 0 {
	//	rsp.Code = 1131
	//	rsp.Msg = "Data[ref_block_num] exception"
	//	return nil
	//}
	//
	//ref_block_prefix, timestamp := GetBlockPrefix(ref_block_num)
	//log.Info("ref_block_prefix:", ref_block_prefix)
	//log.Info("timestamp:", timestamp)
	//time.Sleep(1)
	//if ref_block_prefix < 0 {
	//	rsp.Code = 1132
	//	rsp.Msg = "Data[ref_block_prefix] exception"
	//	return nil
	//}
	//
	//accoutBin := AccountGetBin(req.Username, req.OwnerPubKey, req.ActivePubKey)
	//log.Info("accoutBin:", accoutBin)
	//time.Sleep(1)
	//if accoutBin == "" {
	//	rsp.Code = 1133
	//	rsp.Msg = "Data[account_bin] exception"
	//	return nil
	//}
	//
	//expirationTime := ExpirationTime(timestamp, 20)
	//accountInfo, msg := AccountPushTransaction(ref_block_num, ref_block_prefix, expirationTime, accoutBin)
	//log.Info("accountInfo:", accountInfo)
	//if !accountInfo {
	//	if msg == "Already" {
	//		rsp.Code = 1103
	//		rsp.Msg = "Account has already existed"
	//		return nil
	//	}
	//	rsp.Code = 1101
	//	rsp.Msg = msg
	//	return nil
	//}
	//time.Sleep(1)
	//b, _ := json.Marshal(req.UserInfo)
	//userBin := UserGetBin(req.Username, string(b))
	//log.Info("userBin:", userBin)
	//time.Sleep(1)
	//if userBin == "" {
	//	rsp.Code = 1104
	//	rsp.Msg = "Data[user_bin] exception"
	//	return nil
	//}
	//userInfo := UserPushTransaction(ref_block_num, ref_block_prefix, expirationTime, userBin, req.Username)
	//log.Info("userInfo:", userInfo)
	//if userInfo != "" {
	//	rsp.Code = 0
	//	rsp.Msg = "Registered user success"
	//} else {
	//	rsp.Code = 1102
	//	rsp.Msg = "Registered user failure"
	//}
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


