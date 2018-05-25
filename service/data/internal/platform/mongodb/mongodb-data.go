package mongodb

import (
	"errors"
	"fmt"

	"github.com/bottos-project/magiccube/service/data/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DataMessage struct {
	ID                 string        `bson:"_id"`
	BlockNumber          int           `bson:"block_number"`
	TransactionID      string        `bson:"transaction_id"`
	SequenceNum      int `bson:"sequence_num"`
	BlockHash string        `bson:"block_hash"`
	CursorNum                 int        `bson:"cursor_num"`
	CursorLabel          int           `bson:"cursor_label"`
	Lifetime      int        `bson:"lifetime"`
	Sender      string `bson:"sender"`
	Cntract string        `bson:"contract"`
	Method      string        `bson:"method"`
	Param               struct {
		Filehash      string `bson:"filehash"`
		Info struct {
			Username string     `bson:"username"`
			Filesize       string     `bson:"filesize"`
			Filename       string     `bson:"filename"`
			Filepolicy      string     `bson:"filepolicy"`
			Filenumber       uint64     `bson:"filenumber"`
			Simorass       string     `bson:"simorass"`
			Optype     string     `bson:"optype"`
			Storeaddr     string     `bson:"storeaddr"`

		} `bson:"info"`
	} `bson:"param"`
	SigAlg int `bson:"sig_alg"`
	Signature string `bson:"signature"`
	CreatedTime string `bson:"created_time"`
	Version int `bson:"version"`
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
		return c.Find(bson.M{"method": "datafilereg", "param.filehash": merkleroothash}).All(&mesgs)
	}
	session.SetCollection("pre_datafilereg", query)
	fmt.Println(mesgs)
	var reqs uint64 = 0
	if mesgs != nil {
		reqs = 1
	}
	fmt.Println(session)
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
		return c.Find(bson.M{"method": "datafilereg", "param.filehash": guid}).One(&mesgs)
	}
	session.SetCollection("pre_datafilereg", query)
	fmt.Println("mesgs")
	fmt.Println(mesgs)
	reqs := &util.DataDBInfo{
		mesgs.Param.Filehash,
		mesgs.Param.Info.Username,
		mesgs.Param.Info.Filesize,
		mesgs.Param.Info.Filename,
		mesgs.Param.Info.Filepolicy,
		mesgs.Param.Info.Filenumber,
		mesgs.Param.Info.Simorass,
		mesgs.Param.Info.Optype,
		mesgs.Param.Info.Storeaddr,
		}

	return reqs, err
}
