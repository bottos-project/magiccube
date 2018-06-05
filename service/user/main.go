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
	log "github.com/cihub/seelog"
	"github.com/micro/go-micro"
	user_proto "github.com/bottos-project/magiccube/service/user/proto"
	"golang.org/x/net/context"
	"github.com/bottos-project/magiccube/service/common/data"
	"encoding/hex"
	"github.com/bottos-project/crypto-go/crypto"
	"github.com/protobuf/proto"
	"github.com/bottos-project/magiccube/service/common/util"
	push_sign "github.com/bottos-project/magiccube/service/common/signature/push"
	pack "github.com/bottos-project/msgpack-go"
	"github.com/bottos-project/magiccube/service/common/bean"
	"github.com/bottos-project/magiccube/tools/db/mongodb"
	"gopkg.in/mgo.v2/bson"
	"github.com/bottos-project/magiccube/config"
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
	log.Info("req:", req)
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
		Sender:      "delta",
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
	//配对的pubkey   0454f1c2223d553aa6ee53ea1ccea8b7bf78b8ca99f3ff622a3bb3e62dedc712089033d6091d77296547bc071022ca2838c9e86dec29667cf740e5c9e654b6127f
	seckey, err := hex.DecodeString("b799ef616830cd7b8599ae7958fbee56d4c8168ffd5421a16025a398b8a4be45")
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
		Sender:      "delta",
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
	
	//time.Sleep(time.Duration(2)*time.Second)

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

func (u *User) QueryMyBuy(ctx context.Context, req *user_proto.QueryMyBuyRequest, rsp *user_proto.QueryMyBuyResponse) error {
	var pageNum, pageSize, skip int= 1, 20, 0
	if req.PageNum > 0 {
		pageNum = int(req.PageNum)
	}

	if req.PageSize > 0 && req.PageSize < 20{
		pageSize = int(req.PageSize)
	}

	skip = (pageNum - 1) *  pageSize

	var where = &bson.M{"method":"buydata", "param.info.username": req.Username}

	var ret []bean.Buy

	var mgo = mgo.Session()
	defer mgo.Close()
	count, err:= mgo.DB(config.DB_NAME).C("Transactions").Find(where).Count()
	log.Info(count)
	if err != nil {
		log.Error(err)
	}
	mgo.DB(config.DB_NAME).C("Transactions").Find(where).Sort("-create_time").Skip(skip).Limit(pageSize).All(&ret)

	var rows = []*user_proto.Buy{}
	for _, v := range ret {
		var ret2 bean.AssetBean
		mgo.DB(config.DB_NAME).C("pre_assetreg").Find(bson.M{"param.assetid":v.Param.Info.AssetId, "create_time": bson.M{"$lt": v.CreateTime}}).Sort("-create_time").Limit(1).One(&ret2)
		rows = append(rows, &user_proto.Buy{
			ExchangeId : v.Param.DataExchangeId,
			Username : ret2.Param.Info.UserName,
			AssetId : v.Param.Info.AssetId,
			AssetName : ret2.Param.Info.AssetName,
			AssetType : ret2.Param.Info.AssetType,
			FeatureTag : ret2.Param.Info.FeatureTag,
			Price : ret2.Param.Info.Price,
			SampleHash : ret2.Param.Info.SampleHash,
			StorageHash : ret2.Param.Info.StorageHash,
			Expiretime : uint64(ret2.Param.Info.ExpireTime),
			Timestamp: uint64(v.CreateTime.Unix()),
		})
	}

	rsp.Data = &user_proto.BuyData{
		RowCount: int32(count),
		PageNum: int32(pageNum),
		Row:rows,
	}
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
		micro.Name("go.micro.srv.v3.user"),

		micro.Version("3.0.0"),
	)

	service.Init()

	user_proto.RegisterUserHandler(service.Server(), new(User))

	if err := service.Run(); err != nil {
		log.Critical(err)
	}

}
