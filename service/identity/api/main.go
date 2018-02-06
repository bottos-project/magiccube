package main

import (
	"encoding/json"
	"log"
	"strings"

	user "../proto"
	"github.com/micro/go-micro"
	api "github.com/micro/micro/api/proto"
	"github.com/micro/go-micro/errors"
	"golang.org/x/net/context"
)

type User struct {
	Client user.UserClient
}

func (s *User) Register(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received User.Register API request")

	name, ok := req.Get["userName"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "Name cannot be blank")
	}

	response, err := s.Client.Register(ctx, &user.RegisterRequest{
		Username: strings.Join(name.Values, " "),
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message": response.Msg,
	})
	rsp.Body = string(b)

	return nil
}

func (s *User) Login(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received User.Register API request")

	name, ok := req.Get["userName"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "Name cannot be blank")
	}

	response, err := s.Client.Login(ctx, &user.LoginRequest{
		Username: strings.Join(name.Values, " "),
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message": response.Msg,
	})
	rsp.Body = string(b)

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.user"),
	)

	// parse command line flags
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&User{Client: user.NewUserClient("go.micro.srv.user", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}