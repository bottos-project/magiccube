package main

import (
	"bytes"
	"fmt"
	//"github.com/bitly/go-simplejson"
	baseConfig "github.com/bottos-project/bottos/config"
	"github.com/bottos-project/bottos/service/data/internal/platform/config"
	hash "github.com/bottos-project/bottos/service/data/internal/platform/hash"
	"github.com/bottos-project/bottos/service/data/internal/platform/minio"
	"github.com/bottos-project/bottos/service/data/internal/platform/mongodb"
	proto "github.com/bottos-project/bottos/service/data/proto"
	util "github.com/bottos-project/bottos/service/data/util"
	basicMinio "github.com/minio/minio-go"
	//"github.com/minio/minio-go/pkg/encrypt"
	//	storage "github.com/bottos-project/bottos/service/storage/proto"
	"errors"
	//log "github.com/jeanphorn/log4go"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
	"io/ioutil"
	"math/rand"
	"net/http"
	//	"strconv"
	"io"
	//"os"
	//"strings"
	"time"
)

const (
	BASE_CHAIN_IP           = baseConfig.BASE_CHAIN_IP
	BASE_URL                = baseConfig.BASE_CHAIN_URL
	GET_INFO_URL            = BASE_URL + "v1/chain/get_info"
	GET_BLOCK_URL           = BASE_URL + "v1/chain/get_block"
	ABI_JSON_TO_BIN_URL     = BASE_URL + "v1/chain/abi_json_to_bin"
	PUSH_TRANSACTION_URL    = BASE_URL + "v1/chain/push_transaction"
	GET_TABLE_ROW_BY_STRING = BASE_URL + "v1/chain/get_table_row_by_string_key"
	STORAGE_RPC_URL         = baseConfig.BASE_RPC
)

type DataService struct {
	minioRepo minioRepository
	dbRepo    dbRepository
	mgoRepo   mgoRepository
}
type minioRepository interface {
	GetCacheURL(username string, objectName string) (string, error)
	GetFileDownloadURL(username string, objectName string) (string, error)
	GetCacheFile(username string, objectName string) (*basicMinio.Object, error)
	PutFile(username string, objectName string, reader io.Reader, objectSize int64) (int64, error)
	ComposeFile(dst basicMinio.DestinationInfo, srcs []basicMinio.SourceInfo) error
	GetPutState(username string, objectName string) (int64, error)
}
type dbRepository interface {
}
type mgoRepository interface {
	CallIsDataExists(slicehash string) (uint64, error)
	CallNodeRequest(seedip string) (*util.NodeDBInfo, error)
	CallDataSliceIPRequest(sguid string) (*util.DataDBInfo, error)
}

func NewDataService(minioRepo minioRepository, mgodb mgoRepository) proto.DataHandler {
	return &DataService{minioRepo: minioRepo, mgoRepo: mgodb}
}

func (d *DataService) FileCheck(ctx context.Context, req *proto.FileCheckRequest, rsp *proto.FileCheckResponse) error {

	fmt.Println("Start Check File!")
	if req == nil {
		rsp.Result = 404
		rsp.Message = "para error"
		return errors.New("Missing data request")
	}
	fmt.Println("Check File")

	sliceHash := req.Hash
	fmt.Println("req.Hash")
	fmt.Println(req.Hash)
	var hs []hash.Hash
	for _, filehash := range sliceHash {
		sfilehash := filehash.Hash

		shash := hash.HexToHash(sfilehash)
		hs = append(hs, shash)
	}
	fmt.Println("hs")
	fmt.Println(hs)
	MerkleRootHash := hash.ComputeMerkleRootHash(hs)
	root := MerkleRootHash.ToHexString()
	fmt.Println("root")
	fmt.Println(root)
	fmt.Println("root")
	fmt.Println(root)
	isSlicefileExist, err := d.mgoRepo.CallIsDataExists(root)
	if err != nil {
		rsp.Result = 404
		rsp.Message = "file check failed"
		fmt.Println(err)
		return errors.New("Failed check file")
	}
	rsp.Result = 200
	rsp.Message = "OK"
	rsp.MerkleRootHash = root
	rsp.IsExist = isSlicefileExist
	return nil
}

func (d *DataService) GetFileUploadURL(ctx context.Context, req *proto.GetFileUploadURLRequest, rsp *proto.GetFileUploadURLResponse) error {

	fmt.Println("Start Get File URL!")
	if req == nil {
		rsp.Result = 404
		rsp.Message = "para error"
		return errors.New("Missing storage request")
	}

	userName := req.Username
	fileSlice := req.Slice
	rsp.Url = []*proto.Url{}
	fmt.Println("get FileUploadURL")

	for _, slice := range fileSlice {
		fmt.Println("userName")
		fmt.Println(userName)
		fmt.Println("slice.Sguid")
		fmt.Println(slice.Sguid)
		cacheUrl, err := d.minioRepo.GetCacheURL(userName, slice.Sguid)
		if err != nil {
			rsp.Result = 404
			rsp.Message = "get url failed"
			fmt.Println(err)
			return errors.New("Failed get put url")
		}

		urlTag := &proto.Url{slice.Sguid,
			cacheUrl}
		rsp.Url = append(rsp.Url, urlTag)
	}

	rsp.Result = 200
	rsp.Message = "OK"
	return nil
}
func (d *DataService) GetFileDownloadURL(ctx context.Context, req *proto.GetFileDownloadURLRequest, rsp *proto.GetFileDownloadURLResponse) error {

	fmt.Println("Start Get File URL!")
	if req == nil {
		rsp.Result = 404
		rsp.Message = "para error"
		return errors.New("Missing storage request")
	}
	fmt.Println("get FileDownloadURL")

	userName := req.Username
	guid := req.Guid

	url, err := d.minioRepo.GetFileDownloadURL(userName, guid)
	if err != nil {
		rsp.Result = 404
		rsp.Message = "get url failed"
		fmt.Println(err)
		return errors.New("Failed get download url")
	}

	rsp.Url = url
	rsp.Result = 200
	rsp.Message = "OK"
	return nil
}

func (d *DataService) GetUploadProgress(ctx context.Context, req *proto.GetUploadProgressRequest, rsp *proto.GetUploadProgressResponse) error {
	fmt.Println("Start Get Upload Progress!")
	if req == nil {
		rsp.Result = 404
		rsp.Message = "para error"
		return errors.New("Missing data request")
	}

	userName := req.Username
	//guid := req.Guid
	fileSlice := req.Slice

	rsp.SliceProgressDone = []*proto.Slice{}
	rsp.SliceProgressing = []*proto.Slice{}
	var i int = 0
	var j int = 0
	for _, slice := range fileSlice {
		result, err := d.minioRepo.GetPutState(userName, slice.Sguid)
		fmt.Println("result")
		fmt.Println(result)
		if err != nil {
			j++
			slice1Tag := slice
			rsp.SliceProgressing = append(rsp.SliceProgressing, slice1Tag)
		}
		if result != 0 {
			i++
			slice2Tag := slice
			rsp.SliceProgressDone = append(rsp.SliceProgressDone, slice2Tag)
		}

	}

	fmt.Println("success")

	rsp.Result = 200
	rsp.Message = "OK"
	rsp.ProgressDone = uint64(i)
	rsp.Progressing = uint64(j)
	return nil

}

func (d *DataService) GetFileStorageNode(ctx context.Context, req *proto.GetFileStorageNodeRequest, rsp *proto.GetFileStorageNodeResponse) error {

	fmt.Println("Start Get File Storage Node!")
	if req == nil {
		rsp.Result = 404
		rsp.Message = "para error"
		return errors.New("Missing storage node request")
	}
	fmt.Println("get FileStorageNode")

	fileSlice := req.Slice
	rsp.Ip = []*proto.Ip{}

	nodeInfo, err := d.mgoRepo.CallNodeRequest(BASE_CHAIN_IP)
	if err != nil {
		rsp.Result = 404
		rsp.Message = "get node failed"
		fmt.Println(err)
		return errors.New("Failed get put node")
	}

	i := 0
	n := len(nodeInfo.SlaveIP)
	k := rand.Intn(n)
	for _, slice := range fileSlice {
		j := (i + k) % n
		node := nodeInfo.SlaveIP[j]
		nodeTag := &proto.Ip{slice.Sguid,
			node}
		rsp.Ip = append(rsp.Ip, nodeTag)
		i++
	}
	rsp.Result = 200
	rsp.Message = "OK"
	return nil
}

func (d *DataService) GetFileStorageURL(ctx context.Context, req *proto.GetFileStorageURLRequest, rsp *proto.GetFileStorageURLResponse) error {

	fmt.Println("Start Get File Storage URL!")
	if req == nil {
		rsp.Result = 404
		rsp.Message = "para error"
		return errors.New("Missing storage request")
	}
	fmt.Println("get FileStorageURL")

	userName := req.Username
	guid := req.Guid

	storageUrl, err := d.minioRepo.GetFileDownloadURL(userName, guid)
	if err != nil {
		rsp.Result = 404
		rsp.Message = "get url failed"
		fmt.Println(err)
		return errors.New("Failed get put url")
	}

	rsp.Url = storageUrl

	rsp.Result = 200
	rsp.Message = "OK"
	return nil
}
func (d *DataService) PutFile(ctx context.Context, req *proto.PutFileRequest, rsp *proto.PutFileResponse) error {

	fmt.Println("Start Put file!")
	if req == nil {
		rsp.Result = 404
		rsp.Message = "para error"
		return errors.New("Missing storage request")
	}
	fmt.Println("put file")

	url := req.Url
	userName := req.Username
	guid := req.Guid

	client := &http.Client{}
	//create form data
	bodyBuf := &bytes.Buffer{}
	//get cache file
	file, err := d.minioRepo.GetCacheFile(userName, guid)
	if err != nil {
		fmt.Println("error read file")
		return nil
	}
	//iocopy
	_, err = io.Copy(bodyBuf, file)
	if err != nil {
		return nil
	}
	req_body, err := http.NewRequest("PUT", url, bodyBuf)
	if err != nil {
		return nil
	}
	req_body.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
	req_body.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
	req_body.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	req_body.Header.Set("Cache-Control", "max-age=0")
	req_body.Header.Set("Connection", "keep-alive")
	resp, err := client.Do(req_body)
	defer req_body.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	fmt.Println("resp_body")
	fmt.Println(resp_body)
	rsp.Result = 200
	rsp.Message = "OK"
	return nil
}
func (d *DataService) DownloadFile(ctx context.Context, req *proto.DownloadFileRequest, rsp *proto.DownloadFileResponse) error {
	fmt.Println("Start download file!")
	if req == nil {
		rsp.Result = 404
		rsp.Message = "para error"
		return errors.New("Missing storage request")
	}
	fmt.Println("download file")

	url := req.Url
	userName := req.Username
	guid := req.Guid

	client := &http.Client{}
	req_body, _ := http.NewRequest(http.MethodGet, url, nil)
	req_body.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
	req_body.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
	req_body.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	req_body.Header.Set("Cache-Control", "max-age=0")
	req_body.Header.Set("Connection", "keep-alive")

	resp_http, err := client.Do(req_body)
	defer resp_http.Body.Close()
	fmt.Println("download success")
	//**start
	body_http, err := ioutil.ReadAll(resp_http.Body)
	if err != nil {
		fmt.Println(err)
	}
	file1 := bytes.NewReader(body_http)
	fileSize := int64(len(body_http))

	//**end
	/////getfile
	/*
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("resp_http")
		fmt.Println(resp_http)
		file, err := os.Create(guid)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("file")
		fmt.Println(file)
		defer file.Close()

		_, err = io.Copy(file, resp_http.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("download success")
		////getfile
		////putfile

		file1, err := os.Open(guid)
		if err != nil {
			fmt.Println(err)
		}
		defer file1.Close()

		fileStat, err := file1.Stat()
		if err != nil {
			fmt.Println(err)
		}*/
	////putfile
	fmt.Println("start upload")
	n, err := d.minioRepo.PutFile(userName, guid, file1, fileSize)
	if err != nil {
		fmt.Println(err)
		//return err
	}
	fmt.Println("Successfully uploaded bytes: ", n)
	fmt.Println("upload success")

	rsp.Result = 200
	rsp.Message = "OK"
	return nil
}
func (d *DataService) ComposeFile(ctx context.Context, req *proto.ComposeFileRequest, rsp *proto.ComposeFileResponse) error {

	fmt.Println("Start compose file!")
	if req == nil {
		fmt.Println("Missing file request")

	}
	fmt.Println("compose file")

	userName := req.Username
	guid := req.Guid
	ip := req.Ip

	//***start
	// Create slice of sources.
	var srcs = []basicMinio.SourceInfo{}

	//sseSrc := encrypt.DefaultPBKDF([]byte("password"), []byte("salt"))
	for _, sip := range ip {

		sguid := sip.Sguid
		src := basicMinio.NewSourceInfo(userName, sguid, nil)
		//src.SetMatchETagCond("31624deb84149d2f8ef9c385918b653a")
		srcs = append(srcs, src)

	}

	//sseDst := encrypt.DefaultPBKDF([]byte("new-password"), []byte("new-salt"))
	dst, err := basicMinio.NewDestinationInfo(userName, guid, nil, nil)
	if err != nil {
		fmt.Println(err)
	}

	err = d.minioRepo.ComposeFile(dst, srcs)
	fmt.Println(err)
	fmt.Println("Composed object successfully.")
	return nil

}

func (d *DataService) GetStorageIP(ctx context.Context, req *proto.GetStorageIPRequest, rsp *proto.GetStorageIPResponse) error {

	fmt.Println("Start Get Storage IP!")
	if req == nil {
		rsp.Result = 404
		rsp.Message = "para error"
		return errors.New("Missing storage node request")
	}
	fmt.Println("get StorageIP")

	guid := req.Guid
	rsp.Ip = []*proto.Ip{}

	DataInfo, err := d.mgoRepo.CallDataSliceIPRequest(guid)
	if err != nil {
		rsp.Result = 404
		rsp.Message = "get node failed"
		fmt.Println(err)
		return errors.New("Failed get storage node")
	}
	for i := 0; i < len(DataInfo.Fslice); i++ {
		fs := DataInfo.Fslice[i]
		fmt.Println("fs")
		fmt.Println(fs)
		fmt.Println("fs[0]")
		fmt.Println(fs[0])
		fmt.Println("fs[1]")
		fmt.Println(fs[1])
		nodeTag := &proto.Ip{fs[0],
			fs[1]}
		rsp.Ip = append(rsp.Ip, nodeTag)
	}

	rsp.Result = 200
	rsp.Message = "OK"
	return nil
}

func main() {

	svc := micro.NewService(
		micro.Name("go.micro.srv.v2.data"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Version(config.Version),
	)

	svc.Init()
	minioDataRepository := minio.NewMinioRepository(baseConfig.BASE_MINIO_ADDR, baseConfig.BASE_MINIO_ACCESS_KEY, baseConfig.BASE_MINIO_SECRET_KEY)
	mgoRepository := mongodb.NewMongoRepository(baseConfig.BASE_MONGODB_ADDR)

	repo := NewDataService(minioDataRepository, mgoRepository)
	proto.RegisterDataHandler(svc.Server(), repo)
	if err := svc.Run(); err != nil {
		panic(err)
	}

}
