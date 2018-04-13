package mongodb

import (
	"errors"
	"fmt"

	"time"

	"github.com/code/bottos/service/storage/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type AssetMessage struct {
	ID                 bson.ObjectId `bson:"_id"`
	MessageID          int           `bson:"message_id"`
	TransactionID      string        `bson:"transaction_id"`
	Authorization      []interface{} `bson:"authorization"`
	HandlerAccountName string        `bson:"handler_account_name"`
	Type               string        `bson:"type"`
	Data               struct {
		AssetID   string `bson:"asset_id"`
		BasicInfo struct {
			UserName    string `bson:"user_name"`
			SessionID   string `bson:"session_id"`
			AssetName   string `bson:"asset_name"`
			FeatureTag  uint64 `bson:"feature_tag"`
			SamplePath  string `bson:"sample_path"`
			SampleHash  string `bson:"sample_hash"`
			StoragePath string `bson:"storage_path"`
			StorageHash string `bson:"storage_hash"`
			ExpireTime  uint32 `bson:"expire_time"`
			Price       uint64 `bson:"price"`
			Description string `bson:"description"`
			UploadDate  uint32 `bson:"upload_date"`
			Signature   string `bson:"signature"`
		} `bson:"basic_info"`
	} `bson:"data"`
	CreatedAt time.Time `bson:"createdAt"`
}

func (r *MongoRepository) CallGetAssetListByUser(username string) ([]*util.AssetDBInfo, error) {
	session, err := GetSession(r.mgoEndpoint)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Get session faild" + r.mgoEndpoint)
	}
	//defer session.Close()
	fmt.Println(session)
	var mesgs []AssetMessage
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"type": "assetreg", "data.basic_info.user_name": username}).All(&mesgs)
	}
	session.SetCollection("Messages", query)
	fmt.Println(mesgs)
	var assets = []*util.AssetDBInfo{}
	for i := 0; i < len(mesgs); i++ {
		dbtag := &util.AssetDBInfo{
			mesgs[i].Data.AssetID,
			mesgs[i].Data.BasicInfo.UserName,
			mesgs[i].Data.BasicInfo.AssetName,
			mesgs[i].Data.BasicInfo.FeatureTag,
			mesgs[i].Data.BasicInfo.SamplePath,
			mesgs[i].Data.BasicInfo.SampleHash,
			mesgs[i].Data.BasicInfo.StoragePath,
			mesgs[i].Data.BasicInfo.StorageHash,
			mesgs[i].Data.BasicInfo.ExpireTime,
			mesgs[i].Data.BasicInfo.Price,
			mesgs[i].Data.BasicInfo.Description,
			mesgs[i].Data.BasicInfo.UploadDate}
		assets = append(assets, dbtag)
	}

	fmt.Println(assets)
	return assets, nil
}
func (r *MongoRepository) CallGetUserPurchaseAssetList(username string) ([]*util.AssetDBInfo, error) {
	session, err := GetSession(r.mgoEndpoint)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Get session faild" + r.mgoEndpoint)
	}
	//defer session.Close()
	fmt.Println(session)
	var purMsgs []PurchaseMesssage
	//建议优化，支持多表查询
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"type": "datapurchase", "data.basic_info.user_name": username}).All(&purMsgs)
	}
	session.SetCollection("Messages", query)
	fmt.Println(purMsgs)

	var tfxs = []*util.AssetDBInfo{}
	for i := 0; i < len(purMsgs); i++ {
		asset, err := r.CallGetAssetById(purMsgs[i].Data.BasicInfo.AssetID)
		fmt.Println(purMsgs[i].Data.BasicInfo.AssetID)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("failed CallGetAssetById " + purMsgs[i].Data.BasicInfo.AssetID)
		}
		tfxs = append(tfxs, asset)
	}

	return tfxs, nil
}
func (r *MongoRepository) CallGetAllAssetList() ([]*util.AssetDBInfo, error) {
	session, err := GetSession(r.mgoEndpoint)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Get session faild" + r.mgoEndpoint)
	}
	//defer session.Close()
	fmt.Println(session)
	var mesgs []AssetMessage
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"type": "assetreg"}).All(&mesgs)
	}
	session.SetCollection("Messages", query)
	fmt.Println(mesgs)
	var assets = []*util.AssetDBInfo{}
	for i := 0; i < len(mesgs); i++ {
		dbtag := &util.AssetDBInfo{
			mesgs[i].Data.AssetID,
			mesgs[i].Data.BasicInfo.UserName,
			mesgs[i].Data.BasicInfo.AssetName,
			mesgs[i].Data.BasicInfo.FeatureTag,
			mesgs[i].Data.BasicInfo.SamplePath,
			mesgs[i].Data.BasicInfo.SampleHash,
			mesgs[i].Data.BasicInfo.StoragePath,
			mesgs[i].Data.BasicInfo.StorageHash,
			mesgs[i].Data.BasicInfo.ExpireTime,
			mesgs[i].Data.BasicInfo.Price,
			mesgs[i].Data.BasicInfo.Description,
			mesgs[i].Data.BasicInfo.UploadDate}
		assets = append(assets, dbtag)
	}

	fmt.Println(assets)
	return assets, nil
}
func (r *MongoRepository) CallGetAssetNumByDay(begin time.Time, end time.Time) (uint64, error) {
	session, err := GetSession(r.mgoEndpoint)
	if err != nil {
		fmt.Println(err)
		return 0, errors.New("Get session faild" + r.mgoEndpoint)
	}

	fmt.Println(begin)
	fmt.Println(end)
	//var mesgs []AssetMessage
	query := func(c *mgo.Collection) (int, error) {
		return c.Find(bson.M{"type": "assetreg",
			"createdAt": bson.M{"$gt": begin, "$lt": end}}).Count()
	}
	assetNum, err := session.SetCollectionCount("Messages", query)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return uint64(assetNum), nil

}

func (r *MongoRepository) CallGetAssetById(assertId string) (*util.AssetDBInfo, error) {

	session, err := GetSession(r.mgoEndpoint)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Get session faild" + r.mgoEndpoint)
	}
	//defer session.Close()
	fmt.Println(session)
	var mesgs AssetMessage
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"type": "assetreg", "data.asset_id": assertId}).One(&mesgs)
	}
	session.SetCollection("Messages", query)

	fmt.Println(mesgs)
	dbtag := &util.AssetDBInfo{
		mesgs.Data.AssetID,
		mesgs.Data.BasicInfo.UserName,
		mesgs.Data.BasicInfo.AssetName,
		mesgs.Data.BasicInfo.FeatureTag,
		mesgs.Data.BasicInfo.SamplePath,
		mesgs.Data.BasicInfo.SampleHash,
		mesgs.Data.BasicInfo.StoragePath,
		mesgs.Data.BasicInfo.StorageHash,
		mesgs.Data.BasicInfo.ExpireTime,
		mesgs.Data.BasicInfo.Price,
		mesgs.Data.BasicInfo.Description,
		mesgs.Data.BasicInfo.UploadDate}

	fmt.Println(dbtag)
	return dbtag, nil
}
