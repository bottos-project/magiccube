/*Copyright 2017~2022 The Bottos Authors
  This file is part of the Bottos Service Layer
  Created by Developers Team of Bottos.

  This program is free software: you can distribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with Bottos. If not, see <http://www.gnu.org/licenses/>.
*/
package main

import (
	"encoding/json"
	"os"

	"github.com/bottos-project/magiccube/config"
	errcode "github.com/bottos-project/magiccube/error"
	sign "github.com/bottos-project/magiccube/service/common/signature"
	"github.com/bottos-project/magiccube/service/exchange/proto"
	"github.com/bottos-project/magiccube/tools/db/mongodb"
	log "github.com/cihub/seelog"
	"github.com/micro/go-micro"
	api "github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2/bson"
)

// Exchange struct
type Exchange struct {
	Client exchange.ExchangeClient
}

// BuyAsset on chain
func (s *Exchange) BuyAsset(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.StatusCode = 200

	//verify signature
	isTrue, err := sign.PushVerifySign(req.Body)
	if !isTrue {
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

// IsBuyAsset or not
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

// GrantCredit on chain
func (s *Exchange) GrantCredit(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.StatusCode = 200

	//verify signature
	isTrue, err := sign.PushVerifySign(req.Body)
	if !isTrue {
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

// CancelCredit on chain
func (s *Exchange) CancelCredit(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.StatusCode = 200

	//verify signature
	isTrue, err := sign.PushVerifySign(req.Body)
	if !isTrue {
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
	if err != nil {
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
