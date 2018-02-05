package service

import (
	"github.com/code/bottos/service/storemanagement/proto"
//	"github.com/micro/go-micro/errors"
	"golang.org/x/net/context"

	"fmt"
)

//	"github.com/micro/go-micro/errors"
//	"golang.org/x/net/context"

type storageService struct {
	storageRepo storageRepository
}

type storageRepository interface {
//	GetTx(txid string, account string) (string, error)
	GetPutURL(username string, objectName string) (string, error)
	GetFileDownloadURL(username string, objectName string) (string, error)
}

func NewStorageService(storageRepo storageRepository) storemanagement.StoragemanagementHandler {
	return &storageService{
		storageRepo: storageRepo,
	}
}
func (c *storageService) GetTx(ctx context.Context, request *storemanagement.Request, response *storemanagement.Response) error {

	if request == nil {
		//	response.Message = "para error"
		return nil //errors.BadRequest("", "Missing storage request")
	}
	fmt.Print(request)
	url :="http:"
	//url, err := c.storageRepo.GetTx(request.txid, request.account)
	//if err != nil
	{
		//	response.Message = "get url failed"
		return nil //errors.InternalServerError("", "Failed get put url: %s", err.Error())

	}
	fmt.Print(url)
	//response.Message = "OK"
	return nil
}


func (c *storageService) GetFileUploadURL(ctx context.Context, request *storemanagement.FileUploadRequest, response *storemanagement.FileUploadResponse) error {

	if request == nil {
		response.Result = "404"
		response.Message = "para error"
		return nil //errors.BadRequest("", "Missing storage request")
	}

	url, err := c.storageRepo.GetPutURL(request.Username, request.FileName)
	if err != nil	{
		response.Result = "404"
		response.Message = "get url failed"
		return nil //err //  .InternalServerError("", "Failed get put url: %s", err.Error())

	}
	response.Result = "200"
	response.Message = "OK"
	response.PresignedPutURL = url
	return nil
}

func (c *storageService) GetDownLoadURL(ctx context.Context, request *storemanagement.DownLoadRequest, response *storemanagement.DownLoadResponse) error {

	if request == nil {
		response.Result = "404"
		response.Message = "para error"
		return nil //errors.BadRequest("", "Missing storage request")
	}
	url, err := c.storageRepo.GetFileDownloadURL(request.Username, request.FileName)

	if err != nil	{
		response.Result = "404"
		response.Message = "get url failed"
		return nil //err //  .InternalServerError("", "Failed get put url: %s", err.Error())

	}
	response.Result = "200"
	response.Message = "OK"
	response.PresignedGetURL = url
	return nil
}
