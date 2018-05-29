package main

import (
	log "github.com/cihub/seelog"
	"github.com/micro/go-micro"
	proto "github.com/bottos-project/magiccube/service/exchange/proto"
	"golang.org/x/net/context"
	/*"github.com/bitly/go-simplejson"
	"time"
	"bytes"
	"io/ioutil"
	"net/http"
	cbb "github.com/bottos-project/magiccube/service/asset/cbb"
	"gopkg.in/mgo.v2/bson"
	"github.com/bottos-project/magiccube/service/bean"*/
	"github.com/bottos-project/magiccube/config"
	"github.com/bottos-project/magiccube/service/common/data"
	"os"
)

const (
	BASE_URL             = config.BASE_CHAIN_URL
	GET_INFO_URL         = BASE_URL + "v1/chain/get_info"
	GET_BLOCK_URL        = BASE_URL + "v1/chain/get_block"
	ABI_JSON_TO_BIN_URL  = BASE_URL + "v1/chain/abi_json_to_bin"
	ABI_BIN_TO_JSON_URL  = BASE_URL + "v1/chain/abi_bin_to_json"
	PUSH_TRANSACTION_URL = BASE_URL + "v1/chain/push_transaction"
	GET_TABLE_ROW        = BASE_URL + "v1/chain/get_table_row_by_string_key"
	STORAGE_RPC_URL      = config.BASE_RPC
)

type Exchange struct{}

func (u *Exchange) BuyAsset(ctx context.Context, req *proto.PushRequest, rsp *proto.BuyAssetResponse) error {
	i, err := data.PushTransaction(req)
	if err != nil {
		rsp.Code = 4001
		rsp.Msg = err.Error()
	}
	log.Info(i)
	return nil
}

func (u *Exchange) IsBuyAsset(ctx context.Context, req *proto.IsBuyAssetRequest, rsp *proto.IsBuyAssetResponse) error {
	return nil
}

/*func (u *Exchange) QueryTx(ctx context.Context, req *proto.QueryTxRequest, rsp *proto.QueryTxResponse) error {

	dataBody, signValue, account := "", "", ""
	//dataBody, signValue, account, data := GetSignAndDataCom(req.PostBody)
	log.Info(account)
	//get Public Key
	pubKey := cbb.GetPublicKey("account")
	//Verify Sign Local
	ok, _ := cbb.VerifySign(dataBody, signValue, pubKey)
	log.Info(ok)
	ok = true
	if !ok {
		rsp.Code = 2000
		rsp.Msg = "Verify Signature Failed."
		return nil
	}

	var pageNum, pageSize, skip int = 1, 20, 0
	if req.PageNum > 0 {
		pageNum = int(req.PageNum)
	}

	if req.PageSize > 0 && req.PageSize <= 50 {
		pageSize = int(req.PageSize)
	}

	skip = (pageNum - 1) * pageSize

	var where interface{}
	where = &bson.M{"type": "datapurchase"}

	if req.Username != "" {
		where = &bson.M{"type": "datapurchase", "data.basic_info.user_name": req.Username}
	}

	var ret []bean.TxBean

	var mgo = mgo.Session()
	defer mgo.Close()
	count, err := mgo.DB(config.DB_NAME).C("Messages").Find(where).Count()
	if err != nil {
		log.Error(err)
	}
	mgo.DB(config.DB_NAME).C("Messages").Find(where).Sort("-createdAt").Skip(skip).Limit(pageSize).All(&ret)

	var rows = []*proto.TxRow{}

	var ret2 = bean.AssetBean{}
	for _, v := range ret {
		log.Info(v.Data.BasicInfo.AssetID)
		err := mgo.DB(config.DB_NAME).C("Messages").Find(&bson.M{"type": "assetreg", "data.asset_id": v.Data.BasicInfo.AssetID}).One(&ret2)
		if err != nil {
			log.Error(err)
		}

		rows = append(rows, &proto.TxRow{
			TransactionId: v.TransactionID,
			From:          ret2.Data.BasicInfo.UserName,
			To:            v.Data.BasicInfo.UserName,
			Price:         ret2.Data.BasicInfo.Price,
			Type:          ret2.Data.BasicInfo.AssetType,
			Date:          v.CreatedAt.String(),
			BlockId:       v.BlockNum,
		})
	}

	var data = &proto.QueryTxData{
		RowCount: uint64(count),
		PageNum:  uint64(pageNum),
		Row:      rows,
	}
	log.Info(data)
	rsp.Code = 0
	rsp.Data = data
	return nil

	////Test
	//params := `service=storage&method=Storage.QueryTx&request={
	//"username":"%s"
	//}`
	//userName := req.Username
	////random := req.Random
	//
	////signature := req.Signature
	//
	//s := fmt.Sprintf(params, userName)
	//log.Info("s:", s)
	//resp, err := http.Post(STORAGE_RPC_URL, "application/x-www-form-urlencoded",
	//	strings.NewReader(s))
	//
	//log.Info("resp:", resp)
	////log.Info("err", err)
	//if err != nil {
	//	return err
	//}
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return err
	//} else {
	//	js, _ := simplejson.NewJson([]byte(body))
	//	log.Info("jss", js)
	//	result, _ := json.Marshal(js.Get("FileList"))
	//	if js.Get("code").MustInt() == 1 {
	//
	//		rsp.Code = 1
	//		rsp.Msg = "Get File List Successful!"
	//		rsp.Data = string(result)
	//	}
	//	return nil
	//}
	//end_time := time.Now().UnixNano() / int64(time.Millisecond)
	//log.Info("Time:", end_time-start_time)
	return nil
}*/

func init() {
	logger, err := log.LoggerFromConfigAsFile("./config/exc-log.xml")
	if err != nil{
		log.Error(err)
		panic(err)
	}
	defer logger.Flush()
	log.ReplaceLogger(logger)
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.v3.exchange"),
		micro.Version("3.0.0"),
	)

	service.Init()

	//proto.RegisterUserHandler(service.Server(), new(Asset))
	//user_proto.RegisterUserHandler(service.Server(), new(User))
	proto.RegisterExchangeHandler(service.Server(), new(Exchange))

	if err := service.Run(); err != nil {
		os.Exit(1)
	}
}
