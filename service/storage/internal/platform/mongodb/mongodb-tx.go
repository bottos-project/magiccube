package mongodb

import (
	"errors"
	"fmt"

	"time"

	//	"github.com/code/bottos/service/storage/blockchain"
	"github.com/code/bottos/service/storage/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

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

func (r *MongoRepository) CallGetRecentTxList() ([]*util.TxDBInfo, error) {
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
			purMsgs[i].TransactionID,
			purMsgs[i].Data.BasicInfo.UserName,
			asset.UserName,
			asset.Price,
			asset.FeatureTag,
			purMsgs[i].CreatedAt.String(),
			purMsgs[i].BlockNum}
		tfxs = append(tfxs, dbtag)
	}

	return tfxs, nil
}
func (r *MongoRepository) CallGetUserTxList(username string) ([]*util.TxDBInfo, error) {
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

	var tfxs = []*util.TxDBInfo{}
	for i := 0; i < len(purMsgs); i++ {
		asset, err := r.CallGetAssetById(purMsgs[i].Data.BasicInfo.AssetID)
		fmt.Println(purMsgs[i].Data.BasicInfo.AssetID)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("failed CallGetAssetById " + purMsgs[i].Data.BasicInfo.AssetID)
		}
		dbtag := &util.TxDBInfo{
			purMsgs[i].TransactionID,
			purMsgs[i].Data.BasicInfo.UserName,
			asset.UserName,
			asset.Price,
			asset.FeatureTag,
			purMsgs[i].CreatedAt.String(),
			purMsgs[i].BlockNum}
		tfxs = append(tfxs, dbtag)
	}

	return tfxs, nil
}
func (r *MongoRepository) CallGetSumTxAmount() (uint64, error) {
	session, err := GetSession(r.mgoEndpoint)
	if err != nil {
		fmt.Println(err)
		return 0, errors.New("Get session faild" + r.mgoEndpoint)
	}
	//defer session.Close()
	fmt.Println(session)
	var purMsgs []PurchaseMesssage
	//建议优化，支持多表查询
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
			mesgs[i].TransactionID,
			mesgs[i].Data.From,
			mesgs[i].Data.To,
			mesgs[i].Data.Quantity,
			mesgs[i].CreatedAt.String(),
			mesgs[i].BlockNum}
		tfxs = append(tfxs, dbtag)
	}
	return tfxs, nil
}
