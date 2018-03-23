package main

import (
	log "github.com/jeanphorn/log4go"
	"github.com/micro/go-micro"
	proto "github.com/code/bottos/service/asset/proto"
	"golang.org/x/net/context"
	"github.com/mikemintang/go-curl"
	"github.com/bitly/go-simplejson"
	"time"
	storage "github.com/code/bottos/service/storage/proto"
	"github.com/micro/go-micro/client"
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"fmt"
	"strconv"
	"golang.org/x/net/html/atom"
	"encoding/json"
	"github.com/code/bottos/config"
	"gopkg.in/mgo.v2/bson"
	"github.com/code/bottos/service/bean"
	"github.com/code/bottos/tools/db/mongodb"
	"errors"
)

const (
	BASE_URL             = config.BASE_URL
	GET_INFO_URL         = BASE_URL + "v1/chain/get_info"
	GET_BLOCK_URL        = BASE_URL + "v1/chain/get_block"
	ABI_JSON_TO_BIN_URL  = BASE_URL + "v1/chain/abi_json_to_bin"
	PUSH_TRANSACTION_URL = BASE_URL + "v1/chain/push_transaction"
	GET_TABLE_ROW        = BASE_URL + "v1/chain/get_table_row_by_string_key"
	STORAGE_RPC_URL      = config.BASE_RPC
)

type Asset struct{}

func (u *Asset) GetFileUploadURL(ctx context.Context, req *proto.GetFileUploadURLRequest, rsp *proto.GetFileUploadURLResponse) error {
	log.Info("Start Get File URL!")
	start_time := time.Now().UnixNano() / int64(time.Millisecond)
	log.Info("reqBody:" + req.PostBody)
	//dataBody, signValue, userName := "fd","","13"
	dataBody, signValue, userName := GetSignAndData(req.PostBody)
	//log.Info(userName)
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
	//log.Info(ok)
	//get strore Address
	js, _ := simplejson.NewJson([]byte(req.PostBody))
	log.Info("js", js)

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


func main() {
	log.LoadConfiguration(config.BASE_LOG_CONF)
	defer log.Close()
	log.LOGGER("asset.srv")

	service := micro.NewService(
		micro.Name("go.micro.srv.v2.asset"),
		micro.Version("2.0.0"),
	)

	service.Init()

	//proto.RegisterUserHandler(service.Server(), new(Asset))
	//user_proto.RegisterUserHandler(service.Server(), new(User))
	proto.RegisterAssetHandler(service.Server(), new(Asset))

	if err := service.Run(); err != nil {
		log.Exit(err)
	}
}
