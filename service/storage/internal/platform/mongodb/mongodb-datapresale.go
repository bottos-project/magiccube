package mongodb

import (
	"errors"
	"fmt"

	"github.com/bottos-project/magiccube/service/storage/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

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
			mesgs[i].Data.DataPresaleID,
			mesgs[i].Data.BasicInfo.UserName,
			mesgs[i].Data.BasicInfo.AssetID,
			mesgs[i].Data.BasicInfo.DataReqID,
			mesgs[i].Data.BasicInfo.Consumer}
		pres = append(pres, dbtag)
	}

	fmt.Println(pres)
	return pres, nil
}
