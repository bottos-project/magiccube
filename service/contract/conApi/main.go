package main

import (
	log "github.com/cihub/seelog"

	"github.com/bottos-project/magiccube/service/contract/proto"
	"github.com/micro/go-micro"
	api "github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
	"os"
	"strings"
	"io/ioutil"
	"net/http"
	"github.com/bottos-project/magiccube/service/common/data"
	"errors"
	"encoding/json"
)

type Contract struct {
	Client contract.ContractClient
}

func (s *Contract) Publish(ctx context.Context, req *api.Request, rsp *api.Response) error {

	return nil
}

func (s *Contract) Query(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var queryRequest contract.QueryRequest
	err := json.Unmarshal([]byte(req.Body), &queryRequest)
	if err != nil {
		log.Error(err)
		return err
	}

	params := `service=bottos&method=CoreApi.QueryAbi&request={"contract":"`+queryRequest.Contract+`"}`
	resp, err := http.Post(data.BASE_URL, "application/x-www-form-urlencoded",
		strings.NewReader(params))
	if err != nil {
		log.Error(err)
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if (resp.StatusCode != 200) {
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
