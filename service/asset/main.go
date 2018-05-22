package main

import (
	"github.com/micro/go-micro"
	proto "github.com/bottos-project/bottos/service/asset/proto"
	"golang.org/x/net/context"
	"github.com/bitly/go-simplejson"
	"time"
	storage "github.com/bottos-project/bottos/service/storage/proto"
	"github.com/micro/go-micro/client"
	"io/ioutil"
	"net/http"
	"strings"
	"fmt"
	"strconv"
	"encoding/json"
	"github.com/bottos-project/bottos/config"
	"gopkg.in/mgo.v2/bson"
	"github.com/bottos-project/bottos/service/common/bean"
	"github.com/bottos-project/bottos/tools/db/mongodb"
	log "github.com/cihub/seelog"
	"os"
	"github.com/bottos-project/bottos/service/common/data"
	//"github.com/mikemintang/go-curl"
	"github.com/mikemintang/go-curl"
)

const (
	BASE_URL                = config.BASE_CHAIN_URL
	GET_INFO_URL            = BASE_URL + "v1/chain/get_info"
	GET_BLOCK_URL           = BASE_URL + "v1/chain/get_block"
	ABI_JSON_TO_BIN_URL     = BASE_URL + "v1/chain/abi_json_to_bin"
	PUSH_TRANSACTION_URL    = BASE_URL + "v1/chain/push_transaction"
	GET_TABLE_ROW_BY_STRING = BASE_URL + "v1/chain/get_table_row_by_string_key"
	STORAGE_RPC_URL         = config.BASE_RPC
)

type Asset struct{}

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
func GetSignAndData(postBody string) (string, string, string) {
	js, _ := simplejson.NewJson([]byte(postBody))
	//get signed data
	//TODO
	dataBody := js.Get("signatures").MustString()
	log.Info("dataBody", dataBody)
	//getSignValue
	signValue := js.Get("signatures").MustString()
	log.Info(signValue)
	//get username
	userName := js.Get("userName").MustString()

	//messages := js.Get("messages").GetIndex(0)
	//authorization := messages.Get("authorization").GetIndex(0)
	//log.Info("----------", authorization.Get("account").MustString())

	//postData := map[string]interface{}{
	//	"ref_block_num": js.Get("ref_block_num").MustInt(),
	//}
	return dataBody, signValue, userName
}

func (u *Asset) RegisterFile(ctx context.Context, req *proto.RegisterFileRequest, rsp *proto.RegisterFileResponse) error {
	start_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Info("reqBody:" + req.PostBody)

	rsp.Code = 1005
	//var requestStruct sign_proto.Transaction
	//json.Unmarshal([]byte(req.PostBody), &requestStruct)

	ret, err := data.PushTransaction(req.PostBody)
	if err != nil {
		rsp.Msg = err.Error()
		return nil
	}
	log.Info("ret-file:", ret)
	log.Info(ret.Result.TrxHash)

	//Check the chain for packaging results.
	params := `service=core&method=CoreApi.QueryObject&request={
	"contract":"%s",
	"object":"%s",
	"key":"%s"
	}`
	s := fmt.Sprintf(params, "datafilemng", "datafilereg", ret.Result.TrxHash)
	resp, err := http.Post(BASE_URL, "application/x-www-form-urlencoded",
		strings.NewReader(s))

	log.Info("resp:", resp)
	log.Info("err", err)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Info(body)
	//test
	rsp.Code = 0
	end_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Info("Time:", end_time-start_time)

	return nil
}

func (u *Asset) RegisterAsset(ctx context.Context, req *proto.RegisterRequest, rsp *proto.RegisterResponse) error {
	start_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Info("reqBody:" + req.PostBody)

	rsp.Code = 2001
	//var requestStruct sign_proto.Transaction
	//json.Unmarshal([]byte(req.PostBody), &requestStruct)

	ret, err := data.PushTransaction(req.PostBody)
	if err != nil {
		rsp.Msg = err.Error()
		return nil
	}
	log.Info("ret-file:", ret)
	log.Info(ret.Result.TrxHash)

	//Check the chain for packaging results.
	//params := `service=core&method=CoreApi.QueryTrx&request={
	//"contract":"%s",
	//"object":"%s",
	//"key":"%s"
	//}`
	txHash := "c76387dc13f54c1fb73ec3903a14ac4c006d5be23565fd660aaeb5df124cccce"

	params := `service=core&method=CoreApi.QueryTrx&request={
	"trx_hash":"%s"
	}`
	s := fmt.Sprintf(params, txHash)
	//s := fmt.Sprintf(params, ret.Result.TrxHash)
	resp, err := http.Post(STORAGE_RPC_URL, "application/x-www-form-urlencoded",
		strings.NewReader(s))

	log.Info("resp:", resp)
	log.Info("err", err)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Info(body)
	if err != nil {
		return err
	} else {
		jss, _ := simplejson.NewJson([]byte(body))
		log.Info("jss", jss)
		errcode := jss.Get("errcode").MustInt64()
		log.Info(errcode)
		if errcode == 0 {
		}
	}
	//test
	rsp.Code = 0
	end_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Info("Time:", end_time-start_time)

	return nil
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

func GetAssetList(queryPara *proto.QueryPara) (string, int, error) {
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

func (u *Asset) QueryUploadedData(ctx context.Context, req *proto.QueryRequest, rsp *proto.QueryUploadedDataResponse) error {
	var pageNum, pageSize, skip int = 1, 20, 0
	if req.PageNum > 0 {
		pageNum = int(req.PageNum)
	}

	if req.PageSize > 0 {
		pageSize = int(req.PageSize)
	}

	skip = (pageNum - 1) * pageSize

	var where interface{}
	where = bson.M{"param.info.optype": bson.M{"$in": []int32{1, 2}}}
	//if len(req.Username) > 0{
	//	where = &bson.M{"param.info.optype": bson.M{"$in": []uint32{1,2}}, "param.info.username": req.Username}
	//}

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

/*func (u *Asset) QueryUploadedData(ctx context.Context, req *proto.QueryUploadedData, rsp *proto.QueryUploadedDataResponse) error {
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
	params := `service=storage&method=Storage.GetUserFileList&request={
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
		result, _ := json.Marshal(js.Get("FileList"))
		if js.Get("code").MustInt() == 1 {

			rsp.Code = 1
			rsp.Msg = "Get File List Successful!"
			rsp.Data = string(result)
		}
		return nil
	}
	end_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Info("Time:", end_time-start_time)
	return nil
}*/

func (u *Asset) GetDownLoadURL(ctx context.Context, req *proto.GetDownLoadURLRequest, rsp *proto.GetDownLoadURLResponse) error {
	start_time := time.Now().UnixNano() / int64(time.Millisecond)

	//Test
	params := `service=storage&method=Storage.GetDownLoadURL&request={
	"username":"%s",
	"file_name":"%s"
	}`
	userName := req.Username
	fileName := req.FileName

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
	} else {
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
	end_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Info("Time:", end_time-start_time)
	return nil
}

/*func (u *Asset) Query(ctx context.Context, req *proto.QueryRequest, rsp *proto.QueryResponse) error {

	var pageNum, pageSize, skip int = 1, 20, 0
	if req.PageNum > 0 {
		pageNum = int(req.PageNum)
	}

	if req.PageSize > 0 && req.PageSize <= 50 {
		pageSize = int(req.PageSize)
	}

	skip = (pageNum - 1) * pageSize

	var where interface{}
	where = &bson.M{"type": "assetreg"}
	log.Info(req.Username)
	if req.Username != "" {
		where = &bson.M{"type": "assetreg", "data.basic_info.user_name": req.Username}
		//where = &bson.M{"type": "assetreg", "data.basic_info.user_name": req.Username, "data.basic_info.feature_tag": req.FeatureTag}
	} else {
		//if req.Username != "" {
		where = &bson.M{"type": "assetreg"}
		//}

	}
	log.Info(skip)
	log.Info("where:", where)

	var ret []bean.AssetBean
	//var ret1 []bean.AssetBean

	var mgo = mgo.Session()
	defer mgo.Close()
	count, err := mgo.DB(config.DB_NAME).C("Messages").Find(where).Count()
	if err != nil {
		log.Error(err)
	}

	//mgo.DB(config.DB_NAME).C("Messages").Find(where).Sort("data.asset_id").Skip(skip).Limit(pageSize).All(&ret)
	//mgo.DB(config.DB_NAME).C("Messages").Find(where).Sort("-createdAt").Skip(skip).Limit(int(req.PageSize)).Distinct("data.asset_id",&ret)
	mgo.DB(config.DB_NAME).C("Messages").Find(where).Sort("-createdAt").Skip(skip).Limit(pageSize).All(&ret)
	log.Info("ret:", ret)

	*//*	Remove Duplicates
		a_len := len(ret) - 1
		log.Info(a_len)
		if a_len == 0 {
			ret1 = append(ret1, ret[a_len])
		} else {
			for i := a_len; i >= 0; i-- {
				if (i < a_len && ret[i+1].Data.AssetID == ret[i].Data.AssetID) || len(ret) == 0 {
					continue
				}
				ret1 = append(ret1, ret[i])
			}
		}

		log.Info("ret1:", ret1)*//*

	var rows = []*proto.QueryRow{}
	for _, v := range ret {
		rows = append(rows, &proto.QueryRow{
			AssetId:     v.Data.AssetID,
			Username:    v.Data.BasicInfo.UserName,
			AssetName:   v.Data.BasicInfo.AssetName,
			AssetType:   v.Data.BasicInfo.AssetType,
			FeatureTag1: v.Data.BasicInfo.FeatureTag1,
			FeatureTag2: v.Data.BasicInfo.FeatureTag2,
			FeatureTag3: v.Data.BasicInfo.FeatureTag3,
			SamplePath:  v.Data.BasicInfo.SamplePath,
			SampleHash:  v.Data.BasicInfo.SampleHash,
			StoragePath: "",
			StorageHash: "",
			//SampleHash:  v.Data.BasicInfo.SampleHash,
			//StoragePath: v.Data.BasicInfo.StoragePath,
			ExpireTime:  v.Data.BasicInfo.ExpireTime,
			Price:       v.Data.BasicInfo.Price,
			Description: v.Data.BasicInfo.Description,
			UploadDate:  v.Data.BasicInfo.UploadDate,
			CreateTime:  v.CreatedAt.String(),
		})
	}

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
func (u *Asset) QueryAsset(ctx context.Context, req *proto.QueryRequest, rsp *proto.QueryResponse) error {

	var pageNum, pageSize, skip int = 1, 20, 0
	if req.PageNum > 0 {
		pageNum = int(req.PageNum)
	}

	if req.PageSize > 0 {
		pageSize = int(req.PageSize)
	}

	skip = (pageNum - 1) * pageSize

	var where interface{}
	where = bson.M{"param.info.optype": bson.M{"$in": []int32{1, 2}}}
	//if len(req.Username) > 0{
	//	where = &bson.M{"param.info.optype": bson.M{"$in": []uint32{1,2}}, "param.info.username": req.Username}
	//}

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
			Price:       v.Param.Info.Price,
			OpType:      v.Param.Info.OpType,
			Description: v.Param.Info.Description,
			CreateTime:  uint64(v.CreateTime.Unix()),
		})
	}

	var data = &proto.QueryData{
		RowCount: uint32(count),
		PageNum:  uint32(pageNum),
		Row:      rows,
	}
	log.Info(data)
	rsp.Data = data
	return nil
}

func (u *Asset) QueryAllAsset(ctx context.Context, req *proto.QueryAllAssetRequest, rsp *proto.QueryAllAssetResponse) error {
	start_time := time.Now().UnixNano() / int64(time.Millisecond)

	//Test
	params := `service=storage&method=Storage.GetAllAssetList&request={
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
		result, _ := json.Marshal(js.Get("AssetList"))
		if js.Get("code").MustInt() == 1 {

			rsp.Code = 1
			rsp.Msg = "Get All Asset List Successful!"
			rsp.Data = string(result)
			log.Info(result)
			log.Info(string(result))
		}
		return nil
	}
	end_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Info("Time:", end_time-start_time)
	return nil
}

//func (u *Asset) Modify(ctx context.Context, req *proto.ModifyRequest, rsp *proto.ModifyResponse) error {
//	start_time := time.Now().UnixNano() / int64(time.Millisecond)
//	log.Info("reqBody:" + req.PostBody)
//
//	//Write to BlockChain
//	flag, result := cbb.WriteToBlockChain(req.PostBody, PUSH_TRANSACTION_URL)
//	log.Info("OK1,", result)
//
//	end_time := time.Now().UnixNano() / int64(time.Millisecond)
//	log.Info("Time:", end_time-start_time)
//	//ok1 = true
//	if flag == false {
//		rsp.Code = 2000
//		rsp.Msg = "Modify Asset Failed."
//		rsp.Data = result
//		return nil
//	} else {
//		rsp.Code = 1
//		rsp.Msg = "Modify Asset Successful!"
//		rsp.Data = string(result)
//		log.Info(string(result))
//		return nil
//	}
//
//	/*	获取data JSON格式数据
//	postData1 := map[string]interface{}{
//			"code":    "bto",
//			"action":  "newaccount",
//			"binargs": data,
//		}
//		reqBinToJson := curl.NewRequest()
//		resp, err := reqBinToJson.SetUrl(ABI_BIN_TO_JSON_URL).SetPostData(postData1).Post()
//
//		if err != nil {
//			return nil
//		}*/
//}

/*func (u *Asset) QueryByID(ctx context.Context, req *proto.QueryByIDRequest, rsp *proto.QueryResponse) error {

	var pageNum, pageSize, skip int = 1, 20, 0
	if req.PageNum > 0 {
		pageNum = int(req.PageNum)
	}

	if req.PageSize > 0 && req.PageSize <= 50 {
		pageSize = int(req.PageSize)
	}

	skip = (pageNum - 1) * pageSize
	log.Debug(skip)
	*//*var where interface{}
	where = &bson.M{"type": "assetreg"}
	log.Info(req.AssetID)
	if req.AssetID != "" {
		where = &bson.M{"type": "assetreg", "data.asset_id": req.AssetID}
	} else {
		if req.AssetID != "" {
			where = &bson.M{"type": "assetreg", "data.asset_id": req.AssetID}
		}
	}

	log.Info(where)

	var ret []bean.AssetBean

	var mgo = mgo.Session()
	defer mgo.Close()
	count, err := mgo.DB(config.DB_NAME).C("Messages").Find(where).Count()
	if err != nil {
		log.Error(err)
	}
	mgo.DB(config.DB_NAME).C("Messages").Find(where).Skip(skip).Limit(int(req.PageSize)).All(&ret)
	//mgo.DB(config.DB_NAME).C("Messages").Find(where).Sort("data.basic_info.publish_date").Skip(skip).Limit(int(req.PageSize)).All(&ret)

	var rows = []*proto.QueryRow{}
	for _, v := range ret {
		rows = append(rows, &proto.QueryRow{
			AssetId:   v.Data.AssetID,
			Username:  v.Data.BasicInfo.UserName,
			AssetName: v.Data.BasicInfo.AssetName,
			AssetType: v.Data.BasicInfo.AssetType,
			FeatureTag1: v.Data.BasicInfo.FeatureTag1,
			FeatureTag2: v.Data.BasicInfo.FeatureTag2,
			FeatureTag3: v.Data.BasicInfo.FeatureTag3,
			SamplePath:  v.Data.BasicInfo.SamplePath,
			SampleHash:  v.Data.BasicInfo.SampleHash,
			StoragePath: v.Data.BasicInfo.StoragePath,
			StorageHash: v.Data.BasicInfo.StorageHash,
			ExpireTime:  v.Data.BasicInfo.ExpireTime,
			Price:       v.Data.BasicInfo.Price,
			Description: v.Data.BasicInfo.Description,
			UploadDate:  v.Data.BasicInfo.UploadDate,
		})
	}*//*
	rows, err := GetAssetByIdNoStoPath(req.AssetID)
	if err != nil {
		fmt.Println(err)
		return errors.New("Get session faild" + req.AssetID)
	}
	var data = &proto.QueryData{
		RowCount: 1,
		//RowCount: uint64(count),
		PageNum: uint64(pageNum),
		Row:     rows,
	}
	log.Info("rows:", rows)
	log.Info(data)
	rsp.Code = 0
	rsp.Data = data
	rsp.Msg = "OK"

	return nil
}*/
/*func GetAssetById(assertId string, userName string) ([]*proto.QueryData, error) {
	var where interface{}

	//where = &bson.M{"type": "assetreg"}
	log.Info(assertId)
	if assertId != "" {

		//where = &bson.M{"type": "assetreg", "data.asset_id": assertId}
		where = bson.M{"param.info.optype": bson.M{"$in": []int32{1, 2}}, "param.assetid": assertId}
	} else {
		return nil, nil
	}

	log.Info(where)

	var ret []bean.AssetBean
	var ret1 []bean.AssetBean

	var mgo = mgo.Session()
	defer mgo.Close()

	//mgo.DB(config.DB_NAME).C("pre_assetreg").Find(where).Sort("createdAt").All(&ret)
	mgo.DB(config.DB_NAME).C("pre_assetreg").Find(where).Sort("-_id").All(&ret)

	ret1 = append(ret1, ret[len(ret)-1])

	var rows = []*proto.AssetData{}
	for _, v := range ret1 {
		result, _, _ := GetTableRowByString(userName, v.Data.BasicInfo.StorageHash)
		log.Info("GetTableRowByString:", result)
		if strings.Contains(result, v.Data.BasicInfo.StorageHash) {
			rows = append(rows, &proto.QueryRow{
				AssetId:     v.Data.AssetID,
				Username:    v.Data.BasicInfo.UserName,
				AssetName:   v.Data.BasicInfo.AssetName,
				AssetType:   v.Data.BasicInfo.AssetType,
				FeatureTag1: v.Data.BasicInfo.FeatureTag1,
				FeatureTag2: v.Data.BasicInfo.FeatureTag2,
				FeatureTag3: v.Data.BasicInfo.FeatureTag3,
				SamplePath:  v.Data.BasicInfo.SamplePath,
				SampleHash:  v.Data.BasicInfo.SampleHash,
				StoragePath: v.Data.BasicInfo.StoragePath,
				StorageHash: v.Data.BasicInfo.StorageHash,
				ExpireTime:  v.Data.BasicInfo.ExpireTime,
				Price:       v.Data.BasicInfo.Price,
				Description: v.Data.BasicInfo.Description,
				UploadDate:  v.Data.BasicInfo.UploadDate,
				CreateTime:  v.CreatedAt.String(),
			})
		} else {
			rows = append(rows, &proto.QueryRow{
				AssetId:     v.Data.AssetID,
				Username:    v.Data.BasicInfo.UserName,
				AssetName:   v.Data.BasicInfo.AssetName,
				AssetType:   v.Data.BasicInfo.AssetType,
				FeatureTag1: v.Data.BasicInfo.FeatureTag1,
				FeatureTag2: v.Data.BasicInfo.FeatureTag2,
				FeatureTag3: v.Data.BasicInfo.FeatureTag3,
				SamplePath:  v.Data.BasicInfo.SamplePath,
				SampleHash:  v.Data.BasicInfo.SampleHash,
				StoragePath: "",
				StorageHash: "",
				ExpireTime:  v.Data.BasicInfo.ExpireTime,
				Price:       v.Data.BasicInfo.Price,
				Description: v.Data.BasicInfo.Description,
				UploadDate:  v.Data.BasicInfo.UploadDate,
				CreateTime:  v.CreatedAt.String(),
			})
		}

	}
	return rows, nil
}*/
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
	} else {
		return "", resp.Raw.StatusCode, err
	}
}

func GetAssetByIdNoStoPath(assertId string) ([]*proto.AssetData, error) {
	var where interface{}
	//where = &bson.M{"type": "assetreg"}
	log.Info(assertId)
	if assertId != "" {
		where = bson.M{"param.info.optype": bson.M{"$in": []int32{1, 2}}, "param.assetid": assertId}
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
	if count>1 {
		log.Error(count)
		return nil,nil
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

/*func (u *Asset) QueryByID(ctx context.Context, req *proto.QueryByIDRequest, rsp *proto.QueryByIDResponse) error {
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
	params := `service=storage&method=Storage.GetAssetByAssetId&request={
	"asset_id":"%s"
	}`
	assetID := req.AssetID
	//random := req.Random

	//signature := req.Signature

	s := fmt.Sprintf(params, assetID)
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
		result, _ := json.Marshal(js.Get("asset_info"))
		if js.Get("code").MustInt() == 1 {

			rsp.Code = 1
			rsp.Msg = "Get Asset by ID Successful!"
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

func (u *Asset) PreSaleNotice(ctx context.Context, req *proto.PushTxRequest, rsp *proto.PreSaleNoticeResponse) error {

	i, err := data.PushTransaction(req)

	if i != nil {
		rsp.Code = 1008
		rsp.Msg = err.Error()
	}
	return nil
}

func (u *Asset) QueryMyNotice(ctx context.Context, req *proto.QueryMyNoticeRequest, rsp *proto.QueryMyNoticeResponse) error {

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

	var rows = []*proto.QueryNoticeRow{}
	for _, v := range ret {

		rows = append(rows, &proto.QueryNoticeRow{
			NoticeId : v.Param.Datapresaleid,
			Username : v.Param.Info.Username,
			AssetId : v.Param.Info.Assetid,
			//AssetName : v.Param.Info.ass,
			DataReqId : v.Param.Info.Datareqid,
			Consumer : v.Param.Info.Consumer,
			Time : uint64(v.CreateTime.Unix()),
		})
	}

	var data = &proto.QueryNoticeData{
		RowCount: uint32(count),
		PageNum: uint32(pageNum),
		Row:rows,
	}
	log.Info(data)
	rsp.Data = data
	return nil
}

func (u *Asset) QueryMyPreSale(ctx context.Context, req *proto.QueryMyNoticeRequest, rsp *proto.QueryMyNoticeResponse) error {

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

	var rows = []*proto.QueryNoticeRow{}
	for _, v := range ret {

		rows = append(rows, &proto.QueryNoticeRow{
			NoticeId : v.Param.Datapresaleid,
			Username : v.Param.Info.Username,
			AssetId : v.Param.Info.Assetid,
			//AssetName : v.Param.Info.ass,
			DataReqId : v.Param.Info.Datareqid,
			Consumer : v.Param.Info.Consumer,
			Time : uint64(v.CreateTime.Unix()),
		})
	}

	var data = &proto.QueryNoticeData{
		RowCount: uint32(count),
		PageNum: uint32(pageNum),
		Row:rows,
	}
	log.Info(data)
	rsp.Data = data
	return nil
}

func init() {
	defer log.Flush()
	logger, err := log.LoggerFromConfigAsFile("./config/log.xml")
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
