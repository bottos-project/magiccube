package main

import (
	log "github.com/cihub/seelog"

	"github.com/bottos-project/magiccube/service/contract/proto"
	"github.com/micro/go-micro"
	api "github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
	"os"
)

type Contract struct {
	Client contract.ContractClient
}

func (s *Contract) Publish(ctx context.Context, req *api.Request, rsp *api.Response) error {
	//rsp.StatusCode = 200
	//
	////验签
	//is_true, err := sign.PushVerifySign(req.Body)
	//if !is_true {
	//	rsp.Body = errcode.ReturnError(1000, err)
	//	return nil
	//}
	//
	//var publishRequest contract.PublishRequest
	//err = json.Unmarshal([]byte(req.Body), &publishRequest)
	//if err != nil {
	//	log.Error(err)
	//	return err
	//}
	//
	//response, err := s.Client.Publish(ctx, &publishRequest)
	//if err != nil {
	//	return err
	//}
	//
	//rsp.Body = errcode.Return(response)
	return nil
}

func (s *Contract) Query(ctx context.Context, req *api.Request, rsp *api.Response) error {
	//rsp.StatusCode = 200
	//body := req.Body
	//log.Info(body)
	//var requirementQuery requirement.QueryRequest
	//err := json.Unmarshal([]byte(body), &requirementQuery)
	//if err != nil {
	//	log.Error(err)
	//	return err
	//}
	//requirementQuery.Username = ""
	//response, err := s.Client.Query(ctx, &requirementQuery)
	//if err != nil {
	//	log.Error(err)
	//	return err
	//}
	//
	//rsp.Body = errcode.Return(response)
	return nil
}

func init() {
	logger, err := log.LoggerFromConfigAsFile("./config/con-log.xml")
	if err != nil{
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
