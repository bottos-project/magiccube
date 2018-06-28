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
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/bitly/go-simplejson"
	"github.com/bottos-project/magiccube/config"
	proto "github.com/bottos-project/magiccube/service/asset/proto"
	"github.com/bottos-project/magiccube/service/common/bean"
	"github.com/bottos-project/magiccube/service/common/data"
	"github.com/bottos-project/magiccube/tools/db/mongodb"
	log "github.com/cihub/seelog"
	"github.com/micro/go-micro"
	"github.com/mikemintang/go-curl"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2/bson"
)

const (
	// BASE_URL BASE URL
	BASE_URL = config.BASE_CHAIN_URL
	// GET_TABLE_ROW_BY_STRING get one row service
	GET_TABLE_ROW_BY_STRING = BASE_URL + "v1/chain/get_table_row_by_string_key"
	// STORAGE_RPC_URL strorage url
	STORAGE_RPC_URL = config.BASE_RPC
)

// Asset Struct
type Asset struct{}

// RegisterFile On Chain
func (u *Asset) RegisterFile(ctx context.Context, req *proto.PushTxRequest, rsp *proto.RegisterFileResponse) error {
	i, err := data.PushTransaction(req)
	if err != nil {
		rsp.Code = 2011
		rsp.Msg = err.Error()
	}
	log.Info(i)
	return nil
}

// RegisterAsset On Chain
func (u *Asset) RegisterAsset(ctx context.Context, req *proto.PushTxRequest, rsp *proto.RegisterResponse) error {
	i, err := data.PushTransaction(req)
	if err != nil {
		log.Error(err)
		rsp.Code = 2021
		rsp.Msg = err.Error()
	}
	log.Info(i)
	return nil
}

// QueryAsset from Chain
func (u *Asset) QueryAsset(ctx context.Context, req *proto.QueryRequest, rsp *proto.QueryAssetResponse) error {

	var pageNum, pageSize, skip int = 1, 20, 0
	if req.PageNum > 0 {
		pageNum = int(req.PageNum)
	}

	if req.PageSize > 0 && req.PageSize < 20 {
		pageSize = int(req.PageSize)
	}

	skip = (pageNum - 1) * pageSize

	var where interface{}
	where = bson.M{"param.info.optype": bson.M{"$in": []int32{1, 2}}}
	if len(req.Username) > 0 {
		where = &bson.M{"param.info.optype": bson.M{"$in": []uint32{1, 2}}, "param.info.username": req.Username}
	}

	if req.AssetType > 0 {
		where = bson.M{"param.info.optype": bson.M{"$in": []uint32{1, 2}}, "param.info.assettype": req.AssetType}
	}

	if len(req.Username) > 0 && req.AssetType > 0 {
		where = bson.M{"param.info.optype": bson.M{"$in": []uint32{1, 2}}, "param.info.username": req.Username, "param.info.assettype": req.AssetType}
	}
	log.Info(where)

	var ret []bean.AssetBean

	var mgo = mgo.Session()
	defer mgo.Close()
	count, err := mgo.DB(config.DB_NAME).C("pre_assetreg").Find(where).Count()
	log.Info(count)
	if err != nil {
		log.Error(err)
	}
	mgo.DB(config.DB_NAME).C("pre_assetreg").Find(where).Sort("-_id").Skip(skip).Limit(pageSize).All(&ret)

	var rows = []*proto.AssetData{}
	for _, v := range ret {
		rows = append(rows, &proto.AssetData{
			AssetId:     v.Param.AssetId,
			Username:    v.Param.Info.UserName,
			AssetName:   v.Param.Info.AssetName,
			AssetType:   v.Param.Info.AssetType,
			FeatureTag:  v.Param.Info.FeatureTag,
			SampleHash:  v.Param.Info.SampleHash,
			StorageHash: v.Param.Info.StorageHash,
			ExpireTime:  v.Param.Info.ExpireTime,
			TokenType:v.Param.Info.TokenType,
			Price:       v.Param.Info.Price,
			OpType:      v.Param.Info.OpType,
			Description: v.Param.Info.Description,
			CreateTime:  uint64(v.CreateTime.Unix()),
		})
	}

	var data = &proto.QueryAssetData{
		RowCount: uint32(count),
		PageNum:  uint32(pageNum),
		Row:      rows,
	}
	log.Info(data)
	rsp.Data = data
	return nil
}

// QueryAssetByID from Chain
func (u *Asset) QueryAssetByID(ctx context.Context, req *proto.QueryAssetByIDRequest, rsp *proto.QueryAssetInfoResponse) error {

	var where interface{}
	where = &bson.M{"param.info.optype": bson.M{"$in": []uint32{1, 2}}, "param.assetid": req.AssetId}

	var ret bean.AssetBean

	var mgo = mgo.Session()
	defer mgo.Close()
	err := mgo.DB(config.DB_NAME).C("pre_assetreg").Find(where).Sort("-_id").One(&ret)
	if err != nil {
		log.Error(err)
	}

	var count, count1 = 0, 0
	if len(req.Sender) > 0 {
		var where1 = &bson.M{"param.optype": bson.M{"$in": []uint32{1, 2}}, "param.goodsid": req.AssetId, "param.username": req.Sender}
		log.Debug(where1)
		count, err = mgo.DB(config.DB_NAME).C("pre_favoritepro").Find(where1).Count()
		log.Debug("pre_favoritepro: ", count)
		if err != nil {
			log.Error(err)
		}
		var where2 = &bson.M{"method": "buydata", "param.info.assetid": req.AssetId, "param.info.username": req.Sender}
		log.Debug(where2)
		count1, err = mgo.DB(config.DB_NAME).C("Transactions").Find(where2).Count()
		log.Debug("Transactions: ", count1)
		if err != nil {
			log.Error(err)
		}
	}

	rsp.Data = &proto.QueryAssetInfoData{
		AssetId:        ret.Param.AssetId,
		Username:       ret.Param.Info.UserName,
		AssetName:      ret.Param.Info.AssetName,
		AssetType:      ret.Param.Info.AssetType,
		FeatureTag:     ret.Param.Info.FeatureTag,
		SampleHash:     ret.Param.Info.SampleHash,
		StorageHash:    ret.Param.Info.StorageHash,
		ExpireTime:     ret.Param.Info.ExpireTime,
		TokenType:      ret.Param.Info.TokenType,
		Price:          ret.Param.Info.Price,
		OpType:         ret.Param.Info.OpType,
		Description:    ret.Param.Info.Description,
		CreateTime:     uint64(ret.CreateTime.Unix()),
		FavoriteFlag:   count > 0,
		IsBuyAssetFlag: count1 > 0,
	}

	return nil
}

// QueryUploadedData by My
func (u *Asset) QueryUploadedData(ctx context.Context, req *proto.QueryRequest, rsp *proto.QueryUploadedDataResponse) error {
	var pageNum, pageSize, skip int = 1, 20, 0
	if req.PageNum > 0 {
		pageNum = int(req.PageNum)
	}

	if req.PageSize > 0 && req.PageSize < 20 {
		pageSize = int(req.PageSize)
	}

	skip = (pageNum - 1) * pageSize

	var where interface{}
	//where = bson.M{"param.info.optype": bson.M{"$in": []int32{1, 2}}}
	if len(req.Username) > 0 {
		where = &bson.M{"param.info.optype": bson.M{"$in": []uint32{1, 2}}, "param.info.username": req.Username}
	}

	log.Info(where)

	var ret []bean.FileBean

	var mgo = mgo.Session()
	defer mgo.Close()
	count, err := mgo.DB(config.DB_NAME).C("pre_datafilereg").Find(where).Count()
	log.Info(count)
	if err != nil {
		log.Error(err)
	}
	mgo.DB(config.DB_NAME).C("pre_datafilereg").Find(where).Sort("-_id").Skip(skip).Limit(pageSize).All(&ret)
	log.Debug(ret)
	var rows = []*proto.QueryUploadedRow{}
	for _, v := range ret {

		rows = append(rows, &proto.QueryUploadedRow{
			FileHash:   v.Param.FileId,
			Username:   v.Param.Info.UserName,
			FileSize:   v.Param.Info.FileSize,
			FileName:   v.Param.Info.FileName,
			FilePolicy: v.Param.Info.FilePolicy,
			FileNumber: v.Param.Info.FileNumber,
			SimOrAss:   v.Param.Info.SimOrAss,
			OpType:     v.Param.Info.OpType,
			StoreAddr:  v.Param.Info.StoreAddr,
			CreateTime: uint64(v.CreateTime.Unix()),
		})
	}

	var data = &proto.QueryUploadedData{
		RowCount: uint32(count),
		PageNum:  uint32(pageNum),
		Row:      rows,
	}
	log.Info(data)
	rsp.Data = data
	return nil

	//var pageNum, pageSize, skip int = 1, 20, 0
	//if req.PageNum > 0 {
	//	pageNum = int(req.PageNum)
	//}
	//
	//if req.PageSize > 0 && req.PageSize <= 50 {
	//	pageSize = int(req.PageSize)
	//}
	//
	//skip = (pageNum - 1) * pageSize
	//
	//var where interface{}
	//where = &bson.M{"type": "datafilereg"}
	//log.Info(req.Username)
	//if req.Username != "" {
	//	where = &bson.M{"type": "datafilereg", "data.basic_info.user_name": req.Username}
	//	//where = &bson.M{"type": "assetreg", "data.basic_info.user_name": req.Username, "data.basic_info.feature_tag": req.FeatureTag}
	//} else {
	//	//if req.Username != "" {
	//	//where = &bson.M{"type": "datafilereg"}
	//	//}
	//	return errors.New("usename is nil")
	//}
	//
	//log.Info("where:", where)
	//
	//var ret []bean.FileBean
	//
	//var mgo = mgo.Session()
	//
	//defer mgo.Close()
	//
	//count, err := mgo.DB(config.DB_NAME).C("Messages").Find(where).Count()
	//if err != nil {
	//	log.Error(err)
	//}
	//mgo.DB(config.DB_NAME).C("Messages").Find(where).Skip(skip).Limit(pageSize).All(&ret)
	////mgo.DB(config.DB_NAME).C("Messages").Find(where).Sort("data.basic_info.publish_date").Skip(skip).Limit(int(req.PageSize)).All(&ret)
	//
	//var rows = []*proto.QueryUploadedRow{}
	//for _, v := range ret {
	//	rows = append(rows, &proto.QueryUploadedRow{
	//		Username:   v.Data.BasicInfo.UserName,
	//		FileHash:   v.Data.FileHash,
	//		FileName:   v.Data.BasicInfo.FileName,
	//		FileSize:   v.Data.BasicInfo.FileSize,
	//		FilePolicy: v.Data.BasicInfo.FilePolicy,
	//		FileNumber: v.Data.BasicInfo.FileNumber,
	//		AuthPath:   v.Data.BasicInfo.AuthPath,
	//		CreateTime: v.CreatedAt.String(),
	//	})
	//}
	//
	//var data = &proto.QueryUploadedData{
	//	RowCount: uint64(count),
	//	PageNum:  uint64(pageNum),
	//	Row:      rows,
	//}
	//log.Info(data)
	//rsp.Code = 0
	//rsp.Data = data
	//rsp.Msg = "OK"
	//
	//return nil
}

/*
// Get File Upload URL fs
func (u *Asset) GetFileUploadURL(ctx context.Context, req *proto.GetFileUploadURLRequest, rsp *proto.GetFileUploadURLResponse) error {
	log.Info("Start Get File URL!")
	start_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Info("reqBody:" + req.PostBody)
	//log.Info(userName)
	//get Public Key

	//log.Info(ok)
	//get strore Address
	js, _ := simplejson.NewJson([]byte(req.PostBody))
	log.Info("js", js)
	userName := ""

	userName = js.Get("userName").MustString()
	fileName := js.Get("fileName").MustString()
	FileSize := js.Get("fileSize").MustUint64()
	FilePolicy := js.Get("filePolicy").MustString()
	FileNumber := js.Get("fileNumber").MustUint64()
	Signature := js.Get("signatures").MustString()

	//Test
	params := `service=storage&method=Storage.GetFileUploadURL&request={
	"Username":"%s",
	"file_name":"%s",
	"file_size":%d,
	"file_policy":"%s",
	"file_number":%d,
	"signature":"%s"
	}`
	s := fmt.Sprintf(params, userName, fileName, FileSize, FilePolicy, FileNumber, Signature)
	resp, err := http.Post(STORAGE_RPC_URL, "application/x-www-form-urlencoded",
		strings.NewReader(s))

	log.Error("Get Data from Chain err:", err)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	} else {
		jss, _ := simplejson.NewJson([]byte(body))
		presigned_put_url := jss.Get("presigned_put_url").MustString()
		rsp.Code = 1
		rsp.Msg = "get FileUploadURL Successful!"
		rsp.Data = presigned_put_url
		log.Debug(presigned_put_url)
		return nil
	}
	log.Info(string(body))

	//Test

	//cl := storage.NewStorageClient("storage", client.DefaultClient)
	//rspd, err := cl.GetFileUploadURL(context.Background(), &storage.FileUploadRequest{
	//	Username:   js.Get("userName").MustString(),
	//	FileName:   js.Get("fileName").MustString(),
	//	FileSize:   js.Get("fileSize").MustUint64(),
	//	FilePolicy: js.Get("filePolicy").MustString(),
	//	FileNumber: js.Get("fileNumber").MustUint64(),
	//	Signature:  js.Get("signatures").MustString(),
	//	})
	//log.Info("rspd:", rspd)
	//log.Info(rspd.PresignedPutUrl)

	end_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Info("Time:", end_time-start_time)
	if err != nil {
		log.Info(err)
		return nil
	} else {
		rsp.Code = 1
		rsp.Msg = "get FileUploadURL Successful!"
		//rsp.Data = rspd.PresignedPutUrl
		return nil
	}
}

func (u *Asset) GetFileUploadStat(ctx context.Context, req *proto.GetFileUploadStatRequest, rsp *proto.GetFileUploadStatResponse) error {

	//Test
	params := `service=storage&method=Storage.GetFileUploadStat&request={
	"username":"%s",
	"file_name":"%s"
	}`
	userName := req.Username
	fileName := req.FileName
	log.Info(userName, fileName)
	s := fmt.Sprintf(params, userName, fileName)
	resp, err := http.Post(STORAGE_RPC_URL, "application/x-www-form-urlencoded",
		strings.NewReader(s))

	//log.Info("resp:",resp)
	//log.Info("err", err)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	} else {
		jss, _ := simplejson.NewJson([]byte(body))
		log.Info("jss", jss)
		uploadStat := jss.Get("result").MustString()
		message := jss.Get("message").MustString()
		fileSize := jss.Get("size").MustInt64()
		log.Info(fileSize)
		if uploadStat == "200" {
			rsp.Code = 1
			rsp.Msg = message
			rsp.Data = strconv.FormatInt(fileSize, 10)
			log.Info("filesize:", fileSize)
		}
		return nil
	}
	//log.Info(string(body))

	//Test

	//cl := storage.NewStorageClient("storage", client.DefaultClient)
	//rspd, err := cl.GetFileUploadURL(context.Background(), &storage.FileUploadRequest{
	//	Username:   js.Get("userName").MustString(),
	//	FileName:   js.Get("fileName").MustString(),
	//	FileSize:   js.Get("fileSize").MustUint64(),
	//	FilePolicy: js.Get("filePolicy").MustString(),
	//	FileNumber: js.Get("fileNumber").MustUint64(),
	//	Signature:  js.Get("signatures").MustString(),
	//	})

	//rsp.Code = 1
	//rsp.Msg = "OK"
	//rsp.Data = ""
	//return nil
}

func GetAssetList(queryPara *proto.QueryRequest) (string, int, error) {
	//TODO
	js, _ := simplejson.NewJson([]byte(queryPara.String()))
	cl := storage.NewStorageClient("storage", client.DefaultClient)
	rspd, err := cl.GetFileUploadURL(context.Background(), &storage.FileUploadRequest{
		Username:   js.Get("Type").MustString(),
		FileName:   js.Get("Time").MustString(),
		FileSize:   js.Get("FileSize").MustUint64(),
		FilePolicy: js.Get("FilePolicy").MustString(),
		FileNumber: js.Get("FileNumber").MustUint64(),
		Signature:  js.Get("Signature").MustString(),
	})
	//postData := map[string]interface{}{
	//	"scope":       "usermng",
	//	"code":        "usermng",
	//	"table":       "userreginfo",
	//	"json":        true,
	//	"strkeyvalue": username,
	//}
	//req := curl.NewRequest()
	//resp, err := req.SetUrl(GET_TABLE_ROW).SetPostData(postData).Post()
	if err != nil {
		return rspd.PresignedPutUrl, 1, err
	} else {
		return "Failed", 0, err
	}
}
*/

// GetDownLoadURL -get File Strage Url
func (u *Asset) GetDownLoadURL(ctx context.Context, req *proto.GetDownLoadURLRequest, rsp *proto.GetDownLoadURLResponse) error {

	//
	params := `service=storage&method=Storage.GetDownLoadURL&request={
	"username":"%s",
	"file_name":"%s"
	}`
	userName := req.Username
	fileName := req.FileId

	//signature := req.Signature

	s := fmt.Sprintf(params, userName, fileName)
	log.Info("s:", s)
	resp, err := http.Post(STORAGE_RPC_URL, "application/x-www-form-urlencoded",
		strings.NewReader(s))

	log.Info("resp:", resp)
	//log.Info("err", err)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	js, _ := simplejson.NewJson([]byte(body))
	log.Info("jss", js)
	//result,_ :=json.Marshal(js.Get("FileList"))
	downLoadURL := js.Get("presigned_get_url").MustString()
	if js.Get("result").MustString() == "200" {

		rsp.Code = 1
		rsp.Msg = "Get downLoad URL Successful!"
		rsp.Data = downLoadURL
	}
	//err,_ := json.Marshal(js)
	rsp.Msg = "Get downLoad URL Successful!"
	return nil

}

// GetTableRowByString -get One Row form chain
func GetTableRowByString(username string, fileId string) (string, int, error) {
	postData := map[string]interface{}{
		"scope":       username,
		"code":        "datafilemng",
		"table":       "fileauthinfo",
		"json":        true,
		"strkeyvalue": fileId + username,
	}
	req := curl.NewRequest()
	log.Info("postData:", postData)
	resp, err := req.SetUrl(GET_TABLE_ROW_BY_STRING).SetPostData(postData).Post()
	if err != nil {
		return "", resp.Raw.StatusCode, err
	}
	log.Info(resp.Body)
	if resp.IsOk() {
		return resp.Body, resp.Raw.StatusCode, err
	}
	return "", resp.Raw.StatusCode, err

}

// GetAssetByIdNoStoPath -get one Asset
func GetAssetByIdNoStoPath(assertID string) ([]*proto.AssetData, error) {
	var where interface{}
	//where = &bson.M{"type": "assetreg"}
	log.Info(assertID)
	if assertID != "" {
		where = bson.M{"param.info.optype": bson.M{"$in": []int32{1, 2}}, "param.assetid": assertID}
	} else {
		return nil, nil

	}

	log.Info(where)

	var ret []bean.AssetBean
	//var ret1 []bean.AssetBean

	var mgo = mgo.Session()
	defer mgo.Close()

	count, err := mgo.DB(config.DB_NAME).C("pre_assetreg").Find(where).Count()
	log.Info(count)
	if count > 1 {
		log.Error(count)
		return nil, nil
	}
	if err != nil {
		log.Error(err)
	}

	mgo.DB(config.DB_NAME).C("pre_assetreg").Find(where).Sort("-_id").All(&ret)
	//mgo.DB(config.DB_NAME).C("Messages").Find(where).Sort("data.basic_info.publish_date").Skip(skip).Limit(int(req.PageSize)).All(&ret)

	//ret1 = append(ret1, ret[len(ret)-1])
	var rows = []*proto.AssetData{}
	for _, v := range ret {
		rows = append(rows, &proto.AssetData{
			AssetId:     v.Param.AssetId,
			Username:    v.Param.Info.UserName,
			AssetName:   v.Param.Info.AssetName,
			AssetType:   v.Param.Info.AssetType,
			FeatureTag:  v.Param.Info.FeatureTag,
			SampleHash:  v.Param.Info.SampleHash,
			StorageHash: v.Param.Info.StorageHash,
			ExpireTime:  v.Param.Info.ExpireTime,
			Price:       v.Param.Info.Price,
			OpType:      v.Param.Info.OpType,
			Description: v.Param.Info.Description,
			CreateTime:  uint64(v.CreateTime.Unix()),
		})
	}
	return rows, nil
}

/*func (u *Asset) GetUserPurchaseAssetList(ctx context.Context, req *proto.GetUserPurchaseAssetListRequest, rsp *proto.QueryResponse) error {

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
	log.Info(req.Username)
	if req.Username != "" && req.AssetId != "" {
		where = &bson.M{"type": "datapurchase", "data.basic_info.user_name": req.Username, "data.basic_info.asset_id": req.AssetId}
	} else if req.Username != "" {
		where = &bson.M{"type": "datapurchase", "data.basic_info.user_name": req.Username}
	} else if req.AssetId != "" {
		where = &bson.M{"type": "datapurchase", "data.basic_info.asset_id": req.AssetId}
	}

	log.Info("where:", where)

	var ret []bean.PurchaseMesssageBean

	var mgo = mgo.Session()
	defer mgo.Close()
	count, err := mgo.DB(config.DB_NAME).C("Messages").Find(where).Count()
	if err != nil {
		log.Error(err)
	}
	mgo.DB(config.DB_NAME).C("Messages").Find(where).Skip(skip).Limit(pageSize).All(&ret)
	//mgo.DB(config.DB_NAME).C("Messages").Find(where).Sort("data.basic_info.publish_date").Skip(skip).Limit(int(req.PageSize)).All(&ret)

	var rows = []*proto.QueryRow{}
	for _, v := range ret {
		asset, err := GetAssetById(v.Data.BasicInfo.AssetID, v.Data.BasicInfo.UserName)
		if err != nil {
			fmt.Println(err)
			return errors.New("failed CallGetAssetById " + v.Data.BasicInfo.AssetID)
		}
		rows = append(rows, asset[0])
		//rows = append(rows, &proto.QueryPurchaseRow{
		//	Username: v.Data.BasicInfo.UserName,
		//	AssetId:  v.Data.BasicInfo.AssetID,
		//SampleHash:  v.Data.BasicInfo.SampleHash,
		//StoragePath:  v.Data.BasicInfo.StoragePath,

	}

	//for i := 0; i < len(purMsgs); i++ {
	//asset, err := r.CallGetAssetById(purMsgs[i].Data.BasicInfo.AssetID)
	//fmt.Println(purMsgs[i].Data.BasicInfo.AssetID)
	//if err != nil {
	//fmt.Println(err)
	//return nil, errors.New("failed CallGetAssetById " + purMsgs[i].Data.BasicInfo.AssetID)
	//}
	//tfxs = append(tfxs, asset)
	//}

	var data = &proto.QueryData{
		RowCount: uint64(count),
		PageNum:  uint64(pageNum),
		Row:      rows,
	}
	log.Info(data)
	rsp.Code = 0
	rsp.Data = data
	rsp.Msg = "OK"

	return nil
}*/

/*func (u *Asset) GetUserPurchaseAssetList(ctx context.Context, req *proto.GetUserPurchaseAssetListRequest, rsp *proto.GetUserPurchaseAssetListResponse) error {
	start_time := time.Now().UnixNano() / int64(time.Millisecond)
	dataBody, signValue, account, data := "", "", "", ""
	//dataBody, signValue, account, data := GetSignAndDataCom(req.PostBody)
	log.Info(account, data)
	//get Public Key
	pubKey := GetPublicKey("account")
	//Verify Sign Local
	ok, _ := VerifySign(dataBody, signValue, pubKey)
	log.Info(ok)
	ok = true
	if !ok {
		rsp.Code = 2000
		rsp.Msg = "Verify Signature Failed."
		return nil
	}
	//Test
	params := `service=storage&method=Storage.GetUserPurchaseAssetList&request={
	"username":"%s"
	}`
	userName := req.Username
	//random := req.Random

	//signature := req.Signature

	s := fmt.Sprintf(params, userName)
	log.Info("s:", s)
	resp, err := http.Post(STORAGE_RPC_URL, "application/x-www-form-urlencoded",
		strings.NewReader(s))

	log.Info("resp:", resp)
	//log.Info("err", err)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	} else {
		js, _ := simplejson.NewJson([]byte(body))
		log.Info("jss", js)
		result, _ := json.Marshal(js.Get("UserAssetList"))
		if js.Get("code").MustInt() == 1 {

			rsp.Code = 1
			rsp.Msg = "Get User Purchase Asset Successful!"
			rsp.Data = string(result)
			log.Info(result)
			log.Info(string(result))
		}
		return nil
	}
	end_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Info("Time:", end_time-start_time)
	return nil
}*/

// PreSaleNotice On Chain
func (u *Asset) PreSaleNotice(ctx context.Context, req *proto.PushTxRequest, rsp *proto.PreSaleNoticeResponse) error {

	i, err := data.PushTransaction(req)
	if err != nil {
		log.Error(err)
		rsp.Code = 2031
		rsp.Msg = err.Error()
	}
	log.Info(i)
	return nil
}

// QueryMyNotice My received Notice
func (u *Asset) QueryMyNotice(ctx context.Context, req *proto.QueryMyNoticeRequest, rsp *proto.QueryMyNoticeResponse) error {

	var pageNum, pageSize, skip int = 1, 20, 0
	if req.PageNum > 0 {
		pageNum = int(req.PageNum)
	}

	if req.PageSize > 0 && req.PageSize < 20 {
		pageSize = int(req.PageSize)
	}

	skip = (pageNum - 1) * pageSize

	var where interface{}
	//where = bson.M{"param.info.optype": bson.M{"$in": []int32{1,2}}}
	if len(req.Username) > 0 {
		where = &bson.M{"param.info.consumer": req.Username}
	}

	log.Info(where)

	var ret []bean.PreSaleBean

	var mgo = mgo.Session()
	defer mgo.Close()
	count, err := mgo.DB(config.DB_NAME).C("pre_presale").Find(where).Count()
	log.Info(count)
	if err != nil {
		log.Error(err)
	}
	mgo.DB(config.DB_NAME).C("pre_presale").Find(where).Sort("-_id").Skip(skip).Limit(pageSize).All(&ret)

	var rows = []*proto.QueryNoticeRow{}
	for _, v := range ret {
		var ret2 bean.AssetBean
		mgo.DB(config.DB_NAME).C("pre_assetreg").Find(bson.M{"param.info.optype": bson.M{"$in": []int32{1, 2}}, "param.assetid": v.Param.Info.Assetid}).One(&ret2)
		var ret3 bean.Requirement
		mgo.DB(config.DB_NAME).C("pre_datareqreg").Find(bson.M{"param.info.optype": bson.M{"$in": []int32{1, 2}}, "param.datareqid": v.Param.Info.Datareqid}).One(&ret3)

		rows = append(rows, &proto.QueryNoticeRow{
			NoticeId:    v.Param.Datapresaleid,
			Username:    v.Param.Info.Username,
			AssetId:     v.Param.Info.Assetid,
			AssetName:   ret2.Param.Info.AssetName,
			DataReqId:   v.Param.Info.Datareqid,
			DataReqName: ret3.Param.Info.Reqname,
			Consumer:    v.Param.Info.Consumer,
			IsRead:      v.Param.Info.IsRead,
			Time:        uint64(v.CreateTime.Unix()),
		})
	}

	var data = &proto.QueryNoticeData{
		RowCount: uint32(count),
		PageNum:  uint32(pageNum),
		Row:      rows,
	}
	log.Info(data)
	rsp.Data = data
	return nil
}

// GetUnreadNoticeNum Get Unread Notice Num number
func (u *Asset) GetUnreadNoticeNum(ctx context.Context, req *proto.GetUnreadNoticeNumRequest, rsp *proto.GetUnreadNoticeNumResponse) error {
	var where interface{}
	//where = bson.M{"param.info.optype": bson.M{"$in": []int32{1,2}}}
	if len(req.Username) > 0 {
		where = &bson.M{"param.info.consumer": req.Username, "param.info.isread": bson.M{"$exists": false}}
	}
	log.Info(where)

	var mgo = mgo.Session()
	defer mgo.Close()
	count, err := mgo.DB(config.DB_NAME).C("pre_presale").Find(where).Count()
	log.Info(count)
	if err != nil {
		log.Error(err)
	}

	rsp.Data = uint32(count)
	return nil
}

// QueryMyPreSale Notice
func (u *Asset) QueryMyPreSale(ctx context.Context, req *proto.QueryMyNoticeRequest, rsp *proto.QueryMyNoticeResponse) error {

	var pageNum, pageSize, skip int = 1, 20, 0
	if req.PageNum > 0 {
		pageNum = int(req.PageNum)
	}

	if req.PageSize > 0 && req.PageSize < 20 {
		pageSize = int(req.PageSize)
	}

	skip = (pageNum - 1) * pageSize

	var where interface{}
	//where = bson.M{"param.info.optype": bson.M{"$in": []int32{1,2}}}
	if len(req.Username) > 0 {
		where = &bson.M{"param.info.username": req.Username}
	}

	log.Info(where)

	var ret []bean.PreSaleBean

	var mgo = mgo.Session()
	defer mgo.Close()
	count, err := mgo.DB(config.DB_NAME).C("pre_presale").Find(where).Count()
	log.Info(count)
	if err != nil {
		log.Error(err)
	}
	mgo.DB(config.DB_NAME).C("pre_presale").Find(where).Sort("-_id").Skip(skip).Limit(pageSize).All(&ret)

	var rows = []*proto.QueryNoticeRow{}
	for _, v := range ret {
		var ret2 bean.AssetBean
		mgo.DB(config.DB_NAME).C("pre_assetreg").Find(bson.M{"param.info.optype": bson.M{"$in": []int32{1, 2}}, "param.assetid": v.Param.Info.Assetid}).One(&ret2)
		var ret3 bean.Requirement
		mgo.DB(config.DB_NAME).C("pre_datareqreg").Find(bson.M{"param.info.optype": bson.M{"$in": []int32{1, 2}}, "param.datareqid": v.Param.Info.Datareqid}).One(&ret3)

		rows = append(rows, &proto.QueryNoticeRow{
			NoticeId:    v.Param.Datapresaleid,
			Username:    v.Param.Info.Username,
			AssetId:     v.Param.Info.Assetid,
			AssetName:   ret2.Param.Info.AssetName,
			DataReqId:   v.Param.Info.Datareqid,
			DataReqName: ret3.Param.Info.Reqname,
			Consumer:    v.Param.Info.Consumer,
			Time:        uint64(v.CreateTime.Unix()),
		})
	}

	var data = &proto.QueryNoticeData{
		RowCount: uint32(count),
		PageNum:  uint32(pageNum),
		Row:      rows,
	}
	log.Info(data)
	rsp.Data = data
	return nil
}

// ModifyMyNoticeStatus On DB
func (u *Asset) ModifyMyNoticeStatus(ctx context.Context, req *proto.ModifyMyNoticeStatusRequest, rsp *proto.ModifyMyNoticeStatusResponse) error {
	var where interface{}
	//where = bson.M{"param.info.optype": bson.M{"$in": []int32{1,2}}}
	log.Info(req.NoticeId)
	if len(req.Username) > 0 {
		where = &bson.M{"param.info.consumer": req.Username, "param.datapresaleid": req.NoticeId}
	}

	log.Info(where)

	var ret []bean.PreSaleBean

	var mgo = mgo.Session()
	defer mgo.Close()
	count, err := mgo.DB(config.DB_NAME).C("pre_presale").Find(where).Count()
	log.Info(count)
	if err != nil {
		log.Error(err)
	}
	//mgo.DB(config.DB_NAME).C("pre_presale").Find(where).Sort("-_id").Skip(skip).Limit(pageSize).All(&ret)
	mgo.DB(config.DB_NAME).C("pre_presale").Update(where, bson.M{"$set": bson.M{"param.info.isread": 1}})
	//Update(bson.M{"_id": bson.ObjectIdHex("5204af979955496907000001")}, bson.M{"$pull": bson.M{ "interests": "Golang", }})

	var rows = []*proto.QueryNoticeRow{}
	for _, v := range ret {
		var ret2 bean.AssetBean
		mgo.DB(config.DB_NAME).C("pre_assetreg").Find(bson.M{"param.info.optype": bson.M{"$in": []int32{1, 2}}, "param.assetid": v.Param.Info.Assetid}).One(&ret2)
		var ret3 bean.Requirement
		mgo.DB(config.DB_NAME).C("pre_datareqreg").Find(bson.M{"param.info.optype": bson.M{"$in": []int32{1, 2}}, "param.datareqid": v.Param.Info.Datareqid}).One(&ret3)

		rows = append(rows, &proto.QueryNoticeRow{
			NoticeId:    v.Param.Datapresaleid,
			Username:    v.Param.Info.Username,
			AssetId:     v.Param.Info.Assetid,
			AssetName:   ret2.Param.Info.AssetName,
			DataReqId:   v.Param.Info.Datareqid,
			DataReqName: ret3.Param.Info.Reqname,
			Consumer:    v.Param.Info.Consumer,
			Time:        uint64(v.CreateTime.Unix()),
		})
	}

	var data = &proto.QueryNoticeData{
		RowCount: uint32(count),
		//PageNum:  uint32(pageNum),
		Row: rows,
	}
	log.Info(data)
	//rsp.Data = data
	return nil
}

func init() {
	defer log.Flush()
	logger, err := log.LoggerFromConfigAsFile("./config/ass-log.xml")
	if err != nil {
		log.Critical("err parsing config log file", err)
		os.Exit(1)
		return
	}
	log.ReplaceLogger(logger)
}
func main() {
	log.Info("Asset Service Start")

	service := micro.NewService(
		micro.Name("go.micro.srv.v3.asset"),
		micro.Version("3.0.0"),
	)

	service.Init()

	proto.RegisterAssetHandler(service.Server(), new(Asset))

	if err := service.Run(); err != nil {
		log.Critical("Asset Service Run Failed", err)
	}
}
