package mongodb

import (
	"errors"
	"fmt"

	"github.com/bottos-project/magiccube/service/storage/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

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
			mesgs[i].Data.UserName,
			mesgs[i].Data.OpType,
			mesgs[i].Data.GoodsType,
			mesgs[i].Data.GoodsID}
		favors = append(favors, dbtag)
	}

	fmt.Println(favors)
	return favors, nil
}
