package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/bottos-project/magiccube/service/storage/proto"
)

func (c *StorageService) GetUserFavorit(ctx context.Context, request *storage.UserRequest, response *storage.UserFavoritResponse) error {

	if request == nil {
		response.Code = 0
		return errors.New("Missing storage request")
	}
	fmt.Println("GetUserFavorit")
	favors, err := c.mgoRepo.CallGetFavoritListByUser(request.Username)

	if err != nil {
		response.Code = 0
		fmt.Println(err)
		return errors.New("Failed GetUserFavorit")

	}
	response.FavoritList = []*storage.Favorit{}
	for _, favor := range favors {
		dbTag := &storage.Favorit{favor.UserName,
			favor.OpType,
			favor.GoodsType,
			favor.GoodsID}
		response.FavoritList = append(response.FavoritList, dbTag)
	}
	response.Code = 1
	return nil
}
