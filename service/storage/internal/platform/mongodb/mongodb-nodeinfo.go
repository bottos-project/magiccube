package mongodb

import (
	"errors"
	"fmt"
	"time"

	"github.com/bottos-project/magiccube/service/storage/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

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
			mesgs[i].Data.BasicInfo.NodeIP,
			mesgs[i].Data.BasicInfo.NodePort,
			mesgs[i].Data.BasicInfo.NodeAddress}

		reqs = append(reqs, dbtag)
	}

	fmt.Println(reqs)
	return reqs, nil
}
