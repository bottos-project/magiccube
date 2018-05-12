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
	"github.com/protobuf/proto"
	"github.com/bottos-project/crypto-go/crypto"
	errcode "github.com/bottos-project/bottos/error"
	"crypto/sha256"
	"encoding/hex"
	"github.com/bottos-project/bottos/service/common/proto"
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
			"verify_key": idKeyD,
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

	match,err :=regexp.MatchString("^[1-5a-z.]{3,13}$",registerRequest.Account.Name)
	if err != nil {
		log.Error(err)
		return err
	}
	if !match {
		rsp.Body = errcode.ReturnError(1002)
		return nil
	}

	pubkey,err := hex.DecodeString(registerRequest.Account.Pubkey)
	if err != nil {
		log.Error(err)
		return err
	}

	signature,err := hex.DecodeString(registerRequest.User.Signature)
	if err != nil {
		log.Error(err)
		return err
	}
	registerRequest.User.Signature = ""
	serializeData, err := proto.Marshal(registerRequest.User)
	if err != nil {
		log.Error(err)
		return err
	}

	var d sign.Message;
	proto.Unmarshal(serializeData, &d)
	log.Info(d)
	log.Info(hex.EncodeToString(serializeData))

	h := sha256.New()
	h.Write([]byte(hex.EncodeToString(serializeData)))
	hash := h.Sum(nil)

	if !crypto.VerifySign(pubkey, hash, signature) {
		rsp.Body = errcode.ReturnError(1000)
		return nil
	}

	//response, err := u.Client.Register(ctx, &registerRequest)
	//if err != nil {
	//	return err
	//}

	b, _ := json.Marshal(map[string]interface{}{
		"code": 1,
		"msg": "ok",
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