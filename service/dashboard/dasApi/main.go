package main

import (
	"encoding/json"
	"github.com/code/bottos/service/dashboard/proto"
	"github.com/micro/go-micro"
	api "github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
	log "github.com/jeanphorn/log4go"
	"github.com/code/bottos/config"
)

type Dashboard struct {
	Client dashboard.DashboardClient
}

func (s *Dashboard) GetAllTxNum(ctx context.Context, req *api.Request, rsp *api.Response) error {
	response, err := s.Client.GetAllTxNum(ctx, &dashboard.GetAllTxNumRequest{})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code": response.Code,
		"data": response.Data,
		"msg":"OK",
	})
	rsp.Body = string(b)
	return nil
}

func (s *Dashboard) GetRecentTxList(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	log.Info(body)
	var dashboardRecentTxList dashboard.GetRecentTxListRequest
	err := json.Unmarshal([]byte(body), &dashboardRecentTxList)
	response, err := s.Client.GetRecentTxList(ctx, &dashboardRecentTxList)
	if err != nil {
		log.Error(err)
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code": response.Code,
		"data": response.Data,
		"msg":"OK",
	})
	rsp.Body = string(b)
	return nil
}

func (s *Dashboard) GetBlockList(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	log.Info(body)
	var blockListRequest dashboard.GetBlockListRequest
	err := json.Unmarshal([]byte(body), &blockListRequest)
	response, err := s.Client.GetBlockList(ctx, &blockListRequest)
	if err != nil {
		log.Error(err)
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code": response.Code,
		"data": response.Data,
		"msg":"OK",
	})
	rsp.Body = string(b)
	return nil
}

func (s *Dashboard) GetNodeInfos(ctx context.Context, req *api.Request, rsp *api.Response) error {
	response, err := s.Client.GetNodeInfos(ctx, &dashboard.GetNodeInfosRequest{})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code": response.Code,
		"data": response.Data,
		"msg":"OK",
	})
	rsp.Body = string(b)
	return nil
}

func (s *Dashboard) GetRequirementNumByDay(ctx context.Context, req *api.Request, rsp *api.Response) error {
	response, err := s.Client.GetRequirementNumByDay(ctx, &dashboard.GetRequirementNumByDayRequest{})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code": response.Code,
		"data": response.Data,
		"msg":"OK",
	})
	rsp.Body = string(b)
	return nil
}

func (s *Dashboard) GetAssetNumByDay(ctx context.Context, req *api.Request, rsp *api.Response) error {
	response, err := s.Client.GetAssetNumByDay(ctx, &dashboard.GetAssetNumByDayRequest{})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code": response.Code,
		"data": response.Data,
		"msg":"OK",
	})
	rsp.Body = string(b)
	return nil
}

func (s *Dashboard) GetAccountNumByDay(ctx context.Context, req *api.Request, rsp *api.Response) error {
	response, err := s.Client.GetAccountNumByDay(ctx, &dashboard.GetAccountNumByDayRequest{})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code": response.Code,
		"data": response.Data,
		"msg":"OK",
	})
	rsp.Body = string(b)
	return nil
}

func (s *Dashboard) GetSumTxAmount(ctx context.Context, req *api.Request, rsp *api.Response) error {
	response, err := s.Client.GetSumTxAmount(ctx, &dashboard.GetSumTxAmountRequest{})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code": response.Code,
		"data": response.Data,
		"msg":"OK",
	})
	rsp.Body = string(b)
	return nil
}

func (s *Dashboard) GetTxNumByDay(ctx context.Context, req *api.Request, rsp *api.Response) error {
	response, err := s.Client.GetTxNumByDay(ctx, &dashboard.GetTxNumByDayRequest{})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code": response.Code,
		"data": response.Data,
		"msg":"OK",
	})
	rsp.Body = string(b)
	return nil
}

func (s *Dashboard) GetTxAmountByDay(ctx context.Context, req *api.Request, rsp *api.Response) error {
	response, err := s.Client.GetTxAmountByDay(ctx, &dashboard.GetTxAmountByDayRequest{})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code": response.Code,
		"data": response.Data,
		"msg":"OK",
	})
	rsp.Body = string(b)
	return nil
}

func (s *Dashboard) GetAllTypeTotal(ctx context.Context, req *api.Request, rsp *api.Response) error {
	response, err := s.Client.GetAllTypeTotal(ctx, &dashboard.GetAllTypeTotalRequest{})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code": response.Code,
		"data": response.Data,
		"msg":"OK",
	})
	rsp.Body = string(b)
	return nil
}

func main() {
	log.LoadConfiguration(config.BASE_LOG_CONF)
	defer log.Close()
	log.LOGGER("dashboard.api")

	service := micro.NewService(
		micro.Name("go.micro.api.v2.dashboard"),
	)

	// parse command line flags
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&Dashboard{Client: dashboard.NewDashboardClient("go.micro.srv.dashboard", service.Client())},
		),
	)
	if err := service.Run(); err != nil {
		log.Exit(err)
	}
}