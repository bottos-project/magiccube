package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/code/bottos/service/storage/proto"

	"github.com/code/bottos/service/storage/util"
)

//func (c *StorageService) GetUserInfo(ctx context.Context, request *storage.UserInfoRequest, response *storage.UserInfoResponse) error {
//	if request == nil {
//		fmt.Println("request is nil ")
//		return errors.New("request is nil")
//	}
//	user, err := c.dbRepo.CallGetUserInfo(request.Username)

//	if err != nil {
//		fmt.Println("get CallGetUserInfo ")
//		return errors.New("get CallGetUserInfo failed")

//	}
//	response.Accountname = user.Accountname
//	response.OwnerPubKey = user.Ownerpubkey
//	response.ActivePubKey = user.Activepubkey
//	response.EncyptedInfo = user.EncyptedInfo
//	response.UserType = user.UserType
//	response.RoleType = user.RoleType
//	response.CompanyName = user.CompanyName
//	response.CompanyAddress = user.CompanyAddress
//	return nil
//}
//func (c *StorageService) GetUserNum(ctx context.Context, request *storage.UserNumRequest, response *storage.UserNumResponse) error {

//	if request == nil {
//		fmt.Println("request is nil ")
//		return errors.New("request is nil")
//	}
//	num, err := c.dbRepo.CallGetUserNum()

//	if err != nil {

//		fmt.Println("get CallGetUserNum ")
//		return errors.New("get CallGetUserNum failed")

//	}
//	response.Num = num
//	return nil
//}
func (c *StorageService) InsertUserToken(ctx context.Context, request *storage.InsertTokenRequest, response *storage.InsertTokenResponse) error {
	if request == nil {
		fmt.Println("request is nil ")
		return errors.New("request is nil")
	}
	result, err := c.mgoRepo.CallInsertUserToken(request.Username, request.Token)

	if err != nil {

		fmt.Println("get InsertUserToken ")
		return errors.New("get InsertUserToken failed")

	}
	response.Code = result
	return nil
}
func (c *StorageService) GetUserToken(ctx context.Context, request *storage.TokenRequest, response *storage.TokenResponse) error {
	response.Username = ""
	response.Token = ""
	response.InsertTime = 0
	response.Code = 0
	if request == nil {
		fmt.Println("request is nil ")

		return errors.New("request is nil")
	}
	var err error
	var token *util.TokenDBInfo
	token, err = c.mgoRepo.CallGetUserToken(request.Username, request.Token)

	if err != nil {

		fmt.Println("get GetUserToken ")
		return errors.New("get GetUserToken failed")

	}
	if token == nil {
		fmt.Println("get GetUserToken ")
		return errors.New("get GetUserToken null")
	}
	fmt.Println(token.Username)
	response.Username = token.Username
	response.Token = token.Token
	response.InsertTime = token.InsertTime
	response.Code = 1
	return nil
}
func (c *StorageService) DelUserToken(ctx context.Context, request *storage.DelTokenRequest, response *storage.DelTokenResponse) error {
	if request == nil {
		fmt.Println("request is nil ")
		return errors.New("request is nil")
	}
	code, err := c.mgoRepo.CallDelUserToken(request.Username, request.Token)

	if err != nil {

		fmt.Println("get DelUserToken ")
		return errors.New("get DelUserToken failed")

	}
	response.Code = code
	return nil
}

func (c *StorageService) AgeUserToken(ctx context.Context, request *storage.AgeTokenRequest, response *storage.AgeTokenResponse) error {
	if request == nil {
		fmt.Println("request is nil ")
		return errors.New("request is nil")
	}

	return nil
}
