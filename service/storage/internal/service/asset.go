/*Copyright 2017~2022 The Bottos Authors
  This file is part of the Bottos Data Exchange Client
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

package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/bottos-project/magiccube/service/storage/proto"
	"github.com/bottos-project/magiccube/service/storage/util"
)

// GetUserAssetList from server
func (c *StorageService) GetUserAssetList(ctx context.Context, request *storage.UserAssetListRequest, response *storage.UserAssetListResponse) error {

	if request == nil {
		response.Code = 0
		return errors.New("Missing storage request")
	}
	fmt.Println("GetUserAssetList")
	assets, err := c.mgoRepo.CallGetAssetListByUser(request.Username)
	if err != nil {
		response.Code = 0
		fmt.Println(err)
		return errors.New("Failed get put url")

	}

	response.UserAssetList = []*storage.UserAsset{}
	for _, user := range assets {
		dbTag := &storage.UserAsset{
			AssetId:     user.AssetID,
			AssetName:   user.AssetName,
			FeatureTag:  user.FeatureTag,
			SamplePath:  user.SamplePath,
			SampleHash:  user.SampleHash,
			StoragePath: user.StoragePath,
			StorageHash: user.StorageHash,
			ExpireTime:  user.ExpireTime,
			Price:       user.Price,
			Description: user.Description,
			UploadDate:  user.UploadDate}
		response.UserAssetList = append(response.UserAssetList, dbTag)
	}
	response.Code = 1
	return nil
}

// GetUserPurchaseAssetList from server
func (c *StorageService) GetUserPurchaseAssetList(ctx context.Context, request *storage.UserRequest, response *storage.UserAssetListResponse) error {

	if request == nil {
		response.Code = 0
		return errors.New("Missing storage request")
	}
	fmt.Println("GetUserAssetList")
	assets, err := c.mgoRepo.CallGetUserPurchaseAssetList(request.Username)
	if err != nil {
		response.Code = 0
		fmt.Println(err)
		return errors.New("Failed get put url")

	}

	response.UserAssetList = []*storage.UserAsset{}
	for _, user := range assets {
		dbTag := &storage.UserAsset{
			AssetId:     user.AssetID,
			AssetName:   user.AssetName,
			FeatureTag:  user.FeatureTag,
			SamplePath:  user.SamplePath,
			SampleHash:  user.SampleHash,
			StoragePath: user.StoragePath,
			StorageHash: user.StorageHash,
			ExpireTime:  user.ExpireTime,
			Price:       user.Price,
			Description: user.Description,
			UploadDate:  user.UploadDate}

		response.UserAssetList = append(response.UserAssetList, dbTag)
	}
	response.Code = 1
	return nil
}

// GetAllAssetList from server
func (c *StorageService) GetAllAssetList(ctx context.Context, request *storage.AssetListRequest, response *storage.AssetListResponse) error {

	if request == nil {
		response.Code = 0
		return errors.New("Missing storage request")
	}
	fmt.Println("GetAllAssetList")
	assets, err := c.mgoRepo.CallGetAllAssetList()
	if err != nil {
		response.Code = 0
		fmt.Println(err)
		return errors.New("Failed get put url")

	}
	response.AssetList = []*storage.Asset{}
	for _, asset := range assets {
		dbTag := &storage.Asset{
			AssetId:     asset.AssetID,
			Username:    asset.UserName,
			AssetName:   asset.AssetName,
			FeatureTag:  asset.FeatureTag,
			SamplePath:  asset.SamplePath,
			SampleHash:  asset.SampleHash,
			StoragePath: asset.StoragePath,
			StorageHash: asset.StorageHash,
			ExpireTime:  asset.ExpireTime,
			Price:       asset.Price,
			Description: asset.Description,
			UploadDate:  asset.UploadDate}
		response.AssetList = append(response.AssetList, dbTag)
	}
	response.Code = 1
	return nil
}

// GetAssetByAssetId from server
func (c *StorageService) GetAssetByAssetId(ctx context.Context, request *storage.AssetIdRequest, response *storage.AssetInfoResponse) error {
	if request == nil {
		response.Code = 0
		return errors.New("Missing storage request")
	}
	fmt.Println("GetAssetByAssetId")
	asset, err := c.mgoRepo.CallGetAssetById(request.AssetId)
	if err != nil {
		response.Code = 0
		fmt.Println(err)
		return errors.New("Failed get put url")

	}
	response.AssetInfo = &storage.Asset{
		AssetId:     asset.AssetID,
		Username:    asset.UserName,
		AssetName:   asset.AssetName,
		FeatureTag:  asset.FeatureTag,
		SamplePath:  asset.SamplePath,
		SampleHash:  asset.SampleHash,
		StoragePath: asset.StoragePath,
		StorageHash: asset.StorageHash,
		ExpireTime:  asset.ExpireTime,
		Price:       asset.Price,
		Description: asset.Description,
		UploadDate:  asset.UploadDate}
	response.Code = 1
	return nil
}

// GetAssetNumByDay from server
func (c *StorageService) GetAssetNumByDay(ctx context.Context, request *storage.AllRequest, response *storage.DayAssetNumResponse) error {

	if request == nil {
		response.Code = 0
		return errors.New("Missing storage request")
	}
	fmt.Println("GetAssetNumByDay")
	begin, end := util.YesterdayDur()
	assetNum, err := c.mgoRepo.CallGetAssetNumByDay(begin, end)
	if err != nil {
		response.Code = 0
		fmt.Println(err)
		return errors.New("Failed GetAssetNumByDay")

	}
	response.DayAssetNum = assetNum
	response.Code = 1
	return nil
}

// GetAssetNumByWeek from server
func (c *StorageService) GetAssetNumByWeek(ctx context.Context, request *storage.AllRequest, response *storage.WeekAssetNumResponse) error {

	if request == nil {
		response.Code = 0
		return errors.New("Missing storage request")
	}
	fmt.Println("GetAssetNumByDay")
	response.WeekAssetNum = make([]uint64, 1, 7)
	days := util.WeekDur()
	for _, day := range days {
		assetNum, err := c.mgoRepo.CallGetAssetNumByDay(day.Begin, day.End)
		if err != nil {
			response.Code = 0
			fmt.Println(err)
			return errors.New("Failed CallGetAssetNumByDay")
		}
		response.WeekAssetNum = append(response.WeekAssetNum, assetNum)
	}

	response.Code = 1
	return nil
}
