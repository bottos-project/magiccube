package main

import (
	log "github.com/cihub/seelog"
	"github.com/micro/go-micro"
	user_proto "github.com/bottos-project/bottos/service/user/proto"
	"golang.org/x/net/context"
	"github.com/bottos-project/bottos/service/common/data"
	"encoding/hex"
	"github.com/bottos-project/crypto-go/crypto"
	"github.com/protobuf/proto"
	"github.com/bottos-project/bottos/service/common/util"
	push_sign "github.com/bottos-project/bottos/service/common/signature/push"
	pack "github.com/bottos-project/bottos/core/contract/msgpack"
	"github.com/bottos-project/bottos/service/common/bean"
	"github.com/bottos-project/bottos/tools/db/mongodb"
	"gopkg.in/mgo.v2/bson"
	"github.com/bottos-project/bottos/config"
)

type User struct{}

func (u *User) GetBlockHeader(ctx context.Context, req *user_proto.GetBlockHeaderRequest, rsp *user_proto.GetBlockHeaderResponse) error {
	block_header, err:= data.BlockHeader()
	log.Info(block_header)
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
	tx_account_sign := &push_sign.TransactionSign{
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

	tx_account := &bean.TxBean{
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

	//did
	var did bean.Did
	d, _:= hex.DecodeString(req.User.Param)
	pack.Unmarshal(d, &did)
	log.Info("did:", did)

	//注册用户
	rsp.Code = 1005
	ret_user, err := data.PushTransaction(&req.User)

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
	return nil
}

func (u *User) Favorite(ctx context.Context, req *user_proto.FavoriteRequest, rsp *user_proto.FavoriteResponse) error {

	i, err := data.PushTransaction(req)

	if i == nil {
		rsp.Code = 1008
		rsp.Msg = err.Error()
	}
	return nil
}

func (u *User) GetFavorite(ctx context.Context, req *user_proto.GetFavoriteRequest, rsp *user_proto.GetFavoriteResponse) error {
	var pageNum, pageSize, skip int= 1, 20, 0
	if req.PageNum > 0 {
		pageNum = int(req.PageNum)
	}

	if req.PageSize > 0 && req.PageSize <= 50{
		pageSize = int(req.PageSize)
	}

	if len(req.GoodsType) < 1{
		req.GoodsType = "asset"
	}

	skip = (pageNum - 1) *  pageSize

	var mgo = mgo.Session()
	defer mgo.Close()
	var where = bson.M{
		"param.optype": bson.M{"$in": []int32{1,2}},
		"param.username": req.Username,
		"param.goodstype":req.GoodsType}
	log.Info(where)
	count, err:=mgo.DB(config.DB_NAME).C("pre_favoritepro").Find(where).Count()
	log.Info(count)
	if err != nil {
		log.Error(err)
	}
	var ret []*bean.Favorite
	mgo.DB(config.DB_NAME).C("pre_favoritepro").Find(where).Sort("-_id").Limit(pageSize).Skip(skip).All(&ret)

	var rows []*user_proto.FavoriteData

	if req.GoodsType == "asset" {
		var ret2 bean.AssetBean
		for _, v := range ret {
			err := mgo.DB(config.DB_NAME).C("pre_assetreg").Find(&bson.M{"param.assetid": v.Param.Goodsid}).Sort("-_id").Limit(1).One(&ret2)
			if err != nil {
				log.Error(err)
			}

			rows = append(rows, &user_proto.FavoriteData{
				Username:ret2.Param.Info.UserName,
				GoodsId:v.Param.Goodsid,
				GoodsName:ret2.Param.Info.AssetName,
				Price:ret2.Param.Info.Price,
				Time:uint64(v.CreateTime.Unix()),
			})
		}
	}



	var data = &user_proto.FavoriteArr{
		PageNum: uint64(pageNum),
		RowCount: uint64(count),
		Row: rows,
	}

	rsp.Data = data
	return nil
}

func (u *User) Transfer(ctx context.Context, req *user_proto.PushTxRequest, rsp *user_proto.PushTxResponse) error {

	i, err := data.PushTransaction(req)

	if i == nil {
		rsp.Code = 1008
		rsp.Msg = err.Error()
	}
	return nil
}

func (u *User) QueryMyNotice(ctx context.Context, req *user_proto.QueryMyNoticeRequest, rsp *user_proto.QueryMyNoticeResponse) error {

	var pageNum, pageSize, skip int= 1, 20, 0
	if req.PageNum > 0 {
		pageNum = int(req.PageNum)
	}

	if req.PageSize > 0 {
		pageSize = int(req.PageSize)
	}

	skip = (pageNum - 1) *  pageSize

	var where interface{}
	//where = bson.M{"param.info.optype": bson.M{"$in": []int32{1,2}}}
	if len(req.Username) > 0{
		where = &bson.M{ "param.info.consumer": req.Username}
	}

	log.Info(where)

	var ret []bean.PreSaleBean

	var mgo = mgo.Session()
	defer mgo.Close()
	count, err:= mgo.DB(config.DB_NAME).C("pre_presale").Find(where).Count()
	log.Info(count)
	if err != nil {
		log.Error(err)
	}
	mgo.DB(config.DB_NAME).C("pre_presale").Find(where).Sort("-_id").Skip(skip).Limit(pageSize).All(&ret)

	var rows = []*user_proto.QueryNoticeRow{}
	for _, v := range ret {

		rows = append(rows, &user_proto.QueryNoticeRow{
			NoticeId : v.Param.Datapresaleid,
			Username : v.Param.Info.Username,
			AssetId : v.Param.Info.Assetid,
			//AssetName : v.Param.Info.ass,
			DataReqId : v.Param.Info.Datareqid,
			Consumer : v.Param.Info.Consumer,
			Time : uint64(v.CreateTime.Unix()),
		})
	}

	var data = &user_proto.QueryNoticeData{
		RowCount: uint32(count),
		PageNum: uint32(pageNum),
		Row:rows,
	}
	log.Info(data)
	rsp.Data = data
	return nil
}

func (u *User) QueryMyPreSale(ctx context.Context, req *user_proto.QueryMyNoticeRequest, rsp *user_proto.QueryMyNoticeResponse) error {

	var pageNum, pageSize, skip int= 1, 20, 0
	if req.PageNum > 0 {
		pageNum = int(req.PageNum)
	}

	if req.PageSize > 0 {
		pageSize = int(req.PageSize)
	}

	skip = (pageNum - 1) *  pageSize

	var where interface{}
	//where = bson.M{"param.info.optype": bson.M{"$in": []int32{1,2}}}
	if len(req.Username) > 0{
		where = &bson.M{ "param.info.username": req.Username}
	}

	log.Info(where)

	var ret []bean.PreSaleBean

	var mgo = mgo.Session()
	defer mgo.Close()
	count, err:= mgo.DB(config.DB_NAME).C("pre_presale").Find(where).Count()
	log.Info(count)
	if err != nil {
		log.Error(err)
	}
	mgo.DB(config.DB_NAME).C("pre_presale").Find(where).Sort("-_id").Skip(skip).Limit(pageSize).All(&ret)

	var rows = []*user_proto.QueryNoticeRow{}
	for _, v := range ret {

		rows = append(rows, &user_proto.QueryNoticeRow{
			NoticeId : v.Param.Datapresaleid,
			Username : v.Param.Info.Username,
			AssetId : v.Param.Info.Assetid,
			//AssetName : v.Param.Info.ass,
			DataReqId : v.Param.Info.Datareqid,
			Consumer : v.Param.Info.Consumer,
			Time : uint64(v.CreateTime.Unix()),
		})
	}

	var data = &user_proto.QueryNoticeData{
		RowCount: uint32(count),
		PageNum: uint32(pageNum),
		Row:rows,
	}
	log.Info(data)
	rsp.Data = data
	return nil
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
