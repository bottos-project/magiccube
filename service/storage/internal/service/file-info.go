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
