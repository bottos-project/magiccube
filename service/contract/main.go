package main

import (
	log "github.com/cihub/seelog"
	"github.com/micro/go-micro"
	contract_proto "github.com/bottos-project/magiccube/service/contract/proto"
	"golang.org/x/net/context"
	"os"
)


type Contract struct {}

func (u *Contract) Deploy(ctx context.Context, req *contract_proto.DeployRequest, rsp *contract_proto.DeployResponse) error {
	return nil
}

func (u *Contract) Query(ctx context.Context, req *contract_proto.QueryRequest, rsp *contract_proto.QueryResponse) error {
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
		micro.Name("go.micro.srv.v3.contract"),
		micro.Version("3.0.0"),
	)

	service.Init()

	contract_proto.RegisterContractHandler(service.Server(), new(Contract))

	if err := service.Run(); err != nil {
		os.Exit(1)
	}
}
