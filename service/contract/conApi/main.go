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
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/bottos-project/magiccube/service/common/data"
	"github.com/bottos-project/magiccube/service/contract/proto"
	"github.com/micro/go-micro"
	api "github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
)

// Contract struct
type Contract struct {
	Client contract.ContractClient
}

// Publish Publish
func (s *Contract) Publish(ctx context.Context, req *api.Request, rsp *api.Response) error {

	return nil
}

// Query Query
func (s *Contract) Query(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var queryRequest contract.QueryRequest
	err := json.Unmarshal([]byte(req.Body), &queryRequest)
	if err != nil {
		log.Error(err)
		return err
	}

	params := `service=bottos&method=Chain.GetAbi&request={"contract":"` + queryRequest.Contract + `"}`
	resp, err := http.Post(data.BASE_URL, "application/x-www-form-urlencoded",
		strings.NewReader(params))
	if err != nil {
		log.Error(err)
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		log.Error(resp.Status)
		return errors.New(string(body))
	}
	if err != nil {
		log.Error(err)
		return err
	}

	rsp.StatusCode = 200
	rsp.Body = string(body)
	return nil
}

func init() {
	logger, err := log.LoggerFromConfigAsFile("./config/con-log.xml")
	if err != nil {
		log.Error(err)
		panic(err)
	}
	defer logger.Flush()
	log.ReplaceLogger(logger)
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.v3.contract"),
	)

	// parse command line flags
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&Contract{Client: contract.NewContractClient("go.micro.srv.v3.contract", service.Client())},
		),
	)
	if err := service.Run(); err != nil {
		os.Exit(1)
	}
}
