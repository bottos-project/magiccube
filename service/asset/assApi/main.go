package main

import (
	"encoding/json"
	log "github.com/jeanphorn/log4go"

	"github.com/bottos-project/bottos/service/asset/proto"
	"github.com/micro/go-micro"
	api "github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
	"strconv"
	"github.com/asaskevich/govalidator"
	"github.com/bottos-project/bottos/config"
)

type Asset struct {
	Client asset.AssetClient
}

func (s *Asset) GetFileUploadURL(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Info("Start Get File URL!")
	//header, _ := json.Marshal(req.Header)
	response, err := s.Client.GetFileUploadURL(ctx, &asset.GetFileUploadURLRequest{
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

func (u *Asset) GetFileUploadStat(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	log.Info(body)
	//transfer to struct
	var queryRequest asset.GetFileUploadStatRequest
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

	response, err := u.Client.GetFileUploadStat(ctx, &queryRequest)
	if err != nil {
		return err
	}

	b, _ := json.Marshal(map[string]string{
		"code": strconv.Itoa(int(response.Code)),
		"msg":  response.Msg,
		"data": response.Data,
	})
	rsp.StatusCode = 200
	rsp.Body = string(b)
	return nil
}

func (s *Asset) RegisterFile(ctx context.Context, req *api.Request, rsp *api.Response) error {
	//header, _ := json.Marshal(req.Header)
	response, err := s.Client.RegisterFile(ctx, &asset.RegisterFileRequest{
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

func (u *Asset) QueryUploadedData(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	log.Info(body)
	//transfer to struct
	var queryRequest asset.QueryUploadedDataRequest
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

	response, err := u.Client.QueryUploadedData(ctx, &queryRequest)
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

func (u *Asset) GetDownLoadURL(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	log.Info(body)
	//transfer to struct
	var downLoadRequest asset.GetDownLoadURLRequest
	json.Unmarshal([]byte(body), &downLoadRequest)
	//Checkout data format

	log.Info(downLoadRequest)
	ok, err := govalidator.ValidateStruct(downLoadRequest);
	if !ok {
		b, _ := json.Marshal(map[string]string{
			"code": "-7",
			"msg":  err.Error(),
		})
		rsp.StatusCode = 200
		rsp.Body = string(b)
		return nil
	}

	response, err := u.Client.GetDownLoadURL(ctx, &downLoadRequest)
	if err != nil {
		return err
	}

	b, _ := json.Marshal(map[string]string{
		"code": strconv.Itoa(int(response.Code)),
		"msg":  response.Msg,
		"data": response.Data,
	})
	rsp.StatusCode = 200
	rsp.Body = string(b)
	return nil
}

func (u *Asset) Register(ctx context.Context, req *api.Request, rsp *api.Response) error {
	//header, _ := json.Marshal(req.Header)
	response, err := u.Client.Register(ctx, &asset.RegisterRequest{
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

func (u *Asset) Query(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	log.Info(body)
	//transfer to struct
	var queryRequest asset.QueryRequest
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

	response, err := u.Client.Query(ctx, &queryRequest)
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

func (u *Asset) QueryAllAsset(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	log.Info(body)
	//transfer to struct
	var queryRequest asset.QueryAllAssetRequest
	json.Unmarshal([]byte(body), &queryRequest)
	//Checkout data format

	log.Info(queryRequest)
	ok, err := govalidator.ValidateStruct(queryRequest)
	if !ok {
		b, _ := json.Marshal(map[string]string{
			"code": "-7",
			"msg":  err.Error(),
		})
		rsp.StatusCode = 200
		rsp.Body = string(b)
		return nil
	}

	response, err := u.Client.QueryAllAsset(ctx, &queryRequest)
	if err != nil {
		return err
	}

	b, _ := json.Marshal(map[string]string{
		"code": strconv.Itoa(int(response.Code)),
		"msg":  response.Msg,
		"data": response.Data,
	})
	rsp.StatusCode = 200
	rsp.Body = string(b)
	return nil
}

func (u *Asset) Modify(ctx context.Context, req *api.Request, rsp *api.Response) error {
	//header, _ := json.Marshal(req.Header)
	response, err := u.Client.Modify(ctx, &asset.ModifyRequest{
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

func (u *Asset) QueryByID(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	log.Info(body)
	//transfer to struct
	var queryRequest asset.QueryByIDRequest
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

	response, err := u.Client.QueryByID(ctx, &queryRequest)
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
func (u *Asset) GetUserPurchaseAssetList(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	log.Info(body)
	//transfer to struct
	var queryRequest asset.GetUserPurchaseAssetListRequest
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

	response, err := u.Client.GetUserPurchaseAssetList(ctx, &queryRequest)
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
	log.LOGGER("asset.api")

	service := micro.NewService(
		micro.Name("go.micro.api.v2.asset"),
	)

	// parse command line flags
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&Asset{Client: asset.NewAssetClient("go.micro.srv.v2.asset", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Exit(err)
	}
}
