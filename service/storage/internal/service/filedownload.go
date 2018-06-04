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
