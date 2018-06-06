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

// GetUserRequirementList from db
func (c *StorageService) GetUserRequirementList(ctx context.Context, request *storage.UserRequireListRequest, response *storage.UserRequireListResponse) error {

	if request == nil {
		response.Code = 0
		return errors.New("Missing storage request")
	}
	fmt.Println("GetUserRequirementList")
	requires, err := c.mgoRepo.CallGetUserRequirementList(request.Username)
	if err != nil {
		response.Code = 0
		fmt.Println(err)
		return errors.New("Failed get put url")

	}
	response.RequireList = []*storage.Requirement{}
	for _, require := range requires {
		requiresTag := &storage.Requirement{
			RequirementId:   require.RequirementId,
			Username:        require.Username,
			RequirementName: require.RequirementName,
			FeatureTag:      require.FeatureTag,
			SamplePath:      require.SamplePath,
			SampleHash:      require.SampleHash,
			ExpireTime:      require.ExpireTime,
			Price:           require.Price,
			Description:     require.Description,
			PublishDate:     require.PublishDate}
		response.RequireList = append(response.RequireList, requiresTag)
	}
	response.Code = 1
	return nil
}

// GetRequirementListByFeature from db
func (c *StorageService) GetRequirementListByFeature(ctx context.Context, request *storage.FeatureRequireListRequest, response *storage.FeatureRequireListResponse) error {

	if request == nil {
		response.Code = 0
		return errors.New("Missing storage request")
	}
	fmt.Println("GetRequirementListByFeature")
	requires, err := c.mgoRepo.CallGetRequirementListByFeature(request.FeatureTag)
	if err != nil {
		response.Code = 0
		fmt.Println(err)
		return errors.New("Failed get put url")

	}
	response.RequireList = []*storage.Requirement{}
	for _, require := range requires {
		requiresTag := &storage.Requirement{
			RequirementId:   require.RequirementId,
			Username:        require.Username,
			RequirementName: require.RequirementName,
			FeatureTag:      require.FeatureTag,
			SamplePath:      require.SamplePath,
			SampleHash:      require.SampleHash,
			ExpireTime:      require.ExpireTime,
			Price:           require.Price,
			Description:     require.Description,
			PublishDate:     require.PublishDate}
		response.RequireList = append(response.RequireList, requiresTag)
	}
	response.Code = 1
	return nil
}

// GetRequirementNumByDay from db
func (c *StorageService) GetRequirementNumByDay(ctx context.Context, request *storage.AllRequest, response *storage.DayRequirementNumResponse) error {

	response.DayRequirementNum = 200
	response.Code = 1
	return nil
}

// GetRequirementNumByWeek from db
func (c *StorageService) GetRequirementNumByWeek(ctx context.Context, request *storage.AllRequest, response *storage.WeekRequirementNumResponse) error {

	if request == nil {
		response.Code = 0
		return errors.New("Missing storage request")
	}
	fmt.Println("GetRequirementNumByWeek")
	response.WeekRequirementNum = make([]uint64, 1, 7)
	days := util.WeekDur()
	for _, day := range days {
		requireNum, err := c.mgoRepo.CallGetRequirementNumByDay(day.Begin, day.End)
		if err != nil {
			response.Code = 0
			fmt.Println(err)
			return errors.New("Failed CallGetAssetNumByDay")
		}
		response.WeekRequirementNum = append(response.WeekRequirementNum, requireNum)
	}
	response.Code = 1
	return nil
}

// GetAllRequirementList from db
func (c *StorageService) GetAllRequirementList(ctx context.Context, request *storage.AllRequest, response *storage.AllRequireListResponse) error {

	if request == nil {
		response.Code = 0
		return errors.New("Missing storage request")
	}
	fmt.Println("GetRequirementListByFeature")
	requires, err := c.mgoRepo.CallGetAllRequirementList()
	if err != nil {
		response.Code = 0
		fmt.Println(err)
		return errors.New("Failed get put url")

	}
	response.RequireList = []*storage.Requirement{}
	for _, require := range requires {
		dbTag := &storage.Requirement{
			RequirementId:   require.RequirementId,
			Username:        require.Username,
			RequirementName: require.RequirementName,
			FeatureTag:      require.FeatureTag,
			SamplePath:      require.SamplePath,
			SampleHash:      require.SampleHash,
			ExpireTime:      require.ExpireTime,
			Price:           require.Price,
			Description:     require.Description,
			PublishDate:     require.PublishDate}
		response.RequireList = append(response.RequireList, dbTag)
	}
	response.Code = 1
	return nil
}
