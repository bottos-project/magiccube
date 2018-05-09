package main

import (
	log "github.com/jeanphorn/log4go"
	"github.com/micro/go-micro"
	proto "github.com/bottos-project/bottos/service/exchange/proto"
	"golang.org/x/net/context"
	"github.com/bitly/go-simplejson"
	"time"
	"bytes"
	"io/ioutil"
	"net/http"
	cbb "github.com/bottos-project/bottos/service/asset/cbb"
	"github.com/bottos-project/bottos/tools/db/mongodb"
	"gopkg.in/mgo.v2/bson"
	"github.com/bottos-project/bottos/service/bean"
	"github.com/bottos-project/bottos/config"
	"strings"
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

func (u *Exchange) ConsumerBuy(ctx context.Context, req *proto.ConsumerBuyRequest, rsp *proto.ConsumerBuyResponse) error {
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

	end_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Info("Time:", end_time-start_time)
	//ok1 = true
	if flag == false {
		if strings.Contains(result, "underflow subtracting token balance") {
			rsp.Code = 4001
			rsp.Msg = "Consumer Buy Asset Failed, The balance is too low."
			log.Debug("4001:", result)
			//rsp.Data = result
			return nil
		} else {
			rsp.Code = 4002
			log.Debug("4002:", result)
			rsp.Msg = "Consumer Buy Asset Failed, Unkown error."
			rsp.Data = result
			return nil
		}
	} else {
		rsp.Code = 1
		rsp.Msg = "ConsumerBuy Successful!"
		rsp.Data = string(result)
		return nil
	}

}

func (u *Exchange) QueryTx(ctx context.Context, req *proto.QueryTxRequest, rsp *proto.QueryTxResponse) error {

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

/*func WriteToBlockChain(post string, url string) []byte {
	log.Info(url, post)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(post)))
	// req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	log.Info(resp.Status)
	if resp.StatusCode/100 == 2 {
		body, _ := ioutil.ReadAll(resp.Body)
		js, _ := simplejson.NewJson([]byte(body))
		log.Info(atom.String(body))
		log.Info(string(body))
		log.Info(js)
		js.Get("result").MustString()
		return body
	} else {
		return nil
	}
}

func GetPublicKey(post string) string {
	req, err := http.NewRequest("POST", PUSH_TRANSACTION_URL, bytes.NewBuffer([]byte(post)))
	// req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.Status == "200 OK" {
		body, _ := ioutil.ReadAll(resp.Body)
		js, _ := simplejson.NewJson([]byte(body))
		result := js.Get("result").MustString()
		return result
	} else {
		return ""
	}
}
func GetSignAndDataCom(postBody string) (signData string, account string, sign string, data string) {
	js, _ := simplejson.NewJson([]byte(postBody))
	//get signed data
	//TODO

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
	log.Info(postData)
	//getSignValue
	signValue := js.Get("signatures").MustString()
	log.Info(signValue)
	//get Account
	account = authorization.Get("account").MustString()
	log.Info(account)
	//get sign Data
	delete(postData, "signatures")
	signData = ""
	//signData = string(json.Marshal(postData))
	log.Info(signData)
	//get sign Data
	data = messages.Get("data").MustString()
	log.Info("----------", data)

	*//*	req := curl.NewRequest()
		resp, err := req.SetUrl(PUSH_TRANSACTION_URL).SetPostData(postData).Post()
		if err != nil {
			return
		}*//*
	//return resp.Body, account
	return signData, account, signValue, data
}*/

func VerifySignOnBlockChain(post string, url string) bool {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(post)))
	// req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.Status == "200 OK" {
		body, _ := ioutil.ReadAll(resp.Body)
		js, _ := simplejson.NewJson([]byte(body))
		js.Get("result").MustString()
		//return js.Get("result").MustString()
		return true
	} else {
		return false
	}
}

/*func VerifySign(data string, sign string, pubKey string) (bool, string) {
	//if sign == "" {
	//	//从data中取sign TODO
	//}
	//flag := false
	//ToDO
	//var err string
	if data == sign {
		//flag = true
		return true, "Successful!"
	} else {
		return false, "Failed!"
	}

}*/

func main() {
	log.LoadConfiguration(config.BASE_LOG_CONF)
	defer log.Close()
	log.LOGGER("exchange.srv")

	service := micro.NewService(
		micro.Name("go.micro.srv.v2.exchange"),
		micro.Version("2.0.0"),
	)

	service.Init()

	//proto.RegisterUserHandler(service.Server(), new(Asset))
	//user_proto.RegisterUserHandler(service.Server(), new(User))
	proto.RegisterExchangeHandler(service.Server(), new(Exchange))

	if err := service.Run(); err != nil {
		log.Exit(err)
	}
}
