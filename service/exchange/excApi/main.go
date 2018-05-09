package main

import (
	"encoding/json"
	log "github.com/jeanphorn/log4go"
	"github.com/bottos-project/bottos/service/exchange/proto"
	"github.com/micro/go-micro"
	api "github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
	"github.com/asaskevich/govalidator"
	"strconv"
	"github.com/bottos-project/bottos/config"
)


type Exchange struct {
	Client exchange.ExchangeClient
}

func (u *Exchange) ConsumerBuy(ctx context.Context, req *api.Request, rsp *api.Response) error {
	//header, _ := json.Marshal(req.Header)
	response, err := u.Client.ConsumerBuy(ctx, &exchange.ConsumerBuyRequest{
		PostBody: req.Body,
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	rsp.Body = string(b)

	return nil
}

func (u *Exchange) QueryTx(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	log.Info(body)
	//transfer to struct
	var queryRequest exchange.QueryTxRequest
	json.Unmarshal([]byte(body), &queryRequest)
	//Checkout data format

	log.Info(queryRequest)
	ok, err := govalidator.ValidateStruct(queryRequest);
	if !ok {
		b, _ := json.Marshal(map[string]string{
			"code": "-7",
			"msg":  err.Error(),
		})
		rsp.StatusCode = 200
		rsp.Body = string(b)
		return nil
	}

	response, err := u.Client.QueryTx(ctx, &queryRequest)
	if err != nil {
		return err
	}

	b, _ := json.Marshal(map[string]interface{}{
		"code": strconv.Itoa(int(response.Code)),
		"msg":  response.Msg,
		"data": response.Data,
	})
	rsp.StatusCode = 200
	rsp.Body = string(b)
	return nil
}

func main() {
	log.LoadConfiguration(config.BASE_LOG_CONF)
	defer log.Close()
	log.LOGGER("exchange.api")

	service := micro.NewService(
		micro.Name("go.micro.api.v2.exchange"),
	)

	// parse command line flags
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&Exchange{Client: exchange.NewExchangeClient("go.micro.srv.v2.exchange", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Exit(err)
	}

}
