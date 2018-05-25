package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/bottos-project/magiccube/service/storage/proto"
)

func (c *StorageService) GetDownLoadURL(ctx context.Context, request *storage.DownLoadRequest, response *storage.DownLoadResponse) error {

	if request == nil {
		response.Result = "404"
		response.Message = "para error"
		return errors.New("Missing storage request")
	}
	fmt.Println("get downloand")
	url, err := c.minioRepo.GetFileDownloadURL(request.Username, request.FileName)

	if err != nil {
		response.Result = "404"
		response.Message = "get url failed"
		fmt.Println(err)
		return errors.New("Failed get put url")

	}
	response.Result = "200"
	response.Message = "OK"
	response.PresignedGetUrl = url
	return nil
}
