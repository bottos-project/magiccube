package verifySi

import (
	"log"
	//"github.com/micro/go-micro"
	proto "github.com/code/bottos/service/asset/proto"
	"golang.org/x/net/context"
	"github.com/mikemintang/go-curl"
	"github.com/bitly/go-simplejson"
	"time"
	//"os/user"
	//storage "code/service/storage/proto"
	storage "github.com/code/bottos/service/storage/proto"
	"github.com/micro/go-micro/client"
	//"reflect"
	"bytes"
	"io/ioutil"
	"net/http"
	//"html/template"
	"strings"
	"fmt"
	"strconv"
	"golang.org/x/net/html/atom"
	"encoding/json"
	"os"
	"runtime"
)

const (
	//BASE_URL             = "http://127.0.0.1:8888/"
	BASE_URL             = "http://47.98.47.148:8887/"
	GET_INFO_URL         = BASE_URL + "v1/chain/get_info"
	GET_BLOCK_URL        = BASE_URL + "v1/chain/get_block"
	ABI_JSON_TO_BIN_URL  = BASE_URL + "v1/chain/abi_json_to_bin"
	ABI_BIN_TO_JSON_URL  = BASE_URL + "v1/chain/abi_bin_to_json"
	PUSH_TRANSACTION_URL = BASE_URL + "v1/chain/push_transaction"
	GET_TABLE_ROW        = BASE_URL + "v1/chain/get_table_row_by_string_key"
	STORAGE_RPC_URL      = "http://127.0.0.1:8080/rpc"
	//STORAGE_RPC_URL      = "http://10.104.20.254:8080/rpc"
)

type Asset struct{}


var ostype = runtime.GOOS

func GetProjectPath() string{
	var projectPath string
	projectPath, _ = os.Getwd()
	return projectPath
}

func GetConfigPath() string{
	path := GetProjectPath()
	if ostype == "windows"{
		path = path + "\\" + "config\\log.json"
	}else if ostype == "linux"{
		path = path +"/" + "config/log.json"
	}
	return  path
}

func GetConLogPath() string{
	path := GetProjectPath()
	if ostype == "windows"{
		path = path + "\\log\\"
	}else if ostype == "linux"{
		path = path + "/log/"
	}
	return  path
}

func (u *Asset) GetFileUploadURL(ctx context.Context, req *proto.GetFileUploadURLRequest, rsp *proto.GetFileUploadURLResponse) error {
	start_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Println("reqBody:" + req.PostBody)
	//dataBody, signValue, userName := "fd","","13"
	dataBody, signValue, userName := GetSignAndData(req.PostBody)
	log.Println(userName)
	//get Public Key
	pubKey := GetPublicKey("userName")
	//Verify Sign Local
	ok, _ := VerifySign(dataBody, signValue, pubKey)

	ok = true
	if !ok {
		rsp.Code = 2000
		rsp.Msg = "Verify Signature Failed."
		return nil
	}
	log.Println(ok)
	//get strore Address
	js, _ := simplejson.NewJson([]byte(req.PostBody))
	log.Println("js", js)

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

	log.Println(resp)
	log.Println("err", err)
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
		log.Println(presigned_put_url)
		return nil
	}
	log.Println(string(body))

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
	//log.Println("rspd:", rspd)
	//log.Println(rspd.PresignedPutUrl)

	end_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Println("Time:", end_time-start_time)
	if err != nil {
		log.Println(err)
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
	log.Println("dataBody", dataBody)
	//getSignValue
	signValue := js.Get("signatures").MustString()
	log.Println(signValue)
	//get username
	userName := js.Get("userName").MustString()

	//messages := js.Get("messages").GetIndex(0)
	//authorization := messages.Get("authorization").GetIndex(0)
	//log.Println("----------", authorization.Get("account").MustString())

	//postData := map[string]interface{}{
	//	"ref_block_num": js.Get("ref_block_num").MustInt(),
	//}
	return dataBody, signValue, userName
}

func (u *Asset) RegisterFile(ctx context.Context, req *proto.RegisterFileRequest, rsp *proto.RegisterFileResponse) error {
	start_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Println("reqBody:" + req.PostBody)
	dataBody, signValue, account, data := GetSignAndDataCom(req.PostBody)
	log.Println(account, data)
	//get Public Key
	pubKey := GetPublicKey("account")
	//Verify Sign Local
	ok, _ := VerifySign(dataBody, signValue, pubKey)
	log.Println(ok)
	ok = true
	if !ok {
		rsp.Code = 2000
		rsp.Msg = "Verify Signature Failed."
		return nil
	}

	//Write to BlockChain
	result := WriteToBlockChain(req.PostBody, PUSH_TRANSACTION_URL)
	log.Println("OK1,", result)

	end_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Println("Time:", end_time-start_time)
	//ok1 = true
	if result == nil {
		rsp.Code = 2000
		rsp.Msg = "Register File Failed."
		return nil
	} else {
		rsp.Code = 1
		rsp.Msg = "Register File Successful!"
		rsp.Data = string(result)
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
func WriteToBlockChain(post string, url string) []byte {
	log.Println(url, post)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(post)))
	// req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	log.Println(resp.Status)
	if resp.StatusCode/100 == 2 {
		body, _ := ioutil.ReadAll(resp.Body)
		js, _ := simplejson.NewJson([]byte(body))
		log.Println(atom.String(body))
		log.Println(string(body))
		log.Println(js)
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
	log.Println("----------", authorization.Get("account").MustString())

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
	log.Println(postData)
	//getSignValue
	signValue := js.Get("signatures").MustString()
	log.Println(signValue)
	//get Account
	account = authorization.Get("account").MustString()
	log.Println(account)
	//get sign Data
	delete(postData, "signatures")
	signData = ""
	//signData = string(json.Marshal(postData))
	log.Println(signData)
	//get sign Data
	data = messages.Get("data").MustString()
	log.Println("----------", data)

	/*	req := curl.NewRequest()
		resp, err := req.SetUrl(PUSH_TRANSACTION_URL).SetPostData(postData).Post()
		if err != nil {
			return
		}*/
	//return resp.Body, account
	return signData, account, signValue, data
}

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

func VerifySign(data string, sign string, pubKey string) (bool, string) {
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

}

func (u *Asset) GetFileUploadStat(ctx context.Context, req *proto.GetFileUploadStatRequest, rsp *proto.GetFileUploadStatResponse) error {

	//Test
	params := `service=storage&method=Storage.GetFileUploadStat&request={
	"username":"%s",
	"file_name":"%s"
	}`
	userName := req.Username
	fileName := req.FileName
	log.Println(userName, fileName)
	s := fmt.Sprintf(params, userName, fileName)
	resp, err := http.Post(STORAGE_RPC_URL, "application/x-www-form-urlencoded",
		strings.NewReader(s))

	//log.Println("resp:",resp)
	//log.Println("err", err)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	} else {
		jss, _ := simplejson.NewJson([]byte(body))
		log.Println("jss", jss)
		uploadStat := jss.Get("result").MustString()
		message := jss.Get("message").MustString()
		fileSize := jss.Get("size").MustInt64()
		log.Println(fileSize)
		if uploadStat == "200" {
			rsp.Code = 1
			rsp.Msg = message
			rsp.Data = strconv.FormatInt(fileSize, 10)
			log.Println("filesize:", fileSize)
		}
		return nil
	}
	//log.Println(string(body))

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

func (u *Asset) QueryUploadedData(ctx context.Context, req *proto.QueryUploadedDataRequest, rsp *proto.QueryUploadedDataResponse) error {
	start_time := time.Now().UnixNano() / int64(time.Millisecond)
	dataBody, signValue, account, data := "", "", "", ""
	//dataBody, signValue, account, data := GetSignAndDataCom(req.PostBody)
	log.Println(account, data)
	//get Public Key
	pubKey := GetPublicKey("account")
	//Verify Sign Local
	ok, _ := VerifySign(dataBody, signValue, pubKey)
	log.Println(ok)
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
	log.Println("s:", s)
	resp, err := http.Post(STORAGE_RPC_URL, "application/x-www-form-urlencoded",
		strings.NewReader(s))

	log.Println("resp:", resp)
	//log.Println("err", err)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	} else {
		js, _ := simplejson.NewJson([]byte(body))
		log.Println("jss", js)
		json.Marshal(js.Get("FileList"))
		if js.Get("code").MustInt() == 1 {

			rsp.Code = 1
			rsp.Msg = "Get File List Successful!"
			rsp.Data = nil
		}
		return nil
	}
	end_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Println("Time:", end_time-start_time)
	return nil
}

func (u *Asset) GetDownLoadURL(ctx context.Context, req *proto.GetDownLoadURLRequest, rsp *proto.GetDownLoadURLResponse) error {
	start_time := time.Now().UnixNano() / int64(time.Millisecond)
	dataBody, signValue, account, data := "", "", "", ""
	//dataBody, signValue, account, data := GetSignAndDataCom(req.PostBody)
	log.Println(account, data)
	//get Public Key
	pubKey := GetPublicKey("account")
	//Verify Sign Local
	ok, _ := VerifySign(dataBody, signValue, pubKey)
	log.Println(ok)
	ok = true
	if !ok {
		rsp.Code = 2000
		rsp.Msg = "Verify Signature Failed."
		return nil
	}
	//Test
	params := `service=storage&method=Storage.GetDownLoadURL&request={
	"username":"%s",
	"file_name":"%s"
	}`
	userName := req.Username
	fileName := req.FileName

	//signature := req.Signature

	s := fmt.Sprintf(params, userName, fileName)
	log.Println("s:", s)
	resp, err := http.Post(STORAGE_RPC_URL, "application/x-www-form-urlencoded",
		strings.NewReader(s))

	log.Println("resp:", resp)
	//log.Println("err", err)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	} else {
		js, _ := simplejson.NewJson([]byte(body))
		log.Println("jss", js)
		//result,_ :=json.Marshal(js.Get("FileList"))
		downLoadURL := js.Get("presigned_get_url").MustString()
		if js.Get("result").MustString() == "200" {

			rsp.Code = 1
			rsp.Msg = "Get downLoad URL Successful!"
			rsp.Data = downLoadURL
		}
		//err,_ := json.Marshal(js)
		rsp.Msg="Get downLoad URL Successful!"
		return nil
	}
	end_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Println("Time:", end_time-start_time)
	return nil
}

func (u *Asset) Register(ctx context.Context, req *proto.RegisterRequest, rsp *proto.RegisterResponse) error {
	start_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Println("reqBody:" + req.PostBody)
	dataBody, signValue, account, data := GetSignAndDataCom(req.PostBody)
	log.Println(account, data)
	//get Public Key
	pubKey := GetPublicKey("account")
	//Verify Sign Local
	ok, _ := VerifySign(dataBody, signValue, pubKey)
	log.Println(ok)
	ok = true
	if !ok {
		rsp.Code = 2000
		rsp.Msg = "Verify Signature Failed."
		return nil
	}

	//Write to BlockChain
	result := WriteToBlockChain(req.PostBody, PUSH_TRANSACTION_URL)
	log.Println("OK1,", result)

	end_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Println("Time:", end_time-start_time)
	//ok1 = true
	if result == nil {
		rsp.Code = 2000
		rsp.Msg = "Register Asset Failed."
		return nil
	} else {
		rsp.Code = 1
		rsp.Msg = "Register Asset Successful!"
		rsp.Data = string(result)
		log.Println(string(result))
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

func (u *Asset) Query(ctx context.Context, req *proto.QueryRequest, rsp *proto.QueryResponse) error {
	start_time := time.Now().UnixNano() / int64(time.Millisecond)
	dataBody, signValue, account, data := "", "", "", ""
	//dataBody, signValue, account, data := GetSignAndDataCom(req.PostBody)
	log.Println(account, data)
	//get Public Key
	pubKey := GetPublicKey("account")
	//Verify Sign Local
	ok, _ := VerifySign(dataBody, signValue, pubKey)
	log.Println(ok)
	ok = true
	if !ok {
		rsp.Code = 2000
		rsp.Msg = "Verify Signature Failed."
		return nil
	}
	//Test
	params := `service=storage&method=Storage.GetUserAssetList&request={
	"username":"%s"
	}`
	userName := req.Username
	//random := req.Random

	//signature := req.Signature

	s := fmt.Sprintf(params, userName)
	log.Println("s:", s)
	resp, err := http.Post(STORAGE_RPC_URL, "application/x-www-form-urlencoded",
		strings.NewReader(s))

	log.Println("resp:", resp)
	//log.Println("err", err)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	} else {
		js, _ := simplejson.NewJson([]byte(body))
		log.Println("jss", js)
		result, _ := json.Marshal(js.Get("UserAssetList"))
		if js.Get("code").MustInt() == 1 {

			rsp.Code = 1
			rsp.Msg = "Get File List Successful!"
			rsp.Data = nil
			log.Println(result)
			log.Println(string(result))
		}
		return nil
	}
	end_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Println("Time:", end_time-start_time)
	return nil
}

/*func (u *Asset) QueryAllAsset(ctx context.Context, req *proto.QueryAllAssetRequest, rsp *proto.QueryAllAssetResponse) error {
	start_time := time.Now().UnixNano() / int64(time.Millisecond)
	dataBody, signValue, account, data := "", "", "", ""
	//dataBody, signValue, account, data := GetSignAndDataCom(req.PostBody)
	log.Println(account, data)
	//get Public Key
	pubKey := GetPublicKey("account")
	//Verify Sign Local
	ok, _ := VerifySign(dataBody, signValue, pubKey)
	log.Println(ok)
	ok = true
	if !ok {
		rsp.Code = 2000
		rsp.Msg = "Verify Signature Failed."
		return nil
	}
	//Test
	params := `service=storage&method=Storage.GetAllAssetList&request={
	"username":"%s"
	}`
	userName := req.Username
	//random := req.Random

	//signature := req.Signature

	s := fmt.Sprintf(params, userName)
	log.Println("s:", s)
	resp, err := http.Post(STORAGE_RPC_URL, "application/x-www-form-urlencoded",
		strings.NewReader(s))

	log.Println("resp:", resp)
	//log.Println("err", err)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	} else {
		js, _ := simplejson.NewJson([]byte(body))
		log.Println("jss", js)
		result, _ := json.Marshal(js.Get("UserAssetList"))
		if js.Get("code").MustInt() == 1 {

			rsp.Code = 1
			rsp.Msg = "Get File List Successful!"
			rsp.Data = string(result)
			log.Println(result)
			log.Println(string(result))
		}
		return nil
	}
	end_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Println("Time:", end_time-start_time)
	return nil
}*/

func (u *Asset) Modify(ctx context.Context, req *proto.ModifyRequest, rsp *proto.ModifyResponse) error {
	start_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Println("reqBody:" + req.PostBody)
	dataBody, signValue, account, data := GetSignAndDataCom(req.PostBody)
	log.Println(account, data)
	//get Public Key
	pubKey := GetPublicKey("account")
	//Verify Sign Local
	ok, _ := VerifySign(dataBody, signValue, pubKey)
	log.Println(ok)
	ok = true
	if !ok {
		rsp.Code = 2000
		rsp.Msg = "Verify Signature Failed."
		return nil
	}

	//Write to BlockChain
	result := WriteToBlockChain(req.PostBody, PUSH_TRANSACTION_URL)
	log.Println("OK1,", result)

	end_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Println("Time:", end_time-start_time)
	//ok1 = true
	if result == nil {
		rsp.Code = 2000
		rsp.Msg = "Modify Asset Failed."
		return nil
	} else {
		rsp.Code = 1
		rsp.Msg = "Modify Asset Successful!"
		rsp.Data = string(result)
		log.Println(string(result))
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

func GetBolckNum() (int, int, error) {
	req := curl.NewRequest()
	resp, err := req.SetUrl(GET_INFO_URL).Get()
	if err != nil {
		return 0, resp.Raw.StatusCode, err
	}

	if resp.IsOk() {
		js, _ := simplejson.NewJson([]byte(resp.Body))
		block_num := js.Get("head_block_num").MustInt()
		return block_num, resp.Raw.StatusCode, err
	} else {
		return 0, resp.Raw.StatusCode, err
	}
}

func GetBlockPrefix(block_num int) (int, string, int, error) {
	postData := map[string]interface{}{
		"block_num_or_id": block_num,
	}
	req := curl.NewRequest()
	resp, err := req.SetUrl(GET_BLOCK_URL).SetPostData(postData).Post()
	if err != nil {
		return 0, "", resp.Raw.StatusCode, err
	}

	if resp.IsOk() {
		js, _ := simplejson.NewJson([]byte(resp.Body))
		block_prefix := js.Get("ref_block_prefix").MustInt()
		timestamp := js.Get("timestamp").MustString()
		return block_prefix, timestamp, resp.Raw.StatusCode, err
	} else {
		return 0, "", resp.Raw.StatusCode, err
	}
}

func AccountGetBin(name string, owner_key string, active_key string) (string, int, error) {
	postData := map[string]interface{}{
		"code":   "eos",
		"action": "newaccount",
		"args": map[string]interface{}{
			"creator": "inita",
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
	req := curl.NewRequest()
	resp, err := req.SetUrl(ABI_JSON_TO_BIN_URL).SetPostData(postData).Post()
	if err != nil {
		return "", resp.Raw.StatusCode, err
	}

	js, _ := simplejson.NewJson([]byte(resp.Body))
	binargs := js.Get("binargs").MustString()
	return binargs, resp.Raw.StatusCode, err

}

func UserGetBin(username string, info string) (string, int, error) {
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
	req := curl.NewRequest()
	resp, err := req.SetUrl(ABI_JSON_TO_BIN_URL).SetPostData(postData).Post()
	if err != nil {
		return "", resp.Raw.StatusCode, err
	}

	if resp.Raw.StatusCode/100 == 2 {
		js, _ := simplejson.NewJson([]byte(resp.Body))
		binargs := js.Get("binargs").MustString()
		return binargs, resp.Raw.StatusCode, err
	} else {
		return "", resp.Raw.StatusCode, err
	}
}

/*func main() {

	service := micro.NewService(
		micro.Name("go.micro.srv.v2.asset"),
		micro.Version("2.0.0"),
	)

	service.Init()

	//proto.RegisterUserHandler(service.Server(), new(Asset))
	//user_proto.RegisterUserHandler(service.Server(), new(User))
	proto.RegisterAssetHandler(service.Server(), new(Asset))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}*/
