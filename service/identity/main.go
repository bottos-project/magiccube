package main

import (
	log "github.com/cihub/seelog"
	"github.com/micro/go-micro"
	user_proto "github.com/bottos-project/bottos/service/identity/proto"
	"golang.org/x/net/context"
	"github.com/mikemintang/go-curl"
	"github.com/bitly/go-simplejson"
	"time"
	"encoding/base64"
	"github.com/satori/go.uuid"
	"strings"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"bytes"
	"fmt"
	cbb "github.com/bottos-project/bottos/service/asset/cbb"
	"strconv"
	"github.com/bottos-project/bottos/tools/db/mongodb"
	"gopkg.in/mgo.v2/bson"
	"github.com/bottos-project/bottos/service/bean"
	"github.com/bottos-project/bottos/config"
	"sort"
)

const (
	BASE_URL                = config.BASE_CHAIN_URL
	GET_INFO_URL            = BASE_URL + "v1/chain/get_info"
	GET_BLOCK_URL           = BASE_URL + "v1/chain/get_block"
	ABI_JSON_TO_BIN_URL     = BASE_URL + "v1/chain/abi_json_to_bin"
	PUSH_TRANSACTION_URL    = BASE_URL + "v1/chain/push_transaction"
	GET_TABLE_ROW_BY_STRING = BASE_URL + "v1/chain/get_table_row_by_string_key"
	GET_TABLE_ROWS          = BASE_URL + "v1/chain/get_table_rows"
	STORAGE_RPC_URL         = config.BASE_RPC
)

var index int = 0

type User struct{}

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
	is_login, account := UserLogin(req.Body)
	log.Info(account)
	if is_login {
		token := create_token()
		is_save_token := save_token(account, token)
		if is_save_token {
			rsp.Code = 0
			rsp.Msg = "OK"
			rsp.Token = token
		} else {
			rsp.Code = 1002
			rsp.Msg = "Write token failure"
		}
	} else {
		rsp.Code = 1001
		rsp.Msg = "Access to account information failure"
	}
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

func (u *User) VerifyToken(ctx context.Context, req *user_proto.VerifyTokenRequest, rsp *user_proto.VerifyTokenResponse) error {
	log.Info(req.Token)
	if req.Token == "" {
		rsp.Code = 1999
		rsp.Msg = "Token is nil"
		return nil
	}

	checkToken := CheckToken(req.Token)
	if checkToken {
		rsp.Code = 0
		rsp.Msg = "OK"
	} else {
		rsp.Code = 1999
		rsp.Msg = "Invalid Token"
	}
	return nil
}

type UserTokenBean struct {
	Username string `bson:"username"`
	Token    string `bson:"token"`
	Ctime    int64  `bson:"ctime"`
}

func CheckToken(token string) bool {
	var mgo = mgo.Session()
	defer mgo.Close()
	var ret UserTokenBean
	err := mgo.DB(config.DB_NAME).C("user_token").Find(&bson.M{"token": token}).One(&ret)
	if err != nil {
		log.Error(err)
		return false
	}

	if ret.Token == token {
		if (ret.Ctime + config.TOKEN_EXPIRE_TIME > time.Now().Unix()) {
			return true
		}
	}
	return false
}

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

func (u *User) GetUserInfo(ctx context.Context, req *user_proto.GetUserInfoRequest, rsp *user_proto.GetUserInfoResponse) error {
	data, _, _ := GetTableRow(req.Username)
	rsp.Code = 0
	rsp.Msg = "OK"
	rsp.Data = data
	return nil
}

func (u *User) UpdateUserInfo(ctx context.Context, req *user_proto.UpdateUserInfoRequest, rsp *user_proto.UpdateUserInfoResponse) error {
	ref_block_num := GetBolckNum()
	log.Info("ref_block_num:", ref_block_num)
	time.Sleep(1)
	if ref_block_num < 0 {
		rsp.Code = 1131
		rsp.Msg = "Data[ref_block_num] exception"
		return nil
	}
	ref_block_prefix, timestamp := GetBlockPrefix(ref_block_num)
	log.Info("ref_block_prefix:", ref_block_prefix)
	log.Info("timestamp:", timestamp)
	time.Sleep(1)
	if ref_block_prefix < 0 {
		rsp.Code = 1132
		rsp.Msg = "Data[ref_block_prefix] exception"
		return nil
	}
	expirationTime := ExpirationTime(timestamp, 20)
	time.Sleep(1)
	b, _ := json.Marshal(req.UserInfo)
	userBin := UserGetBin(req.Username, string(b))
	log.Info("userBin:", userBin)
	time.Sleep(1)
	if userBin == "" {
		rsp.Code = 1104
		rsp.Msg = "Data[user_bin] exception"
		return nil
	}
	userInfo := UserPushTransaction(ref_block_num, ref_block_prefix, expirationTime, userBin, req.Username)
	log.Info(userInfo)
	if userInfo != "" {
		rsp.Code = 0
		rsp.Msg = "modify user success"
	} else {
		rsp.Code = 1103
		rsp.Msg = "modify user failure"
	}
	return nil
}

func (u *User) FavoriteMng(ctx context.Context, req *user_proto.FavoriteMngRequest, rsp *user_proto.FavoriteMngResponse) error {
	start_time := time.Now().UnixNano() / int64(time.Millisecond)

	log.Info("reqBody:" + req.PostBody)
	dataBody, signValue, account, data := cbb.GetSignAndDataCom(req.PostBody)
	log.Info(account, data)
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

	//Write to BlockChain
	flag, result := cbb.WriteToBlockChain(req.PostBody, PUSH_TRANSACTION_URL)
	log.Info("OK1,", result)

	end_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Info("Time:", end_time-start_time)
	//ok1 = true
	if flag == false {
		rsp.Code = 1100
		rsp.Msg = "Add or Delete Favorite Failed."
		return nil
	} else {
		rsp.Code = 1
		rsp.Msg = "Add or Delete Favorite Successful!"
		rsp.Data = string(result)
		log.Info(string(result))
		return nil
	}

	/*	获取data JSON格式数据
		postData1 := map[string]interface{}{
				"code":    "bto",
				"action":  "newaccount",
				"binargs": data,
			}
			reqBinToJson := curl.NewRequest()
			resp, err := reqBinToJson.SetUrl(ABI_BIN_TO_JSON_URL).SetPostData(postData1).Post()

			if err != nil {
				return nil
			}*/

}
func (u *User) QueryFavorite(ctx context.Context, req *user_proto.QueryFavoriteRequest, rsp *user_proto.QueryFavoriteResponse) error {
	dataBody, signValue, account, data := "", "", "", ""
	//dataBody, signValue, account, data := GetSignAndDataCom(req.PostBody)
	log.Info(account, data)
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

	if req.PageSize > 0 && req.PageSize <= 50{
		pageSize = int(req.PageSize)
	}

	skip = (pageNum - 1) * pageSize

	var where interface{}
	where = &bson.M{"type": "favoritepro"}
	log.Info(req.Username)
	if req.Username != "" {
		where = &bson.M{"type": "favoritepro", "data.user_name": req.Username}
	}

	var mgo = mgo.Session()
	defer mgo.Close()
	count, err := mgo.DB(config.DB_NAME).C("Messages").Find(where).Count()
	if err != nil {
		log.Error(err)
	}

	var ret []bean.FavoriteBean
	var ret1 []bean.FavoriteBean
	mgo.DB(config.DB_NAME).C("Messages").Find(where).Sort("data.goods_id").Skip(skip).Limit(pageSize).All(&ret)

	log.Info("ret:", ret)
	a_len := len(ret) - 1
	//for i := a_len; i >0; i-- {
	//	if (i < a_len && ret[i+1].Data.GoodsID == ret[i].Data.GoodsID) || len(ret) == 0 {
	//		continue
	//	}
	//	ret1 = append(ret1, ret[i])
	//}
	for i := 0; i < a_len; i++ {
		if (ret[i].Data.GoodsID == ret[i+1].Data.GoodsID) || len(ret) == 0 {
			i++
			continue
		}
		log.Info("i++：", i)
		ret1 = append(ret1, ret[i])
	}
	if a_len != -1 {
		if a_len == 0 || (ret[a_len-1].Data.GoodsID != ret[a_len].Data.GoodsID) {
			ret1 = append(ret1, ret[a_len])
		}
	}

	log.Info("ret1:", ret1)
	var rows = []*user_proto.FavoriteRowData{}
	for _, v := range ret1 {
		rows = append(rows, &user_proto.FavoriteRowData{
			v.Data.UserName,
			v.Data.OpType,
			v.Data.GoodsType,
			v.Data.GoodsID,
		})
	}

	var d = &user_proto.FavoriteData{
		PageNum:  uint64(pageNum),
		RowCount: uint64(count),
		Row:      rows,
	}

	rsp.Code = 1
	rsp.Data = d
	////Test
	//params := `service=storage&method=Storage.GetUserFavorit&request={
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
	//	result, _ := json.Marshal(js.Get("favorit_list"))
	//	if js.Get("code").MustInt() == 1 {
	//
	//		rsp.Code = 1
	//		rsp.Msg = "Query Favorite List Successful!"
	//		rsp.Data = string(result)
	//	}
	//	return nil
	//}
	return nil
}

func (u *User) AddShopCar(ctx context.Context, req *user_proto.AddShopCarRequest, rsp *user_proto.AddShopCarResponse) error {
	start_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Info("reqBody:" + req.PostBody)
	dataBody, signValue, account, data := cbb.GetSignAndDataCom(req.PostBody)
	log.Info(account, data)
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

	//Write to BlockChain
	flag, result := cbb.WriteToBlockChain(req.PostBody, PUSH_TRANSACTION_URL)
	log.Info("OK1,", result)

	end_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Info("Time:", end_time-start_time)
	//ok1 = true
	if flag == false {
		rsp.Code = 2000
		rsp.Msg = "Register Asset Failed."
		return nil
	} else {
		rsp.Code = 1
		rsp.Msg = "Register Asset Successful!"
		rsp.Data = string(result)
		log.Info(string(result))
		return nil
	}

	/*	获取data JSON格式数据
		postData1 := map[string]interface{}{
				"code":    "bto",
				"action":  "newaccount",
				"binargs": data,
			}
			reqBinToJson := curl.NewRequest()
			resp, err := reqBinToJson.SetUrl(ABI_BIN_TO_JSON_URL).SetPostData(postData1).Post()

			if err != nil {
				return nil
			}*/

}
func (u *User) QueryShopCar(ctx context.Context, req *user_proto.QueryShopCarRequest, rsp *user_proto.QueryShopCarResponse) error {
	dataBody, signValue, account, data := "", "", "", ""
	//dataBody, signValue, account, data := GetSignAndDataCom(req.PostBody)
	log.Info(account, data)
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

	if req.PageSize > 0 && req.PageSize <= 50{
		pageSize = int(req.PageSize)
	}

	skip = (pageNum - 1) * pageSize

	var where interface{}
	where = &bson.M{"type": "datafilereg"}
	log.Info(req.Username)
	if req.Username != "" {
		where = &bson.M{"type": "datafilereg", "data.basic_info.user_name": req.Username}
	}

	var mgo = mgo.Session()
	defer mgo.Close()
	count, err := mgo.DB(config.DB_NAME).C("Messages").Find(where).Count()
	if err != nil {
		log.Error(err)
	}

	var ret []bean.FileBean
	mgo.DB(config.DB_NAME).C("Messages").Find(where).Sort("-createAt").Skip(skip).Limit(pageSize).All(&ret)

	var rows = []*user_proto.ShopCarRow{}
	for _, v := range ret {
		rows = append(rows, &user_proto.ShopCarRow{
			v.Data.FileHash,
			v.Data.BasicInfo.UserName,
			v.Data.BasicInfo.FileName,
			v.Data.BasicInfo.FileSize,
			v.Data.BasicInfo.FileNumber,
			v.Data.BasicInfo.FilePolicy,
			v.Data.BasicInfo.AuthPath,
		})
	}

	var d = &user_proto.ShopCarData{
		PageNum:  uint64(pageNum),
		RowCount: uint64(count),
		Row:      rows,
	}

	rsp.Code = 1;
	rsp.Data = d

	////Test
	//params := `service=storage&method=Storage.GetUserFileList&request={
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
	return nil
}

func (u *User) AddNotice(ctx context.Context, req *user_proto.AddNoticeRequest, rsp *user_proto.AddNoticeResponse) error {
	start_time := time.Now().UnixNano() / int64(time.Millisecond)

	log.Info("reqBody:" + req.PostBody)
	dataBody, signValue, account, data := cbb.GetSignAndDataCom(req.PostBody)
	log.Info(account, data)
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

	//Write to BlockChain
	flag, result := cbb.WriteToBlockChain(req.PostBody, PUSH_TRANSACTION_URL)
	log.Info("OK1,", result)

	end_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Info("Time:", end_time-start_time)
	//ok1 = true
	if flag == false {
		rsp.Code = 1200
		rsp.Msg = "Add or Delete Notification Failed."
		return nil
	} else {
		rsp.Code = 1
		rsp.Msg = "Add or Delete Notification Successful!"
		rsp.Data = string(result)
		log.Info(string(result))
		return nil
	}

	/*	获取data JSON格式数据
		postData1 := map[string]interface{}{
				"code":    "bto",
				"action":  "newaccount",
				"binargs": data,
			}
			reqBinToJson := curl.NewRequest()
			resp, err := reqBinToJson.SetUrl(ABI_BIN_TO_JSON_URL).SetPostData(postData1).Post()

			if err != nil {
				return nil
			}*/

}

func (u *User) QueryNotice(ctx context.Context, req *user_proto.QueryNoticeRequest, rsp *user_proto.QueryNoticeResponse) error {

	var pageNum, pageSize, skip int = 1, 20, 0
	if req.PageNum > 0 {
		pageNum = int(req.PageNum)
	}

	if req.PageSize > 0 && req.PageSize <= 50 {
		pageSize = int(req.PageSize)
	}

	skip = (pageNum - 1) * pageSize

	var where interface{}
	where = &bson.M{"type": "datapresale"}
	log.Info(req.Username)
	if req.Username != "" {
		where = &bson.M{"type": "datapresale", "data.basic_info.user_name": req.Username}
		//where = &bson.M{"type": "assetreg", "data.basic_info.user_name": req.Username, "data.basic_info.feature_tag": req.FeatureTag}
	} else {
		//if req.Username != "" {
		//where = &bson.M{"type": "assetreg"}
		//}
	}
	log.Info(skip)
	log.Info("where:", where)

	var ret []bean.DataPreSaleBean

	var mgo = mgo.Session()
	defer mgo.Close()
	count, err := mgo.DB(config.DB_NAME).C("Messages").Find(where).Count()
	if err != nil {
		log.Error(err)
	}

	mgo.DB(config.DB_NAME).C("Messages").Find(where).Sort("-createdAt").Skip(skip).Limit(pageSize).All(&ret)
	//mgo.DB(config.DB_NAME).C("Messages").Find(where).Sort("-createdAt").Skip(skip).Limit(int(req.PageSize)).Distinct("data.asset_id",&ret)
	log.Info("ret:", ret)

	var rows = []*user_proto.QueryNoticeRow{}
	for _, v := range ret {
		timeLayout := "2006-01-02T15:04:05"
		rows = append(rows, &user_proto.QueryNoticeRow{
			Username:    v.Data.BasicInfo.UserName,
			AssetId:     v.Data.BasicInfo.AssetID,
			AssetName:   v.Data.BasicInfo.AssetName,
			DataReqId:   v.Data.BasicInfo.DataReqID,
			DataReqName: v.Data.BasicInfo.DataReqName,
			Consumer:    v.Data.BasicInfo.Consumer,
			CreateTime:  time.Unix(v.CreatedAt.Unix(), 0).Format(timeLayout),
		})
	}

	var data = &user_proto.QueryNoticeData{
		RowCount: uint64(count),
		PageNum:  uint64(pageNum),
		Row:      rows,
	}
	log.Info(data)
	rsp.Code = 0
	rsp.Data = data
	rsp.Msg = "OK"

	return nil
}

func (u *User) GetAccount(ctx context.Context, req *user_proto.GetAccountRequest, rsp *user_proto.GetAccountResponse) error {
	log.Info("req:", req)
	start_time := time.Now().UnixNano() / int64(time.Millisecond)
	dataBody, signValue, account, data := "", "", "", ""
	//dataBody, signValue, account, data := GetSignAndDataCom(req.PostBody)
	log.Info(account, data)
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
	//Test
	params := `{
	"scope":"%s",
	"code":"%s",
	"table":"%s",
	"json":"%s"
	}`

	//{"scope":"buyertest", "code":"currency", "table":"account", "json": true}
	userName := req.Username
	//random := req.Random

	//signature := req.Signature

	s := fmt.Sprintf(params, userName, "currency", "account", "true")
	log.Info("s:", s)
	resp, err := http.Post(GET_TABLE_ROWS, "application/json",
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
		js, _ := simplejson.NewJson(body)
		log.Info("jss", string(body))
		result, _ := json.Marshal(js.Get("rows").GetIndex(0))

		messages := js.Get("rows").GetIndex(0)
		account := messages.Get("balance").MustString()
		log.Info(account);
		log.Info("mess:", string(result))

		rsp.Code = 1
		rsp.Msg = "Get Account Successful!"
		rsp.Data = account
		return nil
	}
	end_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Info("Time:", end_time-start_time)
	return nil
}
func (u *User) Transfer(ctx context.Context, req *user_proto.TransferRequest, rsp *user_proto.TransferResponse) error {
	start_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Info("reqBody:" + req.PostBody)
	dataBody, signValue, account, data := cbb.GetSignAndDataCom(req.PostBody)
	log.Info(account, data)
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

	//Write to BlockChain
	flag, result := cbb.WriteToBlockChain(req.PostBody, PUSH_TRANSACTION_URL)
	log.Info("OK1,", result)

	end_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Info("Time:", end_time-start_time)
	//ok1 = true
	if flag == false {
		if strings.Contains(result, "Account not found") {
			rsp.Code = 1301
			rsp.Msg = "Transfer balance Failed, The destination account does not exist."
			return nil
		}
		if strings.Contains(result, "integer underflow subtracting token balance") {
			rsp.Code = 1302
			rsp.Msg = "Transfer balance Failed, does not have enough balance."
			return nil
		}

		rsp.Code = 1300
		rsp.Msg = "Transfer balance Failed, unknow err."
		rsp.Data = result
		return nil
	} else {
		rsp.Code = 0
		rsp.Msg = "Transfer balance Successful!"
		rsp.Data = string(result)
		log.Info(string(result))
		return nil
	}
}
func (u *User) QueryTransfer(ctx context.Context, req *user_proto.QueryTransferRequest, rsp *user_proto.QueryTransferResponse) error {
	start_time := time.Now().UnixNano() / int64(time.Millisecond)
	dataBody, signValue, account, data := "", "", "", ""
	//dataBody, signValue, account, data := GetSignAndDataCom(req.PostBody)
	log.Info(account, data)
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

	if req.PageSize > 0 && req.PageSize <= 50{
		pageSize = int(req.PageSize)
	}

	skip = (pageNum - 1) * pageSize

	var where interface{}
	where = &bson.M{"type": "transfer"}
	log.Info(req.Username)
	if req.Username != "" {
		where = &bson.M{"type": "transfer", "data.From": req.Username}
	}

	var mgo = mgo.Session()
	defer mgo.Close()
	count, err := mgo.DB(config.DB_NAME).C("Messages").Find(where).Count()
	if err != nil {
		log.Error(err)
	}

	var ret []bean.TransferBean
	mgo.DB(config.DB_NAME).C("Messages").Find(where).Sort("-createAt").Skip(skip).Limit(pageSize).All(&ret)

	var rows = []*user_proto.TransferRow{}
	for _, v := range ret {
		rows = append(rows, &user_proto.TransferRow{
			v.TransactionID,
			v.Data.From,
			v.Data.To,
			v.Data.Quantity,
			v.CreatedAt.String(),
			v.BlockNum,
		})
	}

	var d = &user_proto.TransferData{
		PageNum:  uint64(pageNum),
		RowCount: uint64(count),
		Row:      rows,
	}

	rsp.Code = 0
	rsp.Data = d

	////Test
	//params := `service=storage&method=Storage.GetRecentTransferList&request={
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
	//	result, _ := json.Marshal(js.Get("transfer_list"))
	//	if js.Get("code").MustInt() == 1 {
	//
	//		rsp.Code = 1
	//		rsp.Msg = "Get Transfer List Successful!"
	//		rsp.Data = string(result)
	//		log.Info(result)
	//		log.Info(string(result))
	//	}
	//	return nil
	//}
	end_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Info("Time:", end_time-start_time)
	return nil
}

func (u *User) GetBlockInfo(ctx context.Context, req *user_proto.GetBlockInfoRequest, rsp *user_proto.GetBlockInfoResponse) error {
	ref_block_num := GetBolckNum()
	time.Sleep(1)
	ref_block_prefix, timestamp := GetBlockPrefix(ref_block_num)
	time.Sleep(1)
	expirationTime := ExpirationTime(timestamp, 20)

	rsp.Code = 0
	rsp.Msg = "OK"
	var blockInfo = &user_proto.BlockInfo{
		RefBlockNum:    ref_block_num,
		RefBlockPrefix: int64(ref_block_prefix),
		Expiration:     expirationTime,
	}
	rsp.Data = blockInfo
	return nil
}

func (u *User) GetDataBin(ctx context.Context, req *user_proto.GetDataBinRequest, rsp *user_proto.GetDataBinResponse) error {
	log.Info(req.Info)
	hexBin := GetHexBin(req.Info)
	log.Info(hexBin)
	time.Sleep(1)

	if hexBin != "" {
		rsp.Code = 0
		rsp.Msg = "OK"
		var bin = &user_proto.Bin{
			Bin: hexBin,
		}
		rsp.Data = bin
	} else {
		rsp.Code = 1109
		rsp.Msg = "Gain failure"
	}

	return nil
}

func GetHexBin(post string) string {
	log.Info("bin-post-json", post)
	resp, err := http.Post(ABI_JSON_TO_BIN_URL, "", bytes.NewReader([]byte(post)))
	if err != nil {
		log.Error(err.Error())
		if (index < 2) {
			index++
			GetHexBin(post)
		}
		return ""
	}
	index = 0;
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err.Error())
		return ""
	}
	log.Info("Bin-body:", string(body))
	if resp.StatusCode/100 == 2 {
		js, _ := simplejson.NewJson(body)
		return js.Get("binargs").MustString()
	} else {
		return ""
	}
}

func save_token(username string, token string) bool {
	var mgo = mgo.Session()
	defer mgo.Close()
	count, err := mgo.DB(config.DB_NAME).C("user_token").Find(&bson.M{"username": username}).Count()
	if err != nil {
		log.Error(err)
		return false
	}

	if count > 0 {
		err := mgo.DB(config.DB_NAME).C("user_token").Remove(&bson.M{"username": username})
		if err != nil {
			log.Error(err)
			return false
		}
	}

	err = mgo.DB(config.DB_NAME).C("user_token").Insert(&bson.M{"username": username, "token": token, "ctime": time.Now().Unix()})
	if err != nil {
		log.Error(err)
		return false
	} else {
		return true
	}
}

func create_token() string {
	uuid, _ := uuid.NewV4()
	str := strings.Replace(uuid.String(), "-", "", -1)
	t1 := strings.Trim(base64.StdEncoding.EncodeToString([]byte(str)),"=")
	t2 := strings.Trim(base64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(time.Now().Unix(),10))),"=")
	return  t1 + t2
}

type LoginRet struct {
	Processed struct {
		Expiration string `json:"expiration"`
		Messages []struct {
			Authorization []interface{} `json:"authorization"`
			Code          string        `json:"code"`
			Data struct {
				RandomNum float64 `json:"random_num"`
				UserName  string  `json:"user_name"`
			} `json:"data"`
			HexData string `json:"hex_data"`
			Type    string `json:"type"`
		} `json:"messages"`
		Output []struct {
			DeferredTrxs []interface{} `json:"deferred_trxs"`
			Notify       []interface{} `json:"notify"`
		} `json:"output"`
		RefBlockNum    float64       `json:"ref_block_num"`
		RefBlockPrefix float64       `json:"ref_block_prefix"`
		Scope          []string      `json:"scope"`
		Signatures     []interface{} `json:"signatures"`
	} `json:"processed"`
	TransactionID string `json:"transaction_id"`
}

func UserLogin(info string) (bool, string) {
	//js, _ := simplejson.NewJson([]byte(info))
	//messages := js.Get("messages").GetIndex(0)
	//authorization := messages.Get("authorization").GetIndex(0)
	//postData := map[string]interface{}{
	//	"ref_block_num":    js.Get("ref_block_num").MustInt(),
	//	"ref_block_prefix": js.Get("ref_block_prefix").MustInt(),
	//	"expiration":       js.Get("expiration").MustString(),
	//	"scope":            []string{js.Get("scope").MustString()},
	//	"read_scope":       []string{},
	//	"messages": []interface{}{
	//		map[string]interface{}{
	//			"code": messages.Get("code").MustString(),
	//			"type": messages.Get("type").MustString(),
	//			"authorization": []string{},
	//			//	[]interface{}{
	//			//	map[string]interface{}{
	//			//		"account":    authorization.Get("account").MustString(),
	//			//		"permission": authorization.Get("permission").MustString(),
	//			//	},
	//			//},
	//			"data": messages.Get("data").MustString(),
	//		},
	//	},
	//	"signatures": []string{},
	//}

	//bytesData, err := json.Marshal(postData)
	//log.Info(string(bytesData))
	//if err != nil {
	//	log.Error(err.Error())
	//	return false, ""
	//}
	resp, err := http.Post(PUSH_TRANSACTION_URL, "", bytes.NewReader([]byte(info)))
	if err != nil {
		log.Error(err.Error())
		if (index < 2) {
			index++
			UserLogin(info)
		}
		return false, ""
	}
	index = 0;
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err.Error())
		return false, ""
	}

	log.Info(string(body))
	if resp.StatusCode/100 == 2 {
		log.Info(string(body))
		js, _ := simplejson.NewJson(body)
		transaction_id := js.Get("transaction_id").MustString()
		if transaction_id != "" {
			var loginRet LoginRet
			json.Unmarshal(body, &loginRet)
			username := loginRet.Processed.Messages[0].Data.UserName
			return true, username
		} else {
			return false, ""
		}
	}
	return false, ""
}

func GetTableRow(username string) (string, int, error) {
	postData := map[string]interface{}{
		"scope":       "usermng",
		"code":        "usermng",
		"table":       "userreginfo",
		"json":        true,
		"strkeyvalue": username,
	}
	req := curl.NewRequest()
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

func GetBolckNum() (int64) {
	resp, err := http.Get(GET_INFO_URL)
	if err != nil {
		log.Error(err.Error())
		if index < 2 {
			index++
			GetBolckNum()
		}
		return -1
	}
	index = 0;
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err.Error())
		return -1
	}
	if resp.StatusCode/100 == 2 {
		js, _ := simplejson.NewJson(body)
		block_num := js.Get("head_block_num").MustInt64()
		return block_num
	}

	return -1
}

func GetBlockPrefix(block_num int64) (int64, string) {
	postData := map[string]interface{}{
		"block_num_or_id": block_num,
	}
	bytesData, err := json.Marshal(postData)
	if err != nil {
		log.Error(err.Error())
		return -1, ""
	}
	resp, err := http.Post(GET_BLOCK_URL, "", bytes.NewReader(bytesData))
	if err != nil {
		log.Error(err.Error())
		return GetBlockPrefix(block_num)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err.Error())
		return -1, ""
	}
	if resp.StatusCode/100 == 2 {
		js, _ := simplejson.NewJson(body)
		block_prefix := js.Get("ref_block_prefix").MustInt64()
		timestamp := js.Get("timestamp").MustString()
		return block_prefix, timestamp
	} else {
		return -1, ""
	}
}

func AccountGetBin(name string, owner_key string, active_key string) (string) {
	postData := map[string]interface{}{
		"code":   "bto",
		"action": "newaccount",
		"args": map[string]interface{}{
			"creator": "bto",
			"name":    name,
			"owner": map[string]interface{}{
				"threshold": 1,
				"keys": []interface{}{
					map[string]interface{}{
						"key":    owner_key,
						"weight": 1,
					},
				},
				"accounts": []string{},
			},
			"active": map[string]interface{}{
				"threshold": 1,
				"keys": []interface{}{
					map[string]interface{}{
						"key":    active_key,
						"weight": 1,
					},
				},
				"accounts": []string{},
			},
			"recovery": map[string]interface{}{
				"threshold": 1,
				"keys":      []string{},
				"accounts": []interface{}{map[string]interface{}{
					"permission": map[string]interface{}{
						"account":    "bto",
						"permission": "active",
					},
					"weight": 1,
				},
				},
			},
			"deposit": "0.00000001",
		},
	}

	bytesData, err := json.Marshal(postData)
	log.Info("POST", string(bytesData))
	if err != nil {
		log.Error(err.Error())
		return ""
	}
	resp, err := http.Post(ABI_JSON_TO_BIN_URL, "", bytes.NewReader(bytesData))
	if err != nil {
		log.Error(err.Error())
		if (index < 2) {
			index++
			AccountGetBin(name, owner_key, active_key)
		}
		return ""
	}
	index = 0;
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err.Error())
		return ""
	}
	log.Info("Bin:", string(body))
	if resp.StatusCode/100 == 2 {
		js, _ := simplejson.NewJson(body)
		binargs := js.Get("binargs").MustString()
		return binargs
	}
	return ""
}

func UserGetBin(username string, info string) (string) {
	postData := map[string]interface{}{
		"code":   "usermng",
		"action": "reguser",
		"args": map[string]interface{}{
			"user_name": username,
			"basic_info": map[string]interface{}{
				"info": info,
			},
		},
	}

	bytesData, err := json.Marshal(postData)
	if err != nil {
		log.Error(err.Error())
		return ""
	}
	log.Info("POST", string(bytesData))
	resp, err := http.Post(ABI_JSON_TO_BIN_URL, "", bytes.NewReader(bytesData))
	if err != nil {
		log.Error(err.Error())
		if (index < 2) {
			index++
			UserGetBin(username, info)
		}
		return ""
	}
	index = 0;
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err.Error())
		return ""
	}
	log.Info("body", string(body))
	if resp.StatusCode/100 == 2 {
		js, _ := simplejson.NewJson(body)
		binargs := js.Get("binargs").MustString()
		return binargs
	} else {
		return ""
	}
}

func AccountPushTransaction(ref_block_num int64, ref_block_prefix int64, expirationTime string, data string) (bool, string) {
	postData := map[string]interface{}{
		"ref_block_num":    ref_block_num,
		"ref_block_prefix": ref_block_prefix,
		"expiration":       expirationTime,
		"scope":            []string{"bto"},
		"read_scope":       []string{},
		"messages": []interface{}{
			map[string]interface{}{
				"code":          "bto",
				"type":          "newaccount",
				"authorization": []string{},
				"data":          data,
			},
		},
		"signatures": []string{},
	}
	bytesData, err := json.Marshal(postData)
	if err != nil {
		log.Error(err.Error())
		return false, ""
	}
	resp, err := http.Post(PUSH_TRANSACTION_URL, "", bytes.NewReader(bytesData))
	if err != nil {
		log.Error(err.Error())
		if (index < 2) {
			index++
			AccountPushTransaction(ref_block_num, ref_block_prefix, expirationTime, data)
		}
		return false, ""
	}
	index = 0;
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err.Error())
		return false, ""
	}
	js, _ := simplejson.NewJson(body)
	log.Info("push_transaction_accouunt_body",string(body))
	if resp.StatusCode/100 == 2 {
		transaction_id := js.Get("transaction_id").MustString()
		if transaction_id != "" {
			return true, "success"
		} else {
			return false, "transaction_id"
		}
	} else {
		details := js.Get("details").MustString()
		if strings.Contains(details, "3050000 message_precondition_exception") {
			return false, "Already"
		}
		return false, details
	}

}

func UserPushTransaction(ref_block_num int64, ref_block_prefix int64, expirationTime string, data string, username string) (string) {
	scope := []string{"usermng","bto", username}
	sort.Strings(scope)
	log.Info("scop", scope)
	postData := map[string]interface{}{
		"ref_block_num":    ref_block_num,
		"ref_block_prefix": ref_block_prefix,
		"expiration":       expirationTime,
		"scope":            scope,
		"read_scope":       []string{},
		"messages": []interface{}{
			map[string]interface{}{
				"code":          "usermng",
				"type":          "reguser",
				"authorization": []string{},
				"data":          data,
			},
		},
		"signatures": []string{},
	}

	bytesData, err := json.Marshal(postData)
	if err != nil {
		log.Error(err.Error())
		return ""
	}
	resp, err := http.Post(PUSH_TRANSACTION_URL, "", bytes.NewReader(bytesData))
	if err != nil {
		log.Error(err.Error())
		if (index < 2) {
			index++
			UserPushTransaction(ref_block_num, ref_block_prefix, expirationTime, data, username)
		}
		return ""
	}
	index = 0;
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err.Error())
		return ""
	}
	log.Info("userinfo-body", string(body))
	if resp.StatusCode/100 == 2 {

		return string(body)
	} else {
		return ""
	}
}

func ExpirationTime(toBeChargeTime string, delayTime int64) (string) {
	//获取本地location
	//toBeCharge := "2018-02-09T10:41:19" 								 //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "2006-01-02T15:04:05"                                 //转化所需模板
	loc, _ := time.LoadLocation("Local")                                //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, toBeChargeTime, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                                //转化为时间戳 类型是int64
	//fmt.Println(theTime)                                              //打印输出theTime 2015-01-01 15:15:00 +0800 CST
	//fmt.Println(sr)                                                   //打印输出时间戳 1420041600
	sr2 := sr + delayTime
	//log.Info("sr2", sr2)
	//获取时间戳
	//timestamp := time.Now().Unix()
	//fmt.Println("timestamp:",timestamp)
	//格式化为字符串,tm为Time类型
	//tm := time.Unix(sr2, 0)
	//fmt.Println(tm.Format("2006-01-02 03:04:05 PM"))

	//时间戳转日期
	//dataTimeStr := time.Unix(sr2, 0).Format(timeLayout) 					//设置时间戳 使用模板格式化为日期字符串
	dataTimeStr := time.Unix(sr2, 0).Format(timeLayout)
	//t, _ := time.Parse(timeLayout, dataTimeStr)
	//fmt.Println("haha",t,t.In(time.UTC))

	return dataTimeStr
}

func init() {
	logger, err := log.LoggerFromConfigAsFile("./config/api-user-log.xml")
	if err != nil{
		log.Error(err)
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


