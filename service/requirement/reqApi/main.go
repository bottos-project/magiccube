package main

import (
	"encoding/json"
	log "github.com/jeanphorn/log4go"
	"github.com/bottos-project/bottos/service/requirement/proto"
	"github.com/micro/go-micro"
	api "github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
	"github.com/bottos-project/bottos/config"
)

type Requirement struct {
	Client requirement.RequirementClient
}

func (s *Requirement) Publish(ctx context.Context, req *api.Request, rsp *api.Response) error {
	response, err := s.Client.Publish(ctx, &requirement.PublishRequest{
		Body: req.Body,
	})
	if err != nil {
		return err
	}

	ret, _ := json.Marshal(map[string]interface{}{
		"code": response.Code,
		"data": response.Data,
		"msg":  response.Msg,
	})
	rsp.StatusCode = 200
	rsp.Body = string(ret)
	return nil
}

func (s *Requirement) Query(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	log.Info(body)
	var requirementQuery requirement.QueryRequest
	err := json.Unmarshal([]byte(body), &requirementQuery)

	log.Info(requirementQuery)
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
