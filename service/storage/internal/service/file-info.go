package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/code/bottos/service/storage/proto"
)

func (c *StorageService) GetUserFileList(ctx context.Context, request *storage.UserFileListRequest, response *storage.UserFileListResponse) error {

	if request == nil {
		response.Code = 0
		return errors.New("Missing storage request")
	}
	fmt.Println("GetUserFileList")
	files, err := c.mgoRepo.CallGetUserFileList(request.Username)

	if err != nil {
		response.Code = 0
		fmt.Println(err)
		return errors.New("Failed get put url")

	}
	response.FileList = []*storage.File{}
	for _, file := range files {
		fileTag := &storage.File{file.FileName,
			file.FileSize,
			file.FilePolicy,
			file.FileNumber,
			file.FileHash,
			file.AuthorizedStorage}
		response.FileList = append(response.FileList, fileTag)
	}
	response.Code = 1
	return nil
}
