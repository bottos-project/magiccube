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

	//	"github.com/bottos-project/magiccube/service/storage/blockchain"
	"github.com/bottos-project/magiccube/service/storage/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//TxMessage is transaction message
type TxMessage struct {
	ID             bson.ObjectId `bson:"_id"`
	TransactionID  string        `bson:"transaction_id"`
	SequenceNum    int           `bson:"sequence_num"`
	BlockID        string        `bson:"block_id"`
	RefBlockNum    uint64        `bson:"ref_block_num"`
	RefBlockPrefix string        `bson:"ref_block_prefix"`
	Scope          []interface{} `bson:"scope"`
	ReadScope      []interface{} `bson:"read_scope"`
	Expiration     string        `bson:"expiration"`
	Signatures     []interface{} `bson:"signatures"`
	Messages       []string      `bson:"messages"`
	CreatedAt      time.Time     `bson:"createdAt"`
}

//TransferMessage is transfer message
type TransferMessage struct {
	ID            bson.ObjectId `bson:"_id"`
	MessageID     int           `bson:"message_id"`
	BlockNum      uint64        `bson:"block_num"`
	TransactionID string        `bson:"transaction_id"`
	Authorization []struct {
		Account    string `bson:"account"`
		Permission string `bson:"permission"`
	} `bson:"authorization"`
	HandlerAccountName string `bson:"handler_account_name"`
	Type               string `bson:"type"`
	Data               struct {
		From   string `bson:"from"`
		To     string `bson:"to"`
		Amount string `bson:"amount"`
		Memo   string `bson:"memo"`
	} `bson:"data"`
	CreatedAt time.Time `bson:"createdAt"`
}

//TransferMes is transfer with quantity message
type TransferMes struct {
	ID                 bson.ObjectId `bson:"_id"`
	BlockNum           uint64        `bson:"block_num"`
	MessageID          int           `bson:"message_id"`
	TransactionID      string        `bson:"transaction_id"`
	Authorization      []interface{} `bson:"authorization"`
	HandlerAccountName string        `bson:"handler_account_name"`
	Type               string        `bson:"type"`
	Data               struct {
		From     string `bson:"from"`
		To       string `bson:"to"`
		Quantity uint64 `bson:"quantity"`
	} `bson:"data"`
	CreatedAt time.Time `bson:"createdAt"`
}

//PurchaseMesssage is purchase message
type PurchaseMesssage struct {
	ID                 bson.ObjectId `bson:"_id"`
	BlockNum           uint64        `bson:"block_num"`
	MessageID          int           `bson:"message_id"`
	TransactionID      string        `bson:"transaction_id"`
	Authorization      []interface{} `bson:"authorization"`
	HandlerAccountName string        `bson:"handler_account_name"`
	Type               string        `bson:"type"`
	Data               struct {
		DataDealID string `bson:"data_deal_id"`
		BasicInfo  struct {
			UserName  string `bson:"user_name"`
			SessionID string `bson:"session_id"`
			AssetID   string `bson:"asset_id"`
			RandomNum int    `bson:"random_num"`
			Signature string `bson:"signature"`
		} `bson:"basic_info"`
	} `bson:"data"`
	CreatedAt time.Time `bson:"createdAt"`
}

//CallGetRecentTxList is getting recent transaction lists
func (r *MongoRepository) CallGetRecentTxList() ([]*util.TxDBInfo, error) {
	session, err := GetSession(r.mgoEndpoint)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Get session faild" + r.mgoEndpoint)
	}
	//defer session.Close()
	fmt.Println(session)
	var purMsgs []PurchaseMesssage
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"type": "datapurchase"}).Sort("-createdAt").Limit(15).All(&purMsgs)
	}
	session.SetCollection("Messages", query)
	fmt.Println(purMsgs)

	var tfxs = []*util.TxDBInfo{}
	for i := 0; i < len(purMsgs); i++ {
		asset, err := r.CallGetAssetById(purMsgs[i].Data.BasicInfo.AssetID)
		fmt.Println(purMsgs[i].Data.BasicInfo.AssetID)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("failed CallGetAssetById " + purMsgs[i].Data.BasicInfo.AssetID)
		}
		dbtag := &util.TxDBInfo{
			TransactionID: purMsgs[i].TransactionID,
			From:          purMsgs[i].Data.BasicInfo.UserName,
			To:            asset.UserName,
			Price:         asset.Price,
			Type:          asset.FeatureTag,
			Date:          purMsgs[i].CreatedAt.String(),
			BlockId:       purMsgs[i].BlockNum}
		tfxs = append(tfxs, dbtag)
	}

	return tfxs, nil
}

//CallGetUserTxList is getting User transaction list
func (r *MongoRepository) CallGetUserTxList(username string) ([]*util.TxDBInfo, error) {
	session, err := GetSession(r.mgoEndpoint)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Get session faild" + r.mgoEndpoint)
	}
	//defer session.Close()
	fmt.Println(session)
	var purMsgs []PurchaseMesssage
	
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"type": "datapurchase", "data.basic_info.user_name": username}).All(&purMsgs)
	}
	session.SetCollection("Messages", query)
	fmt.Println(purMsgs)

	var tfxs = []*util.TxDBInfo{}
	for i := 0; i < len(purMsgs); i++ {
		asset, err := r.CallGetAssetById(purMsgs[i].Data.BasicInfo.AssetID)
		fmt.Println(purMsgs[i].Data.BasicInfo.AssetID)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("failed CallGetAssetById " + purMsgs[i].Data.BasicInfo.AssetID)
		}
		dbtag := &util.TxDBInfo{
			TransactionID: purMsgs[i].TransactionID,
			From:          purMsgs[i].Data.BasicInfo.UserName,
			To:            asset.UserName,
			Price:         asset.Price,
			Type:          asset.FeatureTag,
			Date:          purMsgs[i].CreatedAt.String(),
			BlockId:       purMsgs[i].BlockNum}
		tfxs = append(tfxs, dbtag)
	}

	return tfxs, nil
}

//CallGetSumTxAmount is getting sum transaction amount
func (r *MongoRepository) CallGetSumTxAmount() (uint64, error) {
	session, err := GetSession(r.mgoEndpoint)
	if err != nil {
		fmt.Println(err)
		return 0, errors.New("Get session faild" + r.mgoEndpoint)
	}
	//defer session.Close()
	fmt.Println(session)
	var purMsgs []PurchaseMesssage

	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"type": "datapurchase"}).All(&purMsgs)
	}
	session.SetCollection("Messages", query)
	fmt.Println(purMsgs)

	var sum uint64
	for i := 0; i < len(purMsgs); i++ {
		asset, err := r.CallGetAssetById(purMsgs[i].Data.BasicInfo.AssetID)
		fmt.Println(purMsgs[i].Data.BasicInfo.AssetID)
		if err != nil {
			fmt.Println(err)
			return 0, errors.New("failed CallGetAssetById " + purMsgs[i].Data.BasicInfo.AssetID)
		}
		sum += asset.Price
	}

	return sum, nil
}

//CallGetAllTxNum is getting all transaction number
func (r *MongoRepository) CallGetAllTxNum() (uint64, error) {
	session, err := GetSession(r.mgoEndpoint)
	var num uint64
	if err != nil {
		fmt.Println(err)
		return 0, errors.New("Get session faild" + r.mgoEndpoint)
	}
	query := func(c *mgo.Collection) (int, error) {
		return c.Find(nil).Count()
	}
	txnum, err := session.SetCollectionCount("Transactions", query)
	if err != nil {
		fmt.Println(err)
		return 0, errors.New("Transactions c.Find(nil).Count()")
	}
	fmt.Println(txnum)
	num = uint64(txnum)
	return num, nil
}

//CallGetTxNumByDay is getting transaction number by day
func (r *MongoRepository) CallGetTxNumByDay(begin time.Time, end time.Time) (uint64, error) {
	session, err := GetSession(r.mgoEndpoint)
	if err != nil {
		fmt.Println(err)
		return 0, errors.New("Get session faild" + r.mgoEndpoint)
	}

	fmt.Println(begin)
	fmt.Println(end)
	query := func(c *mgo.Collection) (int, error) {
		tfNum, err := c.Find(bson.M{"type": "transfer",
			"createdAt": bson.M{"$gt": begin, "$lt": end}}).Count()
		fmt.Println(tfNum)
		fmt.Println(err)
		purchase, err := c.Find(bson.M{"type": "datapurchase", "createdAt": bson.M{"$gt": begin, "$lt": end}}).Count()
		return tfNum + purchase, err
	}
	daytxnum, err2 := session.SetCollectionCount("Messages", query)
	if err2 != nil {
		fmt.Println(err2)
		return 0, errors.New("Transactions c.Find(nil).Count()")
	}
	fmt.Println("daytxnum", daytxnum)

	return uint64(daytxnum), nil
}

//CallGetRecentTransfersList is getting recent transfers list
func (r *MongoRepository) CallGetRecentTransfersList() ([]*util.TransferDBInfo, error) {
	session, err := GetSession(r.mgoEndpoint)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Get session faild" + r.mgoEndpoint)
	}
	//defer session.Close()
	fmt.Println(session)
	var mesgs []TransferMes
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"type": "transfer"}).Sort("-createdAt").Limit(15).All(&mesgs)
	}
	session.SetCollection("Messages", query)
	fmt.Println(mesgs)

	var tfxs = []*util.TransferDBInfo{}
	for i := 0; i < len(mesgs); i++ {
		dbtag := &util.TransferDBInfo{
			TransactionID: mesgs[i].TransactionID,
			From:          mesgs[i].Data.From,
			To:            mesgs[i].Data.To,
			Price:         mesgs[i].Data.Quantity,
			TxTime:        mesgs[i].CreatedAt.String(),
			BlockNum:      mesgs[i].BlockNum}
		tfxs = append(tfxs, dbtag)
	}
	return tfxs, nil
}
