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
	"time"

	"github.com/bottos-project/magiccube/service/storage/proto"
	"github.com/bottos-project/magiccube/service/storage/util"
)

type StorageService struct {
	minioRepo minioRepository
	dbRepo    dbRepository
	mgoRepo   mgoRepository
}

type minioRepository interface {
	GetPutURL(username string, objectName string) (string, error)
	GetPutState(username string, objectName string) (int64, error)
	GetFileDownloadURL(username string, objectName string) (string, error)
}
type dbRepository interface {
	CallInsertUserInfo(value util.UserDBInfo) error
	CallGetUserInfo(value string) (*util.UserDBInfo, error)
	CallGetUserNum() (uint32, error)
	CallInsertTxInfo(value util.TxDBInfo) error
	CallGetTx(txid string) (*util.TxDBInfo, error)
	CallInsertUserToken(string, string) (uint32, error)
	CallGetUserToken(string, string) (*util.TokenDBInfo, error)
	CallDelUserToken(string, string) (uint32, error)
	CallGetSyncBlockCount() (uint64, error)
}
type mgoRepository interface {
	CallInsertUserToken(string, string) (uint32, error)
	CallGetUserToken(string, string) (*util.TokenDBInfo, error)
	CallDelUserToken(string, string) (uint32, error)
	CallGetUserRequirementList(string) ([]*util.RequirementDBInfo, error)
	CallGetRequirementListByFeature(uint64) ([]*util.RequirementDBInfo, error)
	CallGetRequirementNumByDay(time.Time, time.Time) (uint64, error)
	CallGetAllRequirementList() ([]*util.RequirementDBInfo, error)
	CallGetAssetListByUser(string) ([]*util.AssetDBInfo, error)
	CallGetAllAssetList() ([]*util.AssetDBInfo, error)
	CallGetAssetById(string) (*util.AssetDBInfo, error)
	CallGetUserPurchaseAssetList(string) ([]*util.AssetDBInfo, error)
	CallGetUserFileList(string) ([]*util.FileDBInfo, error)
	CallGetRecentTxList() ([]*util.TxDBInfo, error)
	CallGetUserTxList(string) ([]*util.TxDBInfo, error)
	CallGetRecentTransfersList() ([]*util.TransferDBInfo, error)
	CallGetAssetNumByDay(time.Time, time.Time) (uint64, error)
	CallGetNodeInfos() ([]*util.NodeDBInfo, error)
	CallGetSumTxAmount() (uint64, error)
	CallGetAllTxNum() (uint64, error)
	CallGetTxNumByDay(time.Time, time.Time) (uint64, error)
	CallGetFavoritListByUser(username string) ([]*util.FavoritDBInfo, error)
	CallGetDataPresaleByUser(username string) ([]*util.DataPresaleDBInfo, error)
}

func NewStorageService(minioRepo minioRepository, mgodb mgoRepository) storage.StorageHandler {
	return &StorageService{minioRepo: minioRepo, mgoRepo: mgodb}
}
