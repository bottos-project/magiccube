package mongodb

import (
	"errors"
	"fmt"

	"github.com/bottos-project/bottos/service/data/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DataMessage struct {
	ID                 string        `bson:"_id"`
	MessageID          int           `bson:"message_id"`
	TransactionID      string        `bson:"transaction_id"`
	Authorization      []interface{} `bson:"authorization"`
	HandlerAccountName string        `bson:"handler_account_name"`
	Type               string        `bson:"type"`
	Data               struct {
		Guid      string `bson:"guid"`
		BasicInfo struct {
			MerkleRootHash string     `bson:"merkle_root_hash"`
			FileHash       string     `bson:"file_hash"`
			UserName       string     `bson:"user_name"`
			SessionID      string     `bson:"session_id"`
			FileSize       uint64     `bson:"file_size"`
			FileName       string     `bson:"file_name"`
			FilePolicy     string     `bson:"file_policy"`
			FileNumber     uint64     `bson:"file_number"`
			Signature      string     `bson:"signature"`
			Fslice         [][]string `bson:"fslice"`

			AuthPath string `bson:"auth_path"`
		} `bson:"basic_info"`
	} `bson:"data"`
	CreatedAt string `bson:"createdAt"`
}

func (r *MongoRepository) CallIsDataExists(merkleroothash string) (uint64, error) {
	session, err := GetSession(r.mgoEndpoint)
	if err != nil {
		fmt.Println(err)
		return 0, errors.New("Get session faild" + r.mgoEndpoint)
	}
	//defer session.Close()
	fmt.Println(session)
	var mesgs []DataMessage
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"type": "datafilereg", "data.basic_info.merkle_root_hash": merkleroothash}).All(&mesgs)
	}
	session.SetCollection("bottos", query)
	fmt.Println(mesgs)
	var reqs uint64 = 0
	if mesgs != nil {
		reqs = 1
	}
	return reqs, err
}
func (r *MongoRepository) CallDataSliceIPRequest(guid string) (*util.DataDBInfo, error) {
	session, err := GetSession(r.mgoEndpoint)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Get session faild" + r.mgoEndpoint)
	}
	//defer session.Close()
	fmt.Println(session)
	var mesgs DataMessage
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"type": "datafilereg", "data.guid": guid}).One(&mesgs)
	}

	session.SetCollection("bottos", query)
	fmt.Println("mesgs")
	fmt.Println(mesgs)
	reqs := &util.DataDBInfo{
		mesgs.Data.Guid,
		mesgs.Data.BasicInfo.MerkleRootHash,
		mesgs.Data.BasicInfo.UserName,
		mesgs.Data.BasicInfo.FileName,
		mesgs.Data.BasicInfo.FileSize,
		mesgs.Data.BasicInfo.FileNumber,
		mesgs.Data.BasicInfo.FilePolicy,
		mesgs.Data.BasicInfo.Fslice}

	return reqs, err
}
