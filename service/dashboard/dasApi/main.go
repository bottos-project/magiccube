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
	"github.com/bottos-project/magiccube/service/dashboard/proto"
	log "github.com/cihub/seelog"
	"github.com/micro/go-micro"
	api "github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
)

// Dashboard struct
type Dashboard struct {
	Client dashboard.DashboardClient
}

// GetTxNum on chain
func (s *Dashboard) GetTxNum(ctx context.Context, req *api.Request, rsp *api.Response) error {
	response, err := s.Client.GetTxNum(ctx, &dashboard.GetTxNumRequest{})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	rsp.Body = errcode.Return(response)
	return nil
}

// GetTxList on chain
func (s *Dashboard) GetTxList(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.StatusCode = 200
	body := req.Body
	log.Info(body)
	var getTxList dashboard.GetTxListRequest
	err := json.Unmarshal([]byte(body), &getTxList)
	response, err := s.Client.GetTxList(ctx, &getTxList)
	if err != nil {
		log.Error(err)
	}
	rsp.Body = errcode.Return(response)
	return nil
}

// GetBlockList on chain
func (s *Dashboard) GetBlockList(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.StatusCode = 200
	body := req.Body
	log.Info(body)
	var blockListRequest dashboard.GetBlockListRequest
	err := json.Unmarshal([]byte(body), &blockListRequest)
	response, err := s.Client.GetBlockList(ctx, &blockListRequest)
	if err != nil {
		log.Error(err)
	}

	rsp.Body = errcode.Return(response)
	return nil
}

// GetBlockInfo on chain
func (s *Dashboard) GetBlockInfo(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.StatusCode = 200
	body := req.Body
	log.Info(body)
	var getBlockInfoRequest dashboard.GetBlockInfoRequest
	err := json.Unmarshal([]byte(body), &getBlockInfoRequest)
	response, err := s.Client.GetBlockInfo(ctx, &getBlockInfoRequest)
	if err != nil {
		log.Error(err)
	}

	rsp.Body = errcode.Return(response)
	return nil
}

// GetNodeInfos on chain
func (s *Dashboard) GetNodeInfos(ctx context.Context, req *api.Request, rsp *api.Response) error {
	//response, err := s.Client.GetNodeInfos(ctx, &dashboard.GetNodeInfosRequest{})
	//if err != nil {
	//	return err
	//}

	rsp.StatusCode = 200
	var data []*dashboard.NodeInfoData

	data = append(data, &dashboard.NodeInfoData{
		Ip:   "127.0.0.1",
		Port: "8080",
		Position: &dashboard.Position{
			Longitude: 116.240405,
			Latitude:  39.953014,
		},
	})

	data = append(data, &dashboard.NodeInfoData{
		Ip:   "127.0.0.1",
		Port: "8080",
		Position: &dashboard.Position{
			Longitude: 117.006668,
			Latitude:  39.959003,
		},
	})

	data = append(data, &dashboard.NodeInfoData{
		Ip:   "127.0.0.1",
		Port: "8080",
		Position: &dashboard.Position{
			Longitude: 116.471118,
			Latitude:  39.607907,
		},
	})

	data = append(data, &dashboard.NodeInfoData{
		Ip:   "127.0.0.1",
		Port: "8080",
		Position: &dashboard.Position{
			Longitude: 116.214313,
			Latitude:  40.08658,
		},
	})

	data = append(data, &dashboard.NodeInfoData{
		Ip:   "127.0.0.1",
		Port: "8080",
		Position: &dashboard.Position{
			Longitude: 121.295489,
			Latitude:  31.250492,
		},
	})

	data = append(data, &dashboard.NodeInfoData{
		Ip:   "127.0.0.1",
		Port: "8080",
		Position: &dashboard.Position{
			Longitude: 121.141681,
			Latitude:  31.259884,
		},
	})

	data = append(data, &dashboard.NodeInfoData{
		Ip:   "127.0.0.1",
		Port: "8080",
		Position: &dashboard.Position{
			Longitude: 121.652545,
			Latitude:  31.250492,
		},
	})

	data = append(data, &dashboard.NodeInfoData{
		Ip:   "127.0.0.1",
		Port: "8080",
		Position: &dashboard.Position{
			Longitude: 121.545428,
			Latitude:  31.161224,
		},
	})

	data = append(data, &dashboard.NodeInfoData{
		Ip:   "127.0.0.1",
		Port: "8080",
		Position: &dashboard.Position{
			Longitude: 120.147418,
			Latitude:  30.154743,
		},
	})

	data = append(data, &dashboard.NodeInfoData{
		Ip:   "127.0.0.1",
		Port: "8080",
		Position: &dashboard.Position{
			Longitude: 119.949903,
			Latitude:  30.166617,
		},
	})

	rsp.Body = errcode.Return(dashboard.GetNodeInfosResponse{
		Code: 1,
		Data: data,
	})
	return nil
}

// GetRequirementNumByDay on chain
func (s *Dashboard) GetRequirementNumByDay(ctx context.Context, req *api.Request, rsp *api.Response) error {
	response, err := s.Client.GetRequirementNumByDay(ctx, &dashboard.GetRequirementNumByDayRequest{})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	rsp.Body = errcode.Return(response)
	return nil
}

// GetAssetNumByDay on chain
func (s *Dashboard) GetAssetNumByDay(ctx context.Context, req *api.Request, rsp *api.Response) error {
	response, err := s.Client.GetAssetNumByDay(ctx, &dashboard.GetAssetNumByDayRequest{})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	rsp.Body = errcode.Return(response)
	return nil
}

// GetAccountNumByDay on chain
func (s *Dashboard) GetAccountNumByDay(ctx context.Context, req *api.Request, rsp *api.Response) error {
	response, err := s.Client.GetAccountNumByDay(ctx, &dashboard.GetAccountNumByDayRequest{})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	rsp.Body = errcode.Return(response)
	return nil
}

// GetTxAmount  on chain
func (s *Dashboard) GetTxAmount(ctx context.Context, req *api.Request, rsp *api.Response) error {
	response, err := s.Client.GetTxAmount(ctx, &dashboard.GetTxAmountRequest{})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	rsp.Body = errcode.Return(response)
	return nil
}

// GetTxNumByDay on chain
func (s *Dashboard) GetTxNumByDay(ctx context.Context, req *api.Request, rsp *api.Response) error {
	response, err := s.Client.GetTxNumByDay(ctx, &dashboard.GetTxNumByDayRequest{})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	rsp.Body = errcode.Return(response)
	return nil
}

// GetTxAmountByDay on chain
func (s *Dashboard) GetTxAmountByDay(ctx context.Context, req *api.Request, rsp *api.Response) error {
	response, err := s.Client.GetTxAmountByDay(ctx, &dashboard.GetTxAmountByDayRequest{})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	rsp.Body = errcode.Return(response)
	return nil
}

// GetAllTypeTotal on chain
func (s *Dashboard) GetAllTypeTotal(ctx context.Context, req *api.Request, rsp *api.Response) error {
	response, err := s.Client.GetAllTypeTotal(ctx, &dashboard.GetAllTypeTotalRequest{})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	rsp.Body = errcode.Return(response)
	return nil
}

func init() {
	logger, err := log.LoggerFromConfigAsFile("./config/dash-log.xml")
	if err != nil {
		log.Error(err)
		panic(err)
	}
	defer logger.Flush()
	log.ReplaceLogger(logger)
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.v3.dashboard"),
	)

	// parse command line flags
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&Dashboard{Client: dashboard.NewDashboardClient("go.micro.srv.v3.dashboard", service.Client())},
		),
	)
	if err := service.Run(); err != nil {
		os.Exit(1)
	}
}
