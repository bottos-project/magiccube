package main

import (
	"encoding/json"
	log "github.com/cihub/seelog"
	"github.com/bottos-project/magiccube/service/exchange/proto"
	"github.com/micro/go-micro"
	api "github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
	sign "github.com/bottos-project/magiccube/service/common/signature"
	errcode "github.com/bottos-project/magiccube/error"
	"github.com/bottos-project/magiccube/tools/db/mongodb"
	"github.com/bottos-project/magiccube/config"
	"os"
	"gopkg.in/mgo.v2/bson"
)

type Exchange struct {
	Client exchange.ExchangeClient
}

func (s *Exchange) BuyAsset(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.StatusCode = 200

	//验签
	is_true, err := sign.PushVerifySign(req.Body)
	if !is_true {
		rsp.Body = errcode.ReturnError(1000, err)
		return nil
	}

	var buyAssetRequest exchange.PushRequest
	err = json.Unmarshal([]byte(req.Body), &buyAssetRequest)
	if err != nil {
		log.Error(err)
		return err
	}

	response, err := s.Client.BuyAsset(ctx, &buyAssetRequest)
	if err != nil {
		return err
	}

	rsp.Body = errcode.Return(response)
	return nil
}

func (s *Exchange) IsBuyAsset(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.StatusCode = 200

	var isBuyAssetRequest exchange.IsBuyAssetRequest
	err := json.Unmarshal([]byte(req.Body), &isBuyAssetRequest)
	if err != nil {
		log.Error(err)
		return err
	}

	is, err := sign.QueryVerifySign(req.Body)
	if !is {
		rsp.Body = errcode.ReturnError(1000, err)
		return nil
	}

	var where = &bson.M{"param.info.assetid": isBuyAssetRequest.AssetId, "param.info.username": isBuyAssetRequest.Username}

	var mgo = mgo.Session()
	defer mgo.Close()
	count, err := mgo.DB(config.DB_NAME).C("Transactions").Find(where).Count()
	log.Info(count)

	//response, err := s.Client.GetFavorite(ctx, &isBuyAssetRequest)
	if err != nil {

		return err
	}

	var result exchange.IsBuyAssetResponse
	result.Data = "false"
	if count > 0 {
		result.Data = "true"
	}
	rsp.Body = errcode.Return(result)

	return nil
}

//func (u *Exchange) QueryTx(ctx context.Context, req *api.Request, rsp *api.Response) error {
//	body := req.Body
//	log.Info(body)
//	//transfer to struct
//	var queryRequest exchange.QueryTxRequest
//	json.Unmarshal([]byte(body), &queryRequest)
//	//Checkout data format
//
//	log.Info(queryRequest)
//	ok, err := govalidator.ValidateStruct(queryRequest);
//	if !ok {
//		b, _ := json.Marshal(map[string]string{
//			"code": "-7",
//			"msg":  err.Error(),
//		})
//		rsp.StatusCode = 200
//		rsp.Body = string(b)
//		return nil
//	}
//
//	response, err := u.Client.QueryTx(ctx, &queryRequest)
//	if err != nil {
//		return err
//	}
//
//	b, _ := json.Marshal(map[string]interface{}{
//		"code": strconv.Itoa(int(response.Code)),
//		"msg":  response.Msg,
//		"data": response.Data,
//	})
//	rsp.StatusCode = 200
//	rsp.Body = string(b)
//	return nil
//}

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
		micro.Name("go.micro.api.v3.exchange"),
	)

	// parse command line flags
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&Exchange{Client: exchange.NewExchangeClient("go.micro.srv.v3.exchange", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		os.Exit(1)
	}

}
