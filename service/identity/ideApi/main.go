package main

import (
	"encoding/json"
	log "github.com/jeanphorn/log4go"
	"github.com/bottos-project/bottos/service/identity/proto"
	"github.com/micro/go-micro"
	api "github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
	"github.com/asaskevich/govalidator"
	"strconv"
	"regexp"
	"github.com/bottos-project/bottos/config"
	"github.com/mojocn/base64Captcha"
)

var configCode = base64Captcha.ConfigDigit{
	Height:     80,
	Width:      240,
	MaxSkew:    0.7,
	DotCount:   80,
	CaptchaLen: 5,
}

//var configCode = base64Captcha.ConfigCharacter{
//	Height:             60,
//	Width:              240,
////	//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
//	Mode:               base64Captcha.CaptchaModeNumber,
//	ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
//	ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
//	IsShowHollowLine:   false,
//	IsShowNoiseDot:     false,
//	IsShowNoiseText:    false,
//	IsShowSlimeLine:    false,
//	IsShowSineLine:     false,
//	CaptchaLen:         6,
//}

type User struct {
	Client user.UserClient
}

func (u *User) GetVerificationCode(ctx context.Context, req *api.Request, rsp *api.Response) error {

	idKeyD, capD := base64Captcha.GenerateCaptcha("", configCode)
	//以base64编码
	base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code": 1,
		"data": map[string]interface{}{
			"id_key": idKeyD,
			"img_data": base64stringD,
		},
		"msg": "OK",
	})

	rsp.Body = string(b)
	return nil
}


func (u *User) Register(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	//transfer to struct
	var registerRequest user.RegisterRequest
	json.Unmarshal([]byte(body), &registerRequest)

	if !base64Captcha.VerifyCaptcha(registerRequest.IdKey, registerRequest.VerifyValue) {
		b, _ := json.Marshal(map[string]interface{}{
			"code": -8,
			"msg": "Verification code error",
		})
		rsp.StatusCode = 200
		rsp.Body = string(b)
		return nil
	}
	//Checkout data format
	ok, err := govalidator.ValidateStruct(registerRequest);
	if !ok {
		b, _ := json.Marshal(map[string]interface{}{
			"code": -7,
			"msg": err.Error(),
		})
		rsp.StatusCode = 200
		rsp.Body = string(b)
		return nil
	}
	match,_:=regexp.MatchString("^[1-5a-z.]{3,13}$",registerRequest.Username)
	log.Info(match)
	if !match {
		b, _ := json.Marshal(map[string]interface{}{
			"code": -9,
			"msg": "Username is illegal",
		})
		rsp.StatusCode = 200
		rsp.Body = string(b)
		return nil
	}

	response, err := u.Client.Register(ctx, &registerRequest)
	if err != nil {
		return err
	}
	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code": response.Code,
		"msg": response.Msg,
	})

	rsp.Body = string(b)
	return nil
}

func (s *User) Login(ctx context.Context, req *api.Request, rsp *api.Response) error {
	header, _ := json.Marshal(req.Header)
	response, err := s.Client.Login(ctx, &user.LoginRequest{
		Body: req.Body,
		Header:string(header),
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code": response.Code,
		"token":response.Token,
	})
	rsp.Body = string(b)

	return nil
}

func (s *User) Logout(ctx context.Context, req *api.Request, rsp *api.Response) error {
	token := req.Header["Token"]

	if token == nil {
		rsp.StatusCode = 200
		b, _ := json.Marshal(map[string]interface{}{
			"code": "4001",
			"msg":"Token is nil",
		})
		rsp.Body = string(b)
		return nil
	}
	response, err := s.Client.Logout(ctx, &user.LogoutRequest{
		Token: token.Values[0],
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code": response.Code,
		"msg":response.Msg,
	})
	rsp.Body = string(b)

	return nil
}

func (u *User) GetUserInfo(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	//transfer to struct
	var queryRequest user.GetUserInfoRequest
	json.Unmarshal([]byte(body), &queryRequest)
	//Checkout data format
	ok, err := govalidator.ValidateStruct(queryRequest);
	if !ok {
		b, _ := json.Marshal(map[string]interface{}{
			"code": -7,
			"msg": err.Error(),
		})
		rsp.StatusCode = 200
		rsp.Body = string(b)
		return nil
	}

	response, err := u.Client.GetUserInfo(ctx, &queryRequest)
	if err != nil {
		return err
	}

	b, _ := json.Marshal(map[string]interface{}{
		"code": response.Code,
		"msg": response.Msg,
		"data": response.Data,
	})
	rsp.StatusCode = 200
	rsp.Body = string(b)
	return nil
}

func (u *User) UpdateUserInfo(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	//transfer to struct
	var updateUserInfoRequest user.UpdateUserInfoRequest
	json.Unmarshal([]byte(body), &updateUserInfoRequest)
	//Checkout data format
	ok, err := govalidator.ValidateStruct(updateUserInfoRequest);
	if !ok {
		b, _ := json.Marshal(map[string]interface{}{
			"code": -7,
			"msg": err.Error(),
		})
		rsp.StatusCode = 200
		rsp.Body = string(b)
		return nil
	}

	response, err := u.Client.UpdateUserInfo(ctx, &updateUserInfoRequest)
	if err != nil {
		return err
	}
	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code": response.Code,
		"msg": response.Msg,
	})

	rsp.Body = string(b)
	return nil
}

func (u *User) FavoriteMng(ctx context.Context, req *api.Request, rsp *api.Response) error {
	//header, _ := json.Marshal(req.Header)
	response, err := u.Client.FavoriteMng(ctx, &user.FavoriteMngRequest{
		PostBody:   req.Body,
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code":  response.Code,
		"msg": response.Msg,
		"data": response.Data,
	})
	rsp.Body = string(b)

	return nil
}
func (u *User) QueryFavorite(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	//transfer to struct
	var queryRequest user.QueryFavoriteRequest
	json.Unmarshal([]byte(body), &queryRequest)
	//Checkout data format

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

	response, err := u.Client.QueryFavorite(ctx, &queryRequest)
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
func (u *User) AddNotice(ctx context.Context, req *api.Request, rsp *api.Response) error {
	response, err := u.Client.AddNotice(ctx, &user.AddNoticeRequest{
		PostBody:   req.Body,
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code":  response.Code,
		"msg": response.Msg,
		"data": response.Data,
	})
	rsp.Body = string(b)

	return nil
}
func (u *User) QueryNotice(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	//transfer to struct
	var queryRequest user.QueryNoticeRequest
	json.Unmarshal([]byte(body), &queryRequest)
	//Checkout data format
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

	response, err := u.Client.QueryNotice(ctx, &queryRequest)
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

func (u *User) GetAccount(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	//transfer to struct
	var queryRequest user.GetAccountRequest
	json.Unmarshal([]byte(body), &queryRequest)
	//Checkout data format
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

	response, err := u.Client.GetAccount(ctx, &queryRequest)
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
func (u *User) Transfer(ctx context.Context, req *api.Request, rsp *api.Response) error {
	//header, _ := json.Marshal(req.Header)
	response, err := u.Client.Transfer(ctx, &user.TransferRequest{
		PostBody:   req.Body,
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code":  response.Code,
		"msg": response.Msg,
		"data": response.Data,
	})
	rsp.Body = string(b)

	return nil
}
func (u *User) QueryTransfer(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	//transfer to struct
	var queryRequest user.QueryTransferRequest
	json.Unmarshal([]byte(body), &queryRequest)
	//Checkout data format
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

	response, err := u.Client.QueryTransfer(ctx, &queryRequest)
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

func (u *User) GetBlockInfo(ctx context.Context, req *api.Request, rsp *api.Response) error {
	response, err := u.Client.GetBlockInfo(ctx, &user.GetBlockInfoRequest{})
	if err != nil {
		return err
	}
	b, _ := json.Marshal(map[string]interface{}{
		"code": response.Code,
		"msg": response.Msg,
		"data": response.Data,
	})
	rsp.StatusCode = 200
	rsp.Body = string(b)
	return nil
}

func (u *User) GetDataBin(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body

	response, err := u.Client.GetDataBin(ctx, &user.GetDataBinRequest{
		Info:body,
	})
	if err != nil {
		return err
	}
	b, _ := json.Marshal(map[string]interface{}{
		"code": response.Code,
		"msg": response.Msg,
		"data": response.Data,
	})
	rsp.StatusCode = 200
	rsp.Body = string(b)
	return nil
}

func main() {
	log.LoadConfiguration(config.BASE_LOG_CONF)
	defer log.Close()
	log.LOGGER("user.api")

	service := micro.NewService(
		micro.Name("go.micro.api.v2.user"),
	)

	// parse command line flags
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&User{Client: user.NewUserClient("go.micro.srv.user", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Exit(err)
	}
}