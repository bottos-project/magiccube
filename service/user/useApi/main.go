/*Copyright 2017~2022 The Bottos Authors
  This file is part of the Bottos Service Layer
  Created by Developers Team of Bottos.

  This program is free software: you can distribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with Bottos. If not, see <http://www.gnu.org/licenses/>.
*/
package main

import (
	"encoding/json"
	"github.com/bottos-project/magiccube/config"
	errcode "github.com/bottos-project/magiccube/error"
	sign "github.com/bottos-project/magiccube/service/common/signature"
	"github.com/bottos-project/magiccube/service/user/proto"
	log "github.com/cihub/seelog"
	"github.com/micro/go-micro"
	api "github.com/micro/micro/api/proto"
	"github.com/mojocn/base64Captcha"
	"golang.org/x/net/context"
	"os"
	"regexp"
)

//User struct
type User struct {
	Client user.UserClient
}

//GetVerify is to verify
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
			"verify_id":   idKeyD,
			"verify_data": base64stringD,
		},
		"msg": "OK",
	})

	rsp.Body = string(b)
	return nil
}

//GetBlockHeader is to get blockheader
func (u *User) GetBlockHeader(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.StatusCode = 200
	response, err := u.Client.GetBlockHeader(ctx, &user.GetBlockHeaderRequest{})
	if err != nil {

		return err
	}

	rsp.Body = errcode.Return(response)
	return nil
}

//Register is to Register
func (u *User) Register(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.StatusCode = 200
	log.Info(req.Body)
	var registerRequest user.RegisterRequest
	err := json.Unmarshal([]byte(req.Body), &registerRequest)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Info("registerRequest",registerRequest)

	match, err := regexp.MatchString("^[a-z][a-z1-9]{2,15}$", registerRequest.Account.Name)
	if err != nil {
		log.Error(err)
		return err
	}
	if !match {
		rsp.Body = errcode.ReturnError(1002)
		return nil
	}

	if config.EnableVerification {
		if !base64Captcha.VerifyCaptcha(registerRequest.VerifyId, registerRequest.VerifyValue) {
			rsp.Body = errcode.ReturnError(1001)
			return nil
		}
	}

	userJSONBuf, err := json.Marshal(registerRequest.User)
	if err != nil {
		log.Error(err)
		return err
	}

	isTrue, err := sign.PushVerifySign(string(userJSONBuf), registerRequest.Account.Pubkey)
	if !isTrue {
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

//GetAccountInfo is to get AccountInfo
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

//Login is to Login
func (u *User) Login(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.StatusCode = 200

	var loginRequest user.LoginRequest
	err := json.Unmarshal([]byte(req.Body), &loginRequest)
	if err != nil {
		log.Error(err)
		return err
	}

	if config.EnableVerification {
		if !base64Captcha.VerifyCaptcha(loginRequest.VerifyId, loginRequest.VerifyValue) {
			rsp.Body = errcode.ReturnError(1001)
			return nil
		}
	}

	is, err := sign.QueryVerifySign(req.Body)
	if !is {
		rsp.Body = errcode.ReturnError(1000, err)
		return nil
	}

	rsp.Body = errcode.ReturnError(1)
	return nil
}

//Favorite is to Favorite
func (u *User) Favorite(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.StatusCode = 200

	var favoriteRequest user.FavoriteRequest
	err := json.Unmarshal([]byte(req.Body), &favoriteRequest)
	if err != nil {
		log.Error(err)
		return err
	}

	is, err := sign.PushVerifySign(req.Body)
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

//GetFavorite is to GetFavorite
func (u *User) GetFavorite(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.StatusCode = 200

	var getFavoriteRequest user.GetFavoriteRequest
	err := json.Unmarshal([]byte(req.Body), &getFavoriteRequest)
	if err != nil {
		log.Error(err)
		return err
	}

	is, err := sign.QueryVerifySign(req.Body)
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

// Transfer is to Transfer
func (u *User) Transfer(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.StatusCode = 200

	var pushTxRequest user.PushTxRequest
	err := json.Unmarshal([]byte(req.Body), &pushTxRequest)
	if err != nil {
		log.Error(err)
		return err
	}

	is, err := sign.PushVerifySign(req.Body)
	if !is {
		rsp.Body = errcode.ReturnError(1000, err)
		return nil
	}

	response, err := u.Client.Transfer(ctx, &pushTxRequest)
	if err != nil {

		return err
	}

	rsp.Body = errcode.Return(response)
	return nil
}

// GetTransfer is to QueryMyTransfer List
func (u *User) GetTransfer(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.StatusCode = 200
	body := req.Body
	var queryMyRequest user.GetTransferRequest
	err := json.Unmarshal([]byte(body), &queryMyRequest)
	if err != nil {
		log.Error(err)
		return err
	}

	isTrue, err := sign.QueryVerifySign(req.Body)
	if !isTrue {
		rsp.Body = errcode.ReturnError(1000, err)
		return nil
	}

	response, err := u.Client.GetTransfer(ctx, &queryMyRequest)


	if err != nil {
		log.Error(err)
		return err
	}

	rsp.Body = errcode.Return(response)
	return nil
}

// QueryMyBuy is to QueryMyBuy
func (u *User) QueryMyBuy(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.StatusCode = 200
	body := req.Body
	var queryMyBuyRequest user.QueryMyBuyRequest
	err := json.Unmarshal([]byte(body), &queryMyBuyRequest)
	if err != nil {
		log.Error(err)
		return err
	}

	isTrue, err := sign.QueryVerifySign(req.Body)
	if !isTrue {
		rsp.Body = errcode.ReturnError(1000, err)
		return nil
	}

	response, err := u.Client.QueryMyBuy(ctx, &queryMyBuyRequest)
	if err != nil {
		log.Error(err)
		return err
	}

	rsp.Body = errcode.Return(response)
	return nil
}

//init is to init
func init() {
	logger, err := log.LoggerFromConfigAsFile("./config/user-log.xml")
	if err != nil {
		log.Error(err)
		panic(err)
	}
	defer logger.Flush()
	log.ReplaceLogger(logger)
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.v3.user"),
	)

	// parse command line flags
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&User{Client: user.NewUserClient("go.micro.srv.v3.user", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		os.Exit(1)
	}
}
