package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/bottos-project/magiccube/service/storage/proto"
)

func (c *StorageService) GetFileUploadURL(ctx context.Context, request *storage.FileUploadRequest, response *storage.FileUploadResponse) error {

	if request == nil {
		response.Result = "404"
		response.Message = "para error"
		return errors.New("Missing storage request")
	}
	fmt.Println(request.Username)
	fmt.Println(request.FileName)
	url, err := c.minioRepo.GetPutURL(request.Username, request.FileName)
	if err != nil {
		response.Result = "404"
		response.Message = "get url failed"
		return errors.New("get PUTURL failed")
	}
	fmt.Println("success")
	response.Result = "200"
	response.Message = "OK"
	response.PresignedPutUrl = url
	return nil
}
func (c *StorageService) GetFileUploadStat(ctx context.Context, request *storage.FileUploadStatRequest, response *storage.FileUploadStatResponse) error {

	if request == nil {
		response.Result = "404"
		response.Message = "para error"
		return errors.New("Missing storage request")
	}
	fmt.Println(request.Username)
	fmt.Println(request.FileName)
	size, err := c.minioRepo.GetPutState(request.Username, request.FileName)
	if err != nil {
		response.Result = "404"
		response.Message = "get url failed"
		return errors.New("get PUTURL failed")
	}
	fmt.Println("success")
	response.Result = "200"
	response.Message = "OK"
	response.Size = size
	return nil
}
