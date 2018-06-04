/*Copyright 2017~2022 The Bottos Authors
  This file is part of the Bottos Service Layer
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
package main

import (
	log "github.com/cihub/seelog"
	"encoding/json"
	"github.com/bottos-project/bottos/service/node/proto"
	"github.com/micro/go-micro"
	api "github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
	"github.com/mojocn/base64Captcha"
	"os"
	"github.com/bottos-project/bottos/config"
	"regexp"
	errcode "github.com/bottos-project/bottos/error"
	sign "github.com/bottos-project/bottos/service/common/signature"
)

type nodeTrxInfo struct {
	Client node.NodeClient
}

func (u *User) Register(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.StatusCode = 200
	log.Info(req.Body)
	var registerRequest user.RegisterRequests
	err := json.Unmarshal([]byte(req.Body), &registerRequest)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Info(registerRequest)

	match,err :=regexp.MatchString("^[a-km-z][a-km-z1-9]{2,15}$",registerRequest.nodeUUIDInfo.UserName)
	if err != nil {
		log.Error(err)
		return err
	}
	if !match {
		rsp.Body = errcode.ReturnError(1002)
		return nil
	}

	user_json_buf, err := json.Marshal(registerRequest.SignInfo)
	if err != nil {
		log.Error(err)
		return err
	}

	is_true, err := sign.PushVerifySign(string(user_json_buf), registerRequest.nodeUUIDInfo.PubKey)
	if !is_true {
		rsp.Body = errcode.ReturnError(1000, err)
		return nil
	}

	response, err := u.Client.Register(ctx, &registerRequest)
	if err != nil {
		return err
	}
	rsp.Body = errcode.Return(response)
	return nil
}
