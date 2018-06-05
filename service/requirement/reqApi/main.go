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
	"encoding/json"
	"os"

	errcode "github.com/bottos-project/magiccube/error"
	sign "github.com/bottos-project/magiccube/service/common/signature"
	"github.com/bottos-project/magiccube/service/requirement/proto"
	log "github.com/cihub/seelog"
	"github.com/micro/go-micro"
	api "github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
)

// Requirement struct
type Requirement struct {
	Client requirement.RequirementClient
}

// Publish requirement
func (s *Requirement) Publish(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.StatusCode = 200

	//验签
	isTrue, err := sign.PushVerifySign(req.Body)
	if !isTrue {
		rsp.Body = errcode.ReturnError(1000, err)
		return nil
	}

	var publishRequest requirement.PublishRequest
	err = json.Unmarshal([]byte(req.Body), &publishRequest)
	if err != nil {
		log.Error(err)
		return err
	}

	response, err := s.Client.Publish(ctx, &publishRequest)
	if err != nil {
		return err
	}

	rsp.Body = errcode.Return(response)
	return nil
}

// QueryById on chain
func (s *Requirement) QueryById(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.StatusCode = 200
	body := req.Body
	var queryByIdRequest requirement.QueryByIdRequest
	err := json.Unmarshal([]byte(body), &queryByIdRequest)
	if err != nil {
		log.Error(err)
		return err
	}
	response, err := s.Client.QueryById(ctx, &queryByIdRequest)
	if err != nil {
		log.Error(err)
		return err
	}

	rsp.Body = errcode.Return(response)
	return nil
}

// Query on chain
func (s *Requirement) Query(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.StatusCode = 200
	body := req.Body
	log.Info(body)
	var requirementQuery requirement.QueryRequest
	err := json.Unmarshal([]byte(body), &requirementQuery)
	if err != nil {
		log.Error(err)
		return err
	}
	requirementQuery.Username = ""
	response, err := s.Client.Query(ctx, &requirementQuery)
	if err != nil {
		log.Error(err)
		return err
	}

	rsp.Body = errcode.Return(response)
	return nil
}

// QueryByUsername on chain
func (s *Requirement) QueryByUsername(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.StatusCode = 200
	body := req.Body
	var requirementQuery requirement.QueryRequest
	err := json.Unmarshal([]byte(body), &requirementQuery)
	if err != nil {
		log.Error(err)
		return err
	}

	//验签
	isTrue, err := sign.QueryVerifySign(req.Body)
	if !isTrue {
		rsp.Body = errcode.ReturnError(1000, err)
		return nil
	}

	response, err := s.Client.Query(ctx, &requirementQuery)
	if err != nil {
		log.Error(err)
		return err
	}

	rsp.Body = errcode.Return(response)
	return nil
}

func init() {
	logger, err := log.LoggerFromConfigAsFile("./config/req-log.xml")
	if err != nil {
		log.Error(err)
		panic(err)
	}
	defer logger.Flush()
	log.ReplaceLogger(logger)
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.v3.requirement"),
	)

	// parse command line flags
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&Requirement{Client: requirement.NewRequirementClient("go.micro.srv.v3.requirement", service.Client())},
		),
	)
	if err := service.Run(); err != nil {
		os.Exit(1)
	}
}
