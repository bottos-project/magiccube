package main

import (
	log "github.com/jeanphorn/log4go"
	"github.com/micro/go-micro"
	user_proto "github.com/code/bottos/service/identity/proto"
	"golang.org/x/net/context"
	"github.com/mikemintang/go-curl"
	"github.com/bitly/go-simplejson"
	"time"
	"crypto/sha512"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"github.com/satori/go.uuid"
	"strings"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"bytes"
	"regexp"
	"fmt"
	cbb "github.com/code/bottos/service/asset/cbb"
	"strconv"
	"github.com/code/bottos/tools/db/mongodb"
	"gopkg.in/mgo.v2/bson"
	"github.com/code/bottos/service/bean"
	"github.com/code/bottos/config"
)

const (
	BASE_URL              	= config.BASE_URL
	GET_INFO_URL            = BASE_URL + "v1/chain/get_info"
	GET_BLOCK_URL           = BASE_URL + "v1/chain/get_block"
	ABI_JSON_TO_BIN_URL     = BASE_URL + "v1/chain/abi_json_to_bin"
	PUSH_TRANSACTION_URL    = BASE_URL + "v1/chain/push_transaction"
	GET_TABLE_ROW_BY_STRING = BASE_URL + "v1/chain/get_table_row_by_string_key"
	GET_TABLE_ROWS          = BASE_URL + "v1/chain/get_table_rows"
	STORAGE_RPC_URL         = config.BASE_RPC
	LOG_CONGFIG_FILE 		= "config/log.json"
)
var index int = 0

type User struct{}

func (u *User) Register(ctx context.Context, req *user_proto.RegisterRequest, rsp *user_proto.RegisterResponse) error {
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

	accoutBin := AccountGetBin(req.Username, req.OwnerPubKey, req.ActivePubKey)
	log.Info("accoutBin:", accoutBin)
	time.Sleep(1)
	if accoutBin ==  ""{
		rsp.Code = 1133
		rsp.Msg = "Data[account_bin] exception"
		return nil
	}

	expirationTime := ExpirationTime(timestamp, 20)
	accountInfo, msg:= AccountPushTransaction(ref_block_num, ref_block_prefix, expirationTime, accoutBin)
	log.Info("accountInfo:", accountInfo)
	if !accountInfo {
		if msg == "Already" {
			rsp.Code = 1103
			rsp.Msg = "Account has already existed"
			return nil
		}
		rsp.Code = 1101
		rsp.Msg = msg
		return nil
	}
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
	userInfo := UserPushTransaction(ref_block_num, ref_block_prefix, expirationTime, userBin)
	log.Info("userInfo:", userInfo)
	if userInfo != "" {
		rsp.Code = 0
		rsp.Msg = "Registered user success"
	} else {
		rsp.Code = 1102
		rsp.Msg = "Registered user failure"
	}
	return nil
}

func (u *User) Login(ctx context.Context, req *user_proto.LoginRequest, rsp *user_proto.LoginResponse) error {
	is_login, account := UserLogin(req.Body)
	if is_login {
		token := create_token(req.Header)
		if save_token(account, token) {
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
	if UserLogout(req.Token) {
		rsp.Code = 0
		rsp.Msg = "OK"
	} else {
		rsp.Code = -1
		rsp.Msg = "Error"
	}
	return nil
}

func (u *User) VerifyToken(ctx context.Context, req *user_proto.VerifyTokenRequest, rsp *user_proto.VerifyTokenResponse) error {
	//if CheckToken(req.Token) {
	if false {
		rsp.Code = 0
		rsp.Msg = "OK"
	} else {
		rsp.Code = -1
		rsp.Msg = "Invalid Token"
	}
	return nil
}
type UserTokenBean struct{
	Username string `bson:"username"`
	Token string `bson:"token"`
	Ctime int64 `bson:"ctime"`
}

func CheckToken(token string) bool {
	var mgo = mgo.Session()
	defer mgo.Close()
	var ret UserTokenBean
	err := mgo.DB("local").C("usertoken").Find(&bson.M{"token":token}).One(&ret)
	if err != nil {
		log.Error(err)
		return false
	}

	if ret.Token == token{
		if (ret.Ctime+2*60*60 > time.Now().Unix()) {
			return true
		}
	}
	return false

	//resp, err := http.Post("http://10.104.11.217:8080/rpc",
	//	"application/x-www-form-urlencoded",
	//	strings.NewReader("service=storage&method=Storage.GetUserToken&request={\" token \":\""+token+"\"}"))
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	panic(err)
	//}
	//log.Info(string(body))
	//js, _ := simplejson.NewJson([]byte(body))
	//insert_time := js.Get("insert_time").MustInt64()
	//if (insert_time+15*60 < time.Now().Unix()) {
	//	return true
	//}
	//return false
}

func UserLogout(token string) bool {
	var mgo = mgo.Session()
	defer mgo.Close()
	err := mgo.DB("local").C("usertoken").Remove(&bson.M{"token":token})
	if err != nil {
		log.Error(err)
		return false
	}
	return true
	//resp, err := http.Post("http://10.104.11.217:8080/rpc",
	//	"application/x-www-form-urlencoded",
	//	strings.NewReader("service=storage&method=Storage.DelUserToken&request={\"token\":\""+token+"=\"}"))
	//if err != nil {
	//	log.Error(err)
	//}
	//
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return false
	//}
	//log.Info(string(body))
	//js, _ := simplejson.NewJson([]byte(body))
	//code := js.Get("code").MustInt()
	//if (code == 1) {
	//	return true
	//}
	//return false
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
	userInfo := UserPushTransaction(ref_block_num, ref_block_prefix, expirationTime, userBin)
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
	result := cbb.WriteToBlockChain(req.PostBody, PUSH_TRANSACTION_URL)
	log.Info("OK1,", result)

	end_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Info("Time:", end_time-start_time)
	//ok1 = true
	if result == nil {
		rsp.Code = 2000
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
				"code":    "eos",
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

	var pageNum, pageSize, skip int= 1, 20, 0
	if req.PageNum > 0 {
		pageNum = int(req.PageNum)
	}

	if req.PageSize > 0 {
		pageSize = int(req.PageSize)
	}

	skip = (pageNum - 1) *  pageSize

	var where interface{}
	where = &bson.M{"type": "favoritepro"}
	log.Info(req.Username)
	if req.Username != "" {
		where = &bson.M{"type": "favoritepro","data.user_name": req.Username}
	}

	var mgo = mgo.Session()
	defer mgo.Close()
	count, err:= mgo.DB("bottos").C("Messages").Find(where).Count()
	if err != nil {
		log.Error(err)
	}

	var ret []bean.FavoriteBean
	mgo.DB("bottos").C("Messages").Find(where).Sort("-createAt").Skip(skip).Limit(int(req.PageSize)).All(&ret)

	var rows = []*user_proto.FavoriteRowData{}
	for _, v := range ret {
		rows = append(rows, &user_proto.FavoriteRowData{
			v.Data.UserName,
			v.Data.OpType,
			v.Data.GoodsType,
			v.Data.GoodsID,
		})
	}

	var d = &user_proto.FavoriteData{
		PageNum: uint64(pageNum),
		RowCount:uint64(count),
		Row:rows,
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
	result := cbb.WriteToBlockChain(req.PostBody, PUSH_TRANSACTION_URL)
	log.Info("OK1,", result)

	end_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Info("Time:", end_time-start_time)
	//ok1 = true
	if result == nil {
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
				"code":    "eos",
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

	var pageNum, pageSize, skip int= 1, 20, 0
	if req.PageNum > 0 {
		pageNum = int(req.PageNum)
	}

	if req.PageSize > 0 {
		pageSize = int(req.PageSize)
	}

	skip = (pageNum - 1) *  pageSize

	var where interface{}
	where = &bson.M{"type": "datafilereg"}
	log.Info(req.Username)
	if req.Username != "" {
		where = &bson.M{"type": "datafilereg", "data.basic_info.user_name": req.Username}
	}

	var mgo = mgo.Session()
	defer mgo.Close()
	count, err:= mgo.DB("bottos").C("Messages").Find(where).Count()
	if err != nil {
		log.Error(err)
	}

	var ret []bean.FileBean
	mgo.DB("bottos").C("Messages").Find(where).Sort("-createAt").Skip(skip).Limit(int(req.PageSize)).All(&ret)

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
		PageNum: uint64(pageNum),
		RowCount:uint64(count),
		Row:rows,
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
		result := cbb.WriteToBlockChain(req.PostBody, PUSH_TRANSACTION_URL)
		log.Info("OK1,", result)

		end_time := time.Now().UnixNano() / int64(time.Millisecond)
		log.Info("Time:", end_time-start_time)
		//ok1 = true
		if result == nil {
		rsp.Code = 2000
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
					"code":    "eos",
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
	params := `service=storage&method=Storage.GetUserDataPresale&request={
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
		result, _ := json.Marshal(js.Get("data_presale_list"))
		if js.Get("code").MustInt() == 1 {

			rsp.Code = 1
			rsp.Msg = "Query Favorite List Successful!"
			rsp.Data = string(result)
		}
		return nil
	}
	end_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Info("Time:", end_time-start_time)
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
	resp, err := http.Post(GET_TABLE_ROWS, "application/x-www-form-urlencoded",
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
		result, _ := json.Marshal(js.Get("rows").GetIndex(0))

		messages := js.Get("rows").GetIndex(0)
		account := messages.Get("balance").MustInt()

		log.Info("mess:",string(result))

		rsp.Code = 1
		rsp.Msg = "Get Account Successful!"
		rsp.Data = strconv.Itoa(account)
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
	result := cbb.WriteToBlockChain(req.PostBody, PUSH_TRANSACTION_URL)
	log.Info("OK1,", result)

	end_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Info("Time:", end_time-start_time)
	//ok1 = true
	if result == nil {
		rsp.Code = 2000
		rsp.Msg = "Transfer account Failed."
		return nil
	} else {
		rsp.Code = 1
		rsp.Msg = "Transfer account Successful!"
		rsp.Data = string(result)
		log.Info(string(result))
		return nil
	}

	/*	获取data JSON格式数据
	postData1 := map[string]interface{}{
			"code":    "eos",
			"action":  "newaccount",
			"binargs": data,
		}
		reqBinToJson := curl.NewRequest()
		resp, err := reqBinToJson.SetUrl(ABI_BIN_TO_JSON_URL).SetPostData(postData1).Post()

		if err != nil {
			return nil
		}*/
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

	var pageNum, pageSize, skip int= 1, 20, 0
	if req.PageNum > 0 {
		pageNum = int(req.PageNum)
	}

	if req.PageSize > 0 {
		pageSize = int(req.PageSize)
	}

	skip = (pageNum - 1) *  pageSize

	var where interface{}
	where = &bson.M{"type": "transfer"}
	log.Info(req.Username)
	if req.Username != "" {
		where = &bson.M{"type": "transfer", "data.From": req.Username}
	}

	var mgo = mgo.Session()
	defer mgo.Close()
	count, err:= mgo.DB("bottos").C("Messages").Find(where).Count()
	if err != nil {
		log.Error(err)
	}

	var ret []bean.TransferBean
	mgo.DB("bottos").C("Messages").Find(where).Sort("-createAt").Skip(skip).Limit(int(req.PageSize)).All(&ret)

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
		PageNum: uint64(pageNum),
		RowCount:uint64(count),
		Row:rows,
	}

	rsp.Code = 1;
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
	req, err := http.NewRequest("POST", ABI_JSON_TO_BIN_URL, bytes.NewBuffer([]byte(post)))
	// req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
	}
	defer resp.Body.Close()

	if resp.Status == "200 OK" {
		body, _ := ioutil.ReadAll(resp.Body)
		js, _ := simplejson.NewJson([]byte(body))
		return js.Get("binargs").MustString()
	} else {
		return ""
	}
}



func save_token(username string, token string) bool {

	var mgo = mgo.Session()
	defer mgo.Close()
	count, err := mgo.DB("local").C("usertoken").Find(&bson.M{"username":username,"token":token}).Count()
	if err != nil {
		log.Error(err)
		return false
	}

	if count > 0 {
		err := mgo.DB("local").C("usertoken").Remove(&bson.M{"username":username,"token":token})
		if err != nil {
			log.Error(err)
			return false
		}
	}

	err = mgo.DB("local").C("usertoken").Insert(&bson.M{"username":username,"token":token,"ctime":time.Now().Unix()})
	if err != nil {
		log.Error(err)
		return false
	}else{
		return true
	}


	//resp, err := http.Post("http://10.104.11.217:8080/rpc",
	//	"application/x-www-form-urlencoded",
	//	strings.NewReader("service=storage&method=Storage.GetUserToken&request={\" username \":\""+username+"\"}"))
	//if err != nil {
	//	return false
	//}
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return false
	//}
	//log.Info(string(body))
	//js, _ := simplejson.NewJson([]byte(body))
	//JS_username := js.Get("username").MustString()
	//if (JS_username == username) {
	//	resp1, err1 := http.Post("http://10.104.11.217:8080/rpc",
	//		"application/x-www-form-urlencoded",
	//		strings.NewReader("service=storage&method=Storage.DelUserToken&request={\" username \":\""+username+"\"}"))
	//	if err1 != nil {
	//		return false
	//	}
	//	defer resp1.Body.Close()
	//	body1, err := ioutil.ReadAll(resp1.Body)
	//	if err != nil {
	//		return false
	//	}
	//	js1, _ := simplejson.NewJson([]byte(body1))
	//	code1 := js1.Get("code").MustInt()
	//	if (code1 != 1) {
	//		return false
	//	}
	//}
	//
	//resp2, err2 := http.Post("http://10.104.11.217:8080/rpc",
	//	"application/x-www-form-urlencoded",
	//	strings.NewReader("service=storage&method=Storage.InsertUserToken&request={\" username \":\""+username+"\",\"token\":\""+token+"=\"}"))
	//if err2 != nil {
	//	return false
	//}
	//
	//defer resp.Body.Close()
	//body2, err2 := ioutil.ReadAll(resp2.Body)
	//if err != nil {
	//	return false
	//}
	//log.Info(string(body))
	//js2, _ := simplejson.NewJson([]byte(body2))
	//code2 := js2.Get("code").MustInt()
	//if (code2 == 1) {
	//	return true
	//}
	//return false
}

func create_token(str string) string {
	uuid, _ := uuid.NewV4()
	uu := uuid.String()
	dateStr := time.Now().Local().String()
	h := sha512.New()
	h.Write([]byte(uu + dateStr + str))
	bs := h.Sum(nil)
	//SHA1 值经常以 16 进制输出，例如在 git commit 中。使用%x 来将散列结果格式化为 16 进制字符串。
	hh := md5.New()
	hh.Write([]byte(bs)) // 需要加密的字符串为 123456
	cipherStr := hh.Sum(nil)
	reg := regexp.MustCompile(`=`)
	return reg.ReplaceAllString(base64.StdEncoding.EncodeToString([]byte(hex.EncodeToString(cipherStr))), "")
}
func UserLogin(info string) (bool, string) {

	js, _ := simplejson.NewJson([]byte(info))
	messages := js.Get("messages").GetIndex(0)
	authorization := messages.Get("authorization").GetIndex(0)
	log.Info("----------", authorization.Get("account").MustString())

	postData := map[string]interface{}{
		"ref_block_num":    js.Get("ref_block_num").MustInt(),
		"ref_block_prefix": js.Get("ref_block_prefix").MustInt(),
		"expiration":       js.Get("expiration").MustString(),
		"scope":            []string{js.Get("scope").MustString()},
		"read_scope":       []string{},
		"messages": []interface{}{
			map[string]interface{}{
				"code": messages.Get("code").MustString(),
				"type": messages.Get("type").MustString(),
				"authorization": []interface{}{
					map[string]interface{}{
						"account":    authorization.Get("account").MustString(),
						"permission": authorization.Get("permission").MustString(),
					},
				},
				"data": messages.Get("data").MustString(),
			},
		},
		"signatures": []string{js.Get("signatures").MustString()},
	}

	req := curl.NewRequest()
	resp, err := req.SetUrl(PUSH_TRANSACTION_URL).SetPostData(postData).Post()
	if err != nil {
		log.Error(err)
	}

	if resp.Raw.StatusCode/100 == 2 {
		//js, _ := simplejson.NewJson([]byte(resp.Body))
		//binargs := js.Get("binargs").MustString()
		return true, authorization.Get("account").MustString()
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
		if(index < 2) {
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
	if resp.StatusCode / 100 == 2 {
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
		if(index < 2) {
			index++
			GetBlockPrefix(block_num)
		}
		return -1, ""
	}
	index = 0;
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err.Error())
		return -1, ""
	}
	if resp.StatusCode / 100 == 2 {
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
		"code": "eos",
		"action": "newaccount",
		"args": map[string]interface{}{
			"creator": "inita",
			"name": name,
			"owner": map[string]interface{}{
				"threshold": 1,
				"keys": []interface{}{
					map[string]interface{}{
						"key": owner_key,
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
						"account":    "inita",
						"permission": "active",
					},
					"weight": 1,
				},
				},
			},
			"deposit": "0.0001 EOS",
		},
	}

	bytesData, err := json.Marshal(postData)
	if err != nil {
		log.Error(err.Error())
		return ""
	}
	resp, err := http.Post(ABI_JSON_TO_BIN_URL, "", bytes.NewReader(bytesData))
	if err != nil {
		log.Error(err.Error())
		if(index < 2) {
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
	if resp.StatusCode / 100 == 2 {
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
	resp, err := http.Post(ABI_JSON_TO_BIN_URL, "", bytes.NewReader(bytesData))
	if err != nil {
		log.Error(err.Error())
		if(index < 2) {
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
	if resp.StatusCode / 100 == 2 {
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
		"scope":            []string{"eos", "inita"},
		"read_scope":       []string{},
		"messages": []interface{}{
			map[string]interface{}{
				"code":          "eos",
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
		if(index < 2) {
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
	if resp.StatusCode / 100 == 2 {
		transaction_id := js.Get("transaction_id").MustString()
		if transaction_id != "" {
			return true, "success"
		}else{
			return false, "transaction_id"
		}
	}else {
		details := js.Get("details").MustString()
		if strings.Contains(details, "3050000 message_precondition_exception") {
			return false, "Already"
		}
		return false, details
	}

}

func UserPushTransaction(ref_block_num int64, ref_block_prefix int64, expirationTime string, data string) (string) {
	postData := map[string]interface{}{
		"ref_block_num":    ref_block_num,
		"ref_block_prefix": ref_block_prefix,
		"expiration":       expirationTime,
		"scope":            []string{"usermng"},
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
		if(index < 2) {
			index++
			UserPushTransaction(ref_block_num, ref_block_prefix, expirationTime, data)
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

	if resp.StatusCode / 100 == 2 {
		log.Info(string(body))
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

func main() {
	//log.LoadConfiguration(LOG_CONGFIG_FILE)
	//defer log.Close()
	//log.LOGGER("user.srv")

	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("2.0.0"),
	)

	service.Init()

	user_proto.RegisterUserHandler(service.Server(), new(User))

	if err := service.Run(); err != nil {
		log.Exit(err)
	}

}
