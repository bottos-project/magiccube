/*Copyright 2017~2022 The Bottos Authors
  This file is part of the Bottos Service Layer
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
	"github.com/bottos-project/magiccube/service/data/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	log "github.com/cihub/seelog"
)

type NodeMessage struct {
	ID                 bson.ObjectId `bson:"_id"`
	MessageID          int           `bson:"message_id"`
	TransactionID      string        `bson:"transaction_id"`
	Authorization      []interface{} `bson:"authorization"`
	HandlerAccountName string        `bson:"handler_account_name"`
	Type               string        `bson:"type"`
	Node               struct {
		NodeID    string `bson:"node_id"`
		BasicInfo struct {
			NodeIP      string   `bson:"node_ip"`
			NodePort    string   `bson:"node_port"`
			NodeAddress string   `bson:"node_address"`
			SeedIP      string   `bson:"seed_ip"`
			SlaveIP     []string `bson:"slave_ip"`
		} `bson:"basic_info"`
	} `bson:"node"`
	CreatedAt time.Time `bson:"createdAt"`
}

func (r *MongoRepository) CallNodeRequest(seedip string) (*util.NodeDBInfo, error) {
	log.Info("call node")
	session, err := GetSession(r.mgoEndpoint)
	if err != nil {
		log.Info(err)
		return nil, errors.New("Get session faild" + r.mgoEndpoint)
	}
	//defer session.Close()
	var mesgs NodeMessage
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"type": "nodeinforeg", "node.basic_info.seed_ip": seedip}).One(&mesgs)
	}

	session.SetCollection("pre_node", query)
	reqs := &util.NodeDBInfo{
		mesgs.Node.NodeID,
		mesgs.Node.BasicInfo.NodeIP,
		mesgs.Node.BasicInfo.NodePort,
		mesgs.Node.BasicInfo.NodeAddress,
		mesgs.Node.BasicInfo.SeedIP,
		mesgs.Node.BasicInfo.SlaveIP}

	return reqs, err
}
