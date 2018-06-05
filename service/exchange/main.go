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
	proto "github.com/bottos-project/magiccube/service/exchange/proto"
	log "github.com/cihub/seelog"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
	/*"github.com/bitly/go-simplejson"
	"time"
	"bytes"
	"io/ioutil"
	"net/http"
	cbb "github.com/bottos-project/magiccube/service/asset/cbb"
	"gopkg.in/mgo.v2/bson"
	"github.com/bottos-project/magiccube/service/bean"*/
	"os"

	"github.com/bottos-project/magiccube/config"
	"github.com/bottos-project/magiccube/service/common/data"
)

const (
	// BASE_URL base url
	BASE_URL = config.BASE_CHAIN_URL
)

// Exchange struct
type Exchange struct{}

// BuyAsset on chain
func (u *Exchange) BuyAsset(ctx context.Context, req *proto.PushRequest, rsp *proto.BuyAssetResponse) error {
	i, err := data.PushTransaction(req)
	if err != nil {
		rsp.Code = 4001
		rsp.Msg = err.Error()
	}
	log.Info(i)
	return nil
}

// IsBuyAsset is or not
func (u *Exchange) IsBuyAsset(ctx context.Context, req *proto.IsBuyAssetRequest, rsp *proto.IsBuyAssetResponse) error {
	return nil
}

// GrantCredit on chain
func (u *Exchange) GrantCredit(ctx context.Context, req *proto.PushRequest, rsp *proto.BuyAssetResponse) error {
	i, err := data.PushTransaction(req)
	if err != nil {
		rsp.Code = 4011
		rsp.Msg = err.Error()
	}
	log.Info(i)
	return nil
}

// CancelCredit on chain
func (u *Exchange) CancelCredit(ctx context.Context, req *proto.PushRequest, rsp *proto.BuyAssetResponse) error {
	i, err := data.PushTransaction(req)
	if err != nil {
		rsp.Code = 4021
		rsp.Msg = err.Error()
	}
	log.Info(i)
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
	if err != nil {
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
