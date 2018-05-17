package main

import (
	"encoding/json"
	log "github.com/jeanphorn/log4go"
	"github.com/bottos-project/bottos/service/requirement/proto"
	"github.com/micro/go-micro"
	api "github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
	"github.com/bottos-project/bottos/config"
	errcode "github.com/bottos-project/bottos/error"
	sign "github.com/bottos-project/bottos/service/common/signature"
)

type Requirement struct {
	Client requirement.RequirementClient
}

func (s *Requirement) Publish(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.StatusCode = 200

	//验签
	is_true, err := sign.PushVerifySign(req.Body)
	if !is_true {
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

func (s *Requirement) Query(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	log.Info(body)
	var requirementQuery requirement.QueryRequest
	err := json.Unmarshal([]byte(body), &requirementQuery)
	if err != nil {
		log.Error(err)
	}

	response, err := s.Client.Query(ctx, &requirementQuery)
	if err != nil {
		log.Error(err)
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code": response.Code,
		"data": response.Data,
		"msg":  response.Msg,
	})
	rsp.Body = string(b)
	return nil
}

func main() {
	log.LoadConfiguration(config.BASE_LOG_CONF)
	defer log.Close()
	log.LOGGER("requirement.api")

	service := micro.NewService(
		micro.Name("go.micro.api.v2.requirement"),
	)

	// parse command line flags
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&Requirement{Client: requirement.NewRequirementClient("go.micro.srv.requirement", service.Client())},
		),
	)
	if err := service.Run(); err != nil {
		log.Exit(err)
	}
}
