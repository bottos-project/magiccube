package mongodb

import (
	"errors"
	"fmt"
	"github.com/bottos-project/bottos/service/data/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
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
	session, err := GetSession(r.mgoEndpoint)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Get session faild" + r.mgoEndpoint)
	}
	//defer session.Close()
	fmt.Println(session)
	var mesgs NodeMessage
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"type": "nodeinforeg", "node.basic_info.seed_ip": seedip}).One(&mesgs)
	}

	session.SetCollection("pre_node", query)
	fmt.Println("mesgs")
	fmt.Println(mesgs)
	reqs := &util.NodeDBInfo{
		mesgs.Node.NodeID,
		mesgs.Node.BasicInfo.NodeIP,
		mesgs.Node.BasicInfo.NodePort,
		mesgs.Node.BasicInfo.NodeAddress,
		mesgs.Node.BasicInfo.SeedIP,
		mesgs.Node.BasicInfo.SlaveIP}

	return reqs, err
}
