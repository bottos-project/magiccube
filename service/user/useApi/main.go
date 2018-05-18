package main

import (
	log "github.com/cihub/seelog"
	"encoding/json"
	"github.com/bottos-project/bottos/service/user/proto"
	"github.com/micro/go-micro"
	api "github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
	"github.com/mojocn/base64Captcha"
	"os"
	"github.com/bottos-project/bottos/config"
	"regexp"
	errcode "github.com/bottos-project/bottos/error"
	sign "github.com/bottos-project/bottos/service/common/signature"
)

type User struct {
	Client user.UserClient
}

func (u *User) GetVerify(ctx context.Context, req *api.Request, rsp *api.Response) error {
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
	idKeyD, capD := base64Captcha.GenerateCaptcha("", base64Captcha.ConfigDigit{
		Height:     80,
		Width:      240,
		MaxSkew:    0.7,
		DotCount:   80,
		CaptchaLen: 5,
	})
	base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code": 1,
		"data": map[string]interface{}{
			"verify_id": idKeyD,
			"verify_data": base64stringD,
		},
		"msg": "OK",
	})

	rsp.Body = string(b)
	return nil
}

func (u *User) GetBlockHeader(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.StatusCode = 200
	response, err := u.Client.GetBlockHeader(ctx, &user.GetBlockHeaderRequest{})
	if err != nil {

		return err
	}

	rsp.Body = errcode.Return(response)
	return nil
}

func (u *User) Register(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.StatusCode = 200
	log.Info(req.Body)
	var registerRequest user.RegisterRequest
	err := json.Unmarshal([]byte(req.Body), &registerRequest)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Info(registerRequest)

	if config.Enable_verification {
		if !base64Captcha.VerifyCaptcha(registerRequest.VerifyId, registerRequest.VerifyValue) {
			rsp.Body = errcode.ReturnError(1001)
			return nil
		}
	}

	match,err :=regexp.MatchString("^[a-km-z][a-km-z1-9]{2,15}$",registerRequest.Account.Name)
	if err != nil {
		log.Error(err)
		return err
	}
	if !match {
		rsp.Body = errcode.ReturnError(1002)
		return nil
	}

	user_json_buf, err := json.Marshal(registerRequest.User)
	if err != nil {
		log.Error(err)
		return err
	}

	is_true, err := sign.PushVerifySign(string(user_json_buf), registerRequest.Account.Pubkey)
	if !is_true {
		rsp.Body = errcode.ReturnError(1000, err)
		return nil
	}

	response, err := u.Client.Register(ctx, &registerRequest)
	if err != nil {
		return err
	}

	rsp.Body = errcode.Return(response)
	return nil
}

func (u *User) GetAccountInfo(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.StatusCode = 200
	var getAccountInfoRequest user.GetAccountInfoRequest
	err := json.Unmarshal([]byte(req.Body), &getAccountInfoRequest)
	if err != nil {
		log.Error(err)
		return err
	}
	response, err := u.Client.GetAccountInfo(ctx, &getAccountInfoRequest)
	if err != nil {

		return err
	}

	rsp.Body = errcode.Return(response)
	return nil
}

func (s *User) Login(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.StatusCode = 200

	var loginRequest user.LoginRequest
	err := json.Unmarshal([]byte(req.Body), &loginRequest)
	if err != nil {
		log.Error(err)
		return err
	}

	if config.Enable_verification {
		if !base64Captcha.VerifyCaptcha(loginRequest.VerifyId, loginRequest.VerifyValue) {
			rsp.Body = errcode.ReturnError(1001)
			return nil
		}
	}

	is, err:=sign.QueryVerifySign(req.Body)
	if !is {
		rsp.Body = errcode.ReturnError(1000, err)
		return nil
	}

	rsp.Body = errcode.ReturnError(1)
	return nil
}

func (u *User) Favorite(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.StatusCode = 200

	var favoriteRequest user.FavoriteRequest
	err := json.Unmarshal([]byte(req.Body), &favoriteRequest)
	if err != nil {
		log.Error(err)
		return err
	}

	is, err:=sign.QueryVerifySign(req.Body)
	if !is {
		rsp.Body = errcode.ReturnError(1000, err)
		return nil
	}

	response, err := u.Client.Favorite(ctx, &favoriteRequest)
	if err != nil {

		return err
	}

	rsp.Body = errcode.Return(response)
	return nil
}

func (u *User) GetFavorite(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.StatusCode = 200

	var getFavoriteRequest user.GetFavoriteRequest
	err := json.Unmarshal([]byte(req.Body), &getFavoriteRequest)
	if err != nil {
		log.Error(err)
		return err
	}

	is, err:=sign.QueryVerifySign(req.Body)
	if !is {
		rsp.Body = errcode.ReturnError(1000, err)
		return nil
	}

	response, err := u.Client.GetFavorite(ctx, &getFavoriteRequest)
	if err != nil {

		return err
	}

	rsp.Body = errcode.Return(response)
	return nil
}

func init() {
	logger, err := log.LoggerFromConfigAsFile("./config/user-log.xml")
	if err != nil{
		log.Error(err)
		panic(err)
	}
	defer logger.Flush()
	log.ReplaceLogger(logger)
}

func main() {
	service := micro.NewService(
		micro.Name("bottos.api.v3.user"),
	)

	// parse command line flags
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&User{Client: user.NewUserClient("bottos.srv.user", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		os.Exit(1)
	}
}