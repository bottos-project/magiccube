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

	"github.com/bottos-project/magiccube/service/storage/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DataPresaleMessage is definition of data presale msg
type DataPresaleMessage struct {
	ID                 bson.ObjectId `bson:"_id"`
	BlockNum           int           `bson:"block_num"`
	MessageID          int           `bson:"message_id"`
	TransactionID      string        `bson:"transaction_id"`
	Authorization      []interface{} `bson:"authorization"`
	HandlerAccountName string        `bson:"handler_account_name"`
	Type               string        `bson:"type"`
	Data               struct {
		DataPresaleID string `bson:"data_presale_id"`
		BasicInfo     struct {
			UserName  string `bson:"user_name"`
			SessionID string `bson:"session_id"`
			AssetID   string `bson:"asset_id"`
			DataReqID string `bson:"data_req_id"`
			Consumer  string `bson:"consumer"`
			RandomNum int    `bson:"random_num"`
			Signature string `bson:"signature"`
		} `bson:"basic_info"`
	} `bson:"data"`
	CreatedAt string `bson:"createdAt"`
}

// CallGetDataPresaleByUser is to get data presale by user
func (r *MongoRepository) CallGetDataPresaleByUser(username string) ([]*util.DataPresaleDBInfo, error) {
	session, err := GetSession(r.mgoEndpoint)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Get session faild" + r.mgoEndpoint)
	}
	//defer session.Close()
	fmt.Println(session)
	var mesgs []DataPresaleMessage
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"type": "datapresale", "data.basic_info.user_name": username}).All(&mesgs)
	}
	session.SetCollection("Messages", query)
	fmt.Println(mesgs)
	var pres = []*util.DataPresaleDBInfo{}
	for i := 0; i < len(mesgs); i++ {
		dbtag := &util.DataPresaleDBInfo{
			DataPresaleID: mesgs[i].Data.DataPresaleID,
			UserName:      mesgs[i].Data.BasicInfo.UserName,
			AssetID:       mesgs[i].Data.BasicInfo.AssetID,
			DataReqID:     mesgs[i].Data.BasicInfo.DataReqID,
			Consumer:      mesgs[i].Data.BasicInfo.Consumer}
		pres = append(pres, dbtag)
	}

	fmt.Println(pres)
	return pres, nil
}
