package main

import (
	log "github.com/cihub/seelog"
	"encoding/json"
	"github.com/bottos-project/magiccube/service/dashboard/proto"
	"github.com/micro/go-micro"
	api "github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
	"os"
	errcode "github.com/bottos-project/magiccube/error"
)

type Dashboard struct {
	Client dashboard.DashboardClient
}

func (s *Dashboard) GetTxNum(ctx context.Context, req *api.Request, rsp *api.Response) error {
	response, err := s.Client.GetTxNum(ctx, &dashboard.GetTxNumRequest{})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	rsp.Body = errcode.Return(response)
	return nil
}

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

func (s *Dashboard) GetNodeInfos(ctx context.Context, req *api.Request, rsp *api.Response) error {
	response, err := s.Client.GetNodeInfos(ctx, &dashboard.GetNodeInfosRequest{})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	rsp.Body = errcode.Return(response)
	return nil
}

func (s *Dashboard) GetRequirementNumByDay(ctx context.Context, req *api.Request, rsp *api.Response) error {
	response, err := s.Client.GetRequirementNumByDay(ctx, &dashboard.GetRequirementNumByDayRequest{})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	rsp.Body = errcode.Return(response)
	return nil
}

func (s *Dashboard) GetAssetNumByDay(ctx context.Context, req *api.Request, rsp *api.Response) error {
	response, err := s.Client.GetAssetNumByDay(ctx, &dashboard.GetAssetNumByDayRequest{})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	rsp.Body = errcode.Return(response)
	return nil
}

func (s *Dashboard) GetAccountNumByDay(ctx context.Context, req *api.Request, rsp *api.Response) error {
	response, err := s.Client.GetAccountNumByDay(ctx, &dashboard.GetAccountNumByDayRequest{})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	rsp.Body = errcode.Return(response)
	return nil
}

func (s *Dashboard) GetTxAmount(ctx context.Context, req *api.Request, rsp *api.Response) error {
	response, err := s.Client.GetTxAmount(ctx, &dashboard.GetTxAmountRequest{})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	rsp.Body = errcode.Return(response)
	return nil
}

func (s *Dashboard) GetTxNumByDay(ctx context.Context, req *api.Request, rsp *api.Response) error {
	response, err := s.Client.GetTxNumByDay(ctx, &dashboard.GetTxNumByDayRequest{})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	rsp.Body = errcode.Return(response)
	return nil
}

func (s *Dashboard) GetTxAmountByDay(ctx context.Context, req *api.Request, rsp *api.Response) error {
	response, err := s.Client.GetTxAmountByDay(ctx, &dashboard.GetTxAmountByDayRequest{})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	rsp.Body = errcode.Return(response)
	return nil
}

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
	if err != nil{
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