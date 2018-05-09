package mongodb

import (
	"errors"
	"fmt"

	"github.com/bottos-project/bottos/service/storage/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type FileMessage struct {
	ID                 string        `bson:"_id"`
	MessageID          int           `bson:"message_id"`
	TransactionID      string        `bson:"transaction_id"`
	Authorization      []interface{} `bson:"authorization"`
	HandlerAccountName string        `bson:"handler_account_name"`
	Type               string        `bson:"type"`
	Data               struct {
		FileHash  string `bson:"file_hash"`
		BasicInfo struct {
			UserName   string `bson:"user_name"`
			SessionID  string `bson:"session_id"`
			FileSize   uint64 `bson:"file_size"`
			FileName   string `bson:"file_name"`
			FilePolicy string `bson:"file_policy"`
			FileNumber uint64 `bson:"file_number"`
			Signature  string `bson:"signature"`
			AuthPath   string `bson:"auth_path"`
		} `bson:"basic_info"`
	} `bson:"data"`
	CreatedAt string `bson:"createdAt"`
}

func (r *MongoRepository) CallGetUserFileList(username string) ([]*util.FileDBInfo, error) {
	session, err := GetSession(r.mgoEndpoint)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Get session faild" + r.mgoEndpoint)
	}
	//defer session.Close()
	fmt.Println(session)
	var mesgs []FileMessage
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"type": "datafilereg", "data.basic_info.user_name": username}).All(&mesgs)
	}
	session.SetCollection("Messages", query)
	fmt.Println(mesgs)
	var reqs = []*util.FileDBInfo{}
	for i := 0; i < len(mesgs); i++ {
		dbtag := &util.FileDBInfo{
			mesgs[i].Data.FileHash,
			mesgs[i].Data.BasicInfo.UserName,
			mesgs[i].Data.BasicInfo.FileName,
			mesgs[i].Data.BasicInfo.FileSize,
			mesgs[i].Data.BasicInfo.FileNumber,
			mesgs[i].Data.BasicInfo.FilePolicy,
			mesgs[i].Data.BasicInfo.AuthPath}
		reqs = append(reqs, dbtag)
	}

	fmt.Println(reqs)
	return reqs, nil
}
