package main

import (
	log "github.com/cihub/seelog"
	"encoding/json"
	"github.com/bottos-project/magiccube/service/data/proto"
	"github.com/micro/go-micro"
	api "github.com/micro/micro/api/proto"

	"github.com/bitly/go-simplejson"
	"golang.org/x/net/context"
	"io/ioutil"
	"fmt"

	"net/http"
	"strings"
)

type Data struct {
	Client data.DataClient
}

//global map
var mslice = make(map[string]int)
var msliceip = make(map[string]string)

func (d *Data) FileCheck(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	log.Info("Start File Check !")
	var req1 data.FileCheckRequest
	json.Unmarshal([]byte(body), &req1)

	sliceHash := req1.Hash

	rep, err := d.Client.FileCheck(ctx, &data.FileCheckRequest{
		Hash: sliceHash,
	})
	if err != nil {
		log.Error(err)
	}
	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"result":           rep.Result,
		"message":          rep.Message,
		"merkle_root_hash": rep.MerkleRootHash,
		"is_exist":         rep.IsExist,
	})
	rsp.Body = string(b)
	log.Info("rsp.Body")
	log.Info(rsp.Body)
	return nil
}
func (d *Data) GetFileUploadURL(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	log.Info("Start Get File URL!")
	var req1 data.GetFileUploadURLRequest
	json.Unmarshal([]byte(body), &req1)

	userName := req1.Username
	fileSlice := req1.Slice

	rep, err := d.Client.GetFileUploadURL(ctx, &data.GetFileUploadURLRequest{
		Username: userName,
		Slice:    fileSlice,
	})
	log.Info("GetFileUploadURLResult")
	log.Info(rep)
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"result":  rep.Result,
		"message": rep.Message,
		"url":     rep.Url,
	})
	rsp.Body = string(b)

	return nil
}
func (d *Data) GetUploadProgress(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	log.Info("Start Get Upload Progress !")
	var req1 data.GetUploadProgressRequest
	json.Unmarshal([]byte(body), &req1)

	userName := req1.Username
	fileSlice := req1.Slice
	//1 check cache upload status
	uploadCacheResult, err := d.Client.GetUploadProgress(ctx, &data.GetUploadProgressRequest{
		Username: userName,
		Slice:    fileSlice,
	})
	log.Info("GetUploadProgress Result")
	log.Info(uploadCacheResult)
	if err != nil {
		log.Error(err)
	}
	m := int(uploadCacheResult.ProgressDone)

	//2.1 get slice IPlist
	log.Info("get slice IPlist")
	nodes, err := d.Client.GetFileStorageNode(ctx, &data.GetFileStorageNodeRequest{
		Username: userName,
		Slice:    fileSlice,
	})
	log.Info("GetFileStorageNode Result")
	log.Info(nodes)
	if err != nil {
		log.Info(err)
		return err
	}
	//2.2 storage
	storageOK := 0
	sliceIp := []*data.Ip{}
	sliceIp = nil
	for i := 0; i < m; i++ {
		sguid := fileSlice[i].Sguid
		if mslice[sguid] == 0 {
			//2.2.1 get slice ip
			log.Info("get slice ip")
			Sip := nodes.Ip[i].SnodeIp
			//Sip := "127.0.0.1"
			addr := "http://" + Sip + ":8080/rpc"
			//2.2.2 get slice storage url
            log.Info("get slice storage url")
			params := `service=go.micro.srv.v3.data&method=Data.GetFileStorageURL&request={
					"username":"%s",
					"guid":"%s"}`
			s := fmt.Sprintf(params, userName, sguid)
			resp_body, err := http.Post(addr, "application/x-www-form-urlencoded",
				strings.NewReader(s))
			log.Info("GetFileStorageURL Result")
			log.Info(resp_body)	
			if err != nil {
				log.Info(err)
			}
			defer resp_body.Body.Close()
			body, err := ioutil.ReadAll(resp_body.Body)
			var url string
			if err != nil {
				log.Info(err)
			} else {
				jss, _ := simplejson.NewJson([]byte(body))
				url = jss.Get("url").MustString()

			}
			

			//2.2.3 storage slice file
			log.Info("storage slice file")
			putResult, err := d.Client.PutFile(ctx, &data.PutFileRequest{
				Username: userName,
				Guid:     sguid,
				Url:      url,
			})
			log.Info("PutFile Result")
			log.Info(putResult)
			if err != nil {
				log.Info(err)
				return err
			}
			mslice[sguid] = 1
			msliceip[sguid] = Sip
			nodeTag := &data.Ip{sguid,
			Sip}
			sliceIp = append(sliceIp, nodeTag)
		}else{
			nodeTag := &data.Ip{sguid,
			msliceip[sguid]}
			sliceIp = append(sliceIp, nodeTag)
		}
		storageOK++
	}

	//response
	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"result":              uploadCacheResult.Result,
		"message":             uploadCacheResult.Message,
		"slice_progress_done": uploadCacheResult.SliceProgressDone,
		"slice_progressing":   uploadCacheResult.SliceProgressing,
		"storage_done":         storageOK,
		"storage_ip": 		   sliceIp,
	})
	rsp.Body = string(b)
	return nil

}

func (d *Data) GetStorageIP(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	log.Info("Start Get Storage IP !")
	var req1 data.GetStorageIPRequest
	json.Unmarshal([]byte(body), &req1)

	guid := req1.Guid

	rep, err := d.Client.GetStorageIP(ctx, &data.GetStorageIPRequest{
		Guid: guid,
	})
	log.Info("GetStorageIP Result")
	log.Info(rep)
	if err != nil {
		log.Error(err)
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"result":  rep.Result,
		"message": rep.Message,
		"storage_addr":      rep.StorageAddr,
		"file_name": rep.FileName,
	})
	rsp.Body = string(b)

	return nil
}
func (d *Data) GetFileDownloadURL(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	log.Info("Start Get File Download URL!")
	var req1 data.GetFileDownloadURLRequest

	json.Unmarshal([]byte(body), &req1)

	userName := req1.Username
	guid := req1.Guid
	ip := req1.Ip
	n := len(ip)
	if userName == "" || guid == "" || ip == nil {
		rsp.StatusCode = 200
		b, _ := json.Marshal(map[string]interface{}{
			"result":  "404",
			"message": "Missing  request para",
			"url":     nil,
		})
		rsp.Body = string(b)
		return nil
	}
	for i := 0; i < n; i++ {
		downloadOK := 0
		//1.1 get slice ip
		Sip := ip[i].SnodeIp
		//Sip := "127.0.0.1"
		addr := "http://" + Sip + ":8080/rpc"
		//1.2 get slice storage url
		sguid := ip[i].Sguid
		log.Info("get slice storage url")
		params := `service=go.micro.srv.v3.data&method=Data.GetFileStorageURL&request={
					"username":"%s",
					"guid":"%s"}`
		s := fmt.Sprintf(params, userName, sguid)
		resp_body, err := http.Post(addr, "application/x-www-form-urlencoded",
			strings.NewReader(s))
		log.Info("GetFileStorageURL Result")
	    log.Info(resp_body)
		if err != nil {
			return err
		}
		defer resp_body.Body.Close()
		dbody, err := ioutil.ReadAll(resp_body.Body)
		var url string
		if err != nil {
			return err
		} else {
			jss, _ := simplejson.NewJson([]byte(dbody))
			url = jss.Get("url").MustString()

		}
		//1.3 storage slice file
		log.Info("storage slice file")
		downloadResult, err := d.Client.DownloadFile(ctx, &data.DownloadFileRequest{
			Username: userName,
			Guid:     sguid,
			Url:      url,
		})
		log.Info("DownloadFile Result")
	    log.Info(downloadResult)
		if err != nil {
			log.Error(err)
			return err
		}
		downloadOK++

	}
	//2.composeFile
	log.Info("composeFile")
	d.Client.ComposeFile(ctx, &data.ComposeFileRequest{
		Username: userName,
		Guid:     guid,
		Ip:       ip,
	})

	//3.get download url
	log.Info("get download url")
	rep, err := d.Client.GetFileStorageURL(ctx, &data.GetFileStorageURLRequest{
		Username: userName,
		Guid:     guid,
	})
	log.Info("GetFileStorageURL Result")
	log.Info(rep)
	if err != nil {
		log.Error(err)
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"result":  rep.Result,
		"message": rep.Message,
		"url":     rep.Url,
	})
	rsp.Body = string(b)

	return nil
}

func init() {
	logger, err := log.LoggerFromConfigAsFile("./config/data-log.xml")
	if err != nil {
		log.Error(err)
		panic(err)
	}
	defer logger.Flush()
	log.ReplaceLogger(logger)
}

func main() {

	service := micro.NewService(
		micro.Name("go.micro.api.v3.data"),


	)

	// parse command line flags
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&Data{Client: data.NewDataClient("go.micro.srv.v3.data", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Critical(err)
	}

}
