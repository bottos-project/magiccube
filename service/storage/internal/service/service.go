package service

import (
	//"github.com/bottos-project/magiccube/service/storage/controller"
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

	//	CallAgeUserToken(timeout int64) (uint32, error)
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
