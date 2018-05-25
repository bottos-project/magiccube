package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/bottos-project/magiccube/service/storage/proto"
)

func (c *StorageService) GetUserDataPresale(ctx context.Context, request *storage.UserRequest, response *storage.UserDataPresaleResponse) error {

	if request == nil {
		response.Code = 0
		return errors.New("Missing storage request")
	}
	fmt.Println("GetUserFavorit")
	presales, err := c.mgoRepo.CallGetDataPresaleByUser(request.Username)

	if err != nil {
		response.Code = 0
		fmt.Println(err)
		return errors.New("Failed GetUserFavorit")

	}
	response.DataPresaleList = []*storage.DataPresale{}
	for _, presale := range presales {
		dbTag := &storage.DataPresale{
			presale.DataPresaleID,
			presale.UserName,
			presale.AssetID,
			presale.DataReqID,
			presale.Consumer}
		response.DataPresaleList = append(response.DataPresaleList, dbTag)
	}
	response.Code = 1
	return nil
}
