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

	log "github.com/cihub/seelog"

	"os"

	errcode "github.com/bottos-project/magiccube/error"
	"github.com/bottos-project/magiccube/service/asset/proto"
	sign "github.com/bottos-project/magiccube/service/common/signature"
	"github.com/micro/go-micro"
	api "github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
)

// Asset struct
type Asset struct {
	Client asset.AssetClient
}

/*func (s *Asset) GetFileUploadURL(ctx context.Context, req *api.Request, rsp *api.Response) error {
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
}*/

// RegisterFile on Chain
func (s *Asset) RegisterFile(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Info("RegisterFile Service Start")
	rsp.StatusCode = 200

	//verify signature
	isTrue, err := sign.PushVerifySign(req.Body)
	if !isTrue {

		rsp.Body = errcode.ReturnError(1000, err)
		return nil
	}

	var publishRequest asset.PushTxRequest
	err = json.Unmarshal([]byte(req.Body), &publishRequest)
	if err != nil {
		log.Error(err)
		return err
	}

	response, err := s.Client.RegisterFile(ctx, &publishRequest)
	if err != nil {
		return err
	}

	rsp.Body = errcode.Return(response)
	return nil
}

// QueryUploadedData from chain
func (s *Asset) QueryUploadedData(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Info("QueryUploadedData Service Start")
	rsp.StatusCode = 200
	body := req.Body
	log.Debug(body)
	var queryRequest asset.QueryRequest
	err := json.Unmarshal([]byte(body), &queryRequest)
	if err != nil {
		log.Error(err)
		return err
	}

	//verify signature
	isTrue, err := sign.QueryVerifySign(req.Body)

	if !isTrue {
		rsp.Body = errcode.ReturnError(1000, err)
		return nil
	}

	response, err := s.Client.QueryUploadedData(ctx, &queryRequest)
	if err != nil {
		log.Error(err)
		return err
	}

	rsp.Body = errcode.Return(response)
	return nil

}

//func (u *Asset) GetDownLoadURL(ctx context.Context, req *api.Request, rsp *api.Response) error {
//	body := req.Body
//	log.Info(body)
//	//transfer to struct
//	var downLoadRequest asset.GetDownLoadURLRequest
//	json.Unmarshal([]byte(body), &downLoadRequest)
//	//Checkout data format
//
//	log.Info(downLoadRequest)
//	ok, err := govalidator.ValidateStruct(downLoadRequest);
//	if !ok {
//		b, _ := json.Marshal(map[string]string{
//			"code": "-7",
//			"msg":  err.Error(),
//		})
//		rsp.StatusCode = 200
//		rsp.Body = string(b)
//		return nil
//	}
//
//	response, err := u.Client.GetDownLoadURL(ctx, &downLoadRequest)
//	if err != nil {
//		return err
//	}
//
//	b, _ := json.Marshal(map[string]string{
//		"code": strconv.Itoa(int(response.Code)),
//		"msg":  response.Msg,
//		"data": response.Data,
//	})
//	rsp.StatusCode = 200
//	rsp.Body = string(b)
//	return nil
//}

// RegisterAsset on Chain
func (s *Asset) RegisterAsset(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.StatusCode = 200

	//verify signature
	isTrue, err := sign.PushVerifySign(req.Body)

	if !isTrue {
		rsp.Body = errcode.ReturnError(1000, err)
		return nil
	}

	var publishRequest asset.PushTxRequest
	err = json.Unmarshal([]byte(req.Body), &publishRequest)
	if err != nil {
		log.Error(err)
		return err
	}

	response, err := s.Client.RegisterAsset(ctx, &publishRequest)
	if err != nil {
		return err
	}

	rsp.Body = errcode.Return(response)
	return nil
}

// QueryMyAsset from chain
func (s *Asset) QueryMyAsset(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Info("QueryMyAsset Service Start")
	rsp.StatusCode = 200
	body := req.Body
	var assetQuery asset.QueryRequest
	err := json.Unmarshal([]byte(body), &assetQuery)
	if err != nil {
		log.Error(err)
		return err
	}

	//verify signature
	isTrue, err := sign.QueryVerifySign(req.Body)

	if !isTrue {
		rsp.Body = errcode.ReturnError(1000, err)
		return nil
	}

	response, err := s.Client.QueryAsset(ctx, &assetQuery)
	if err != nil {
		log.Error(err)
		return err
	}

	rsp.Body = errcode.Return(response)
	return nil
}

// QueryAllAsset from Chain
func (s *Asset) QueryAllAsset(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Info("QueryAllAsset Service Start")
	rsp.StatusCode = 200
	body := req.Body
	log.Info(body)
	var assetQuery asset.QueryRequest
	err := json.Unmarshal([]byte(body), &assetQuery)
	if err != nil {
		log.Error(err)
		return err
	}
	assetQuery.Username = ""
	response, err := s.Client.QueryAsset(ctx, &assetQuery)
	if err != nil {
		log.Error(err)
		return err
	}

	rsp.Body = errcode.Return(response)
	return nil
}

/*func (u *Asset) Modify(ctx context.Context, req *api.Request, rsp *api.Response) error {
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
}*/

// QueryAssetByID from chain
func (s *Asset) QueryAssetByID(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Info("QueryAllAsset Service Start")
	rsp.StatusCode = 200
	body := req.Body
	log.Info(body)
	var assetQuery asset.QueryAssetByIDRequest
	err := json.Unmarshal([]byte(body), &assetQuery)
	if err != nil {
		log.Error(err)
		return err
	}
	response, err := s.Client.QueryAssetByID(ctx, &assetQuery)
	if err != nil {
		log.Error(err)
		return err
	}

	rsp.Body = errcode.Return(response)
	return nil
}

/*func (s *Asset) GetUserPurchaseAssetList(ctx context.Context, req *api.Request, rsp *api.Response) error {
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

	response, err := s.Client.GetUserPurchaseAssetList(ctx, &queryRequest)
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
}*/

// PreSaleNotice on Chain
func (s *Asset) PreSaleNotice(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Info("PreSaleNotice Service Start")
	rsp.StatusCode = 200

	//verify signature
	isTrue, err := sign.PushVerifySign(req.Body)

	if !isTrue {
		rsp.Body = errcode.ReturnError(1000, err)
		return nil
	}

	var publishRequest asset.PushTxRequest
	err = json.Unmarshal([]byte(req.Body), &publishRequest)
	if err != nil {
		log.Error(err)
		return err
	}

	response, err := s.Client.PreSaleNotice(ctx, &publishRequest)
	log.Error(response)
	if err != nil {
		return err
	}

	rsp.Body = errcode.Return(response)
	return nil
}

// QueryMyNotice from chain
func (s *Asset) QueryMyNotice(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Info("QueryMyNotice Service Start")
	rsp.StatusCode = 200
	body := req.Body
	var queryMyNotice asset.QueryMyNoticeRequest
	err := json.Unmarshal([]byte(body), &queryMyNotice)
	if err != nil {
		log.Error(err)
		return err
	}

	//verify signature
	isTrue, err := sign.QueryVerifySign(req.Body)

	if !isTrue {
		rsp.Body = errcode.ReturnError(1000, err)
		return nil
	}

	response, err := s.Client.QueryMyNotice(ctx, &queryMyNotice)
	if err != nil {
		log.Error(err)
		return err
	}

	rsp.Body = errcode.Return(response)
	return nil
}

// QueryMyPreSale from chain
func (s *Asset) QueryMyPreSale(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Info("QueryMyPreSale Service Start")
	rsp.StatusCode = 200
	body := req.Body
	var queryMyNotice asset.QueryMyNoticeRequest
	err := json.Unmarshal([]byte(body), &queryMyNotice)
	if err != nil {
		log.Error(err)
		return err
	}

	//verify signature
	isTrue, err := sign.QueryVerifySign(req.Body)

	if !isTrue {
		rsp.Body = errcode.ReturnError(1000, err)
		return nil
	}

	response, err := s.Client.QueryMyPreSale(ctx, &queryMyNotice)
	if err != nil {
		log.Error(err)
		return err
	}

	rsp.Body = errcode.Return(response)
	return nil
}

func init() {
	defer log.Flush()
	logger, err := log.LoggerFromConfigAsFile("./config/ass-log.xml")
	if err != nil {
		log.Critical("err parsing config log file", err)
		os.Exit(1)
		return
	}
	log.ReplaceLogger(logger)
}
func main() {
	log.Info("Asset API Service Start")

	service := micro.NewService(
		micro.Name("go.micro.api.v3.asset"),
	)

	// parse command line flags
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&Asset{Client: asset.NewAssetClient("go.micro.srv.v3.asset", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Critical("Asset API Service Run Failed", err)
	}
}
