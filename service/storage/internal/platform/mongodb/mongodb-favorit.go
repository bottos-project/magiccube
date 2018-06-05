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

// FavoritMessage is definition of favorite msg
type FavoritMessage struct {
	ID                 bson.ObjectId `bson:"_id"`
	BlockNum           int           `bson:"block_num"`
	MessageID          int           `bson:"message_id"`
	TransactionID      string        `bson:"transaction_id"`
	Authorization      []interface{} `bson:"authorization"`
	HandlerAccountName string        `bson:"handler_account_name"`
	Type               string        `bson:"type"`
	Data               struct {
		UserName  string `bson:"user_name"`
		SessionID string `bson:"session_id"`
		OpType    string `bson:"op_type"`
		GoodsType string `bson:"goods_type"`
		GoodsID   string `bson:"goods_id"`
		Signature string `bson:"signature"`
	} `bson:"data"`
	CreatedAt string `bson:"createdAt"`
}

// CallGetFavoritListByUser is to get favorite list by user
func (r *MongoRepository) CallGetFavoritListByUser(username string) ([]*util.FavoritDBInfo, error) {
	session, err := GetSession(r.mgoEndpoint)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Get session faild" + r.mgoEndpoint)
	}
	//defer session.Close()
	fmt.Println(session)
	var mesgs []FavoritMessage
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"type": "favoritepro", "data.user_name": username}).All(&mesgs)
	}
	session.SetCollection("Messages", query)
	fmt.Println(mesgs)
	var favors = []*util.FavoritDBInfo{}
	for i := 0; i < len(mesgs); i++ {
		dbtag := &util.FavoritDBInfo{
			UserName  : mesgs[i].Data.UserName,
			OpType    : mesgs[i].Data.OpType,
			GoodsType : mesgs[i].Data.GoodsType,
			GoodsID   : mesgs[i].Data.GoodsID}
		favors = append(favors, dbtag)
	}

	fmt.Println(favors)
	return favors, nil
}
