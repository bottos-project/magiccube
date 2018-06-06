/*Copyright 2017~2022 The Bottos Authors
  This file is part of the Bottos Data Exchange Client
  Created by Developers Team of Bottos.

  This program is free software: you can distribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with Bottos. If not, see <http://www.gnu.org/licenses/>.
*/

package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/bottos-project/magiccube/service/storage/proto"
)

// GetFileUploadURL from db
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

// GetFileUploadStat from minio
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
