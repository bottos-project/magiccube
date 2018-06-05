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

// NodeMessage is definition os node msg
type NodeMessage struct {
	ID                 bson.ObjectId `bson:"_id"`
	MessageID          int           `bson:"message_id"`
	TransactionID      string        `bson:"transaction_id"`
	Authorization      []interface{} `bson:"authorization"`
	HandlerAccountName string        `bson:"handler_account_name"`
	Type               string        `bson:"type"`
	Data               struct {
		NodeID    string `bson:"node_id"`
		BasicInfo struct {
			NodeIP      string `bson:"node_ip"`
			NodePort    string `bson:"node_port"`
			NodeAddress string `bson:"node_address"`
		} `bson:"basic_info"`
	} `bson:"data"`
	CreatedAt time.Time `bson:"createdAt"`
}

// CallGetNodeInfos is to get node info
func (r *MongoRepository) CallGetNodeInfos() ([]*util.NodeDBInfo, error) {
	session, err := GetSession(r.mgoEndpoint)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Get session faild" + r.mgoEndpoint)
	}
	//defer session.Close()
	fmt.Println(session)
	var mesgs []NodeMessage
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"type": "nodeinforeg"}).All(&mesgs)
	}
	session.SetCollection("Messages", query)
	fmt.Println(mesgs)
	var reqs = []*util.NodeDBInfo{}
	for i := 0; i < len(mesgs); i++ {

		dbtag := &util.NodeDBInfo{
			NodeId:   mesgs[i].Data.BasicInfo.NodeIP,
			NodeIP:   mesgs[i].Data.BasicInfo.NodePort,
			NodePort: mesgs[i].Data.BasicInfo.NodeAddress}

		reqs = append(reqs, dbtag)
	}

	fmt.Println(reqs)
	return reqs, nil
}
