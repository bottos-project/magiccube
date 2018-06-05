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

package mongodb

import (
	"errors"
	"fmt"
	"time"

	"github.com/bottos-project/magiccube/service/storage/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// AssetMessage is definition of asset msg
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

// CallGetAssetListByUser is to get asset list by user
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
			AssetID:     mesgs[i].Data.AssetID,
			UserName:    mesgs[i].Data.BasicInfo.UserName,
			AssetName:   mesgs[i].Data.BasicInfo.AssetName,
			FeatureTag:  mesgs[i].Data.BasicInfo.FeatureTag,
			SamplePath:  mesgs[i].Data.BasicInfo.SamplePath,
			SampleHash:  mesgs[i].Data.BasicInfo.SampleHash,
			StoragePath: mesgs[i].Data.BasicInfo.StoragePath,
			StorageHash: mesgs[i].Data.BasicInfo.StorageHash,
			ExpireTime:  mesgs[i].Data.BasicInfo.ExpireTime,
			Price:       mesgs[i].Data.BasicInfo.Price,
			Description: mesgs[i].Data.BasicInfo.Description,
			UploadDate:  mesgs[i].Data.BasicInfo.UploadDate}
		assets = append(assets, dbtag)
	}

	fmt.Println(assets)
	return assets, nil
}

// CallGetUserPurchaseAssetList is to get user pruchase asset list
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

// CallGetAllAssetList is to get all asset list
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
			AssetID:     mesgs[i].Data.AssetID,
			UserName:    mesgs[i].Data.BasicInfo.UserName,
			AssetName:   mesgs[i].Data.BasicInfo.AssetName,
			FeatureTag:  mesgs[i].Data.BasicInfo.FeatureTag,
			SamplePath:  mesgs[i].Data.BasicInfo.SamplePath,
			SampleHash:  mesgs[i].Data.BasicInfo.SampleHash,
			StoragePath: mesgs[i].Data.BasicInfo.StoragePath,
			StorageHash: mesgs[i].Data.BasicInfo.StorageHash,
			ExpireTime:  mesgs[i].Data.BasicInfo.ExpireTime,
			Price:       mesgs[i].Data.BasicInfo.Price,
			Description: mesgs[i].Data.BasicInfo.Description,
			UploadDate:  mesgs[i].Data.BasicInfo.UploadDate}
		assets = append(assets, dbtag)
	}

	fmt.Println(assets)
	return assets, nil
}

// CallGetAssetNumByDay is to get asset num by day
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

// CallGetAssetById is to get asset by id
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
		AssetID:     mesgs.Data.AssetID,
		UserName:    mesgs.Data.BasicInfo.UserName,
		AssetName:   mesgs.Data.BasicInfo.AssetName,
		FeatureTag:  mesgs.Data.BasicInfo.FeatureTag,
		SamplePath:  mesgs.Data.BasicInfo.SamplePath,
		SampleHash:  mesgs.Data.BasicInfo.SampleHash,
		StoragePath: mesgs.Data.BasicInfo.StoragePath,
		StorageHash: mesgs.Data.BasicInfo.StorageHash,
		ExpireTime:  mesgs.Data.BasicInfo.ExpireTime,
		Price:       mesgs.Data.BasicInfo.Price,
		Description: mesgs.Data.BasicInfo.Description,
		UploadDate:  mesgs.Data.BasicInfo.UploadDate}

	fmt.Println(dbtag)
	return dbtag, nil
}
