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
	log "github.com/cihub/seelog"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DataMessage struct
type DataMessage struct {
	ID            string `bson:"_id"`
	BlockNumber   int    `bson:"block_number"`
	TransactionID string `bson:"transaction_id"`
	SequenceNum   int    `bson:"sequence_num"`
	BlockHash     string `bson:"block_hash"`
	CursorNum     int    `bson:"cursor_num"`
	CursorLabel   int    `bson:"cursor_label"`
	Lifetime      int    `bson:"lifetime"`
	Sender        string `bson:"sender"`
	Cntract       string `bson:"contract"`
	Method        string `bson:"method"`
	Param         struct {
		Filehash string `bson:"filehash"`
		Info     struct {
			Username   string `bson:"username"`
			Filename   string `bson:"filename"`
			Filesize   uint64 `bson:"filesize"`
			Filepolicy string `bson:"filepolicy"`
			Filenumber uint64 `bson:"filenumber"`
			Simorass   uint64 `bson:"simorass"`
			Optype     uint64 `bson:"optype"`
			Storeaddr  string `bson:"storeaddr"`
		} `bson:"info"`
	} `bson:"param"`
	SigAlg      int    `bson:"sig_alg"`
	Signature   string `bson:"signature"`
	CreatedTime string `bson:"created_time"`
	Version     int    `bson:"version"`
}

// CallIsDataExists check
func (r *MongoRepository) CallIsDataExists(merkleroothash string) (uint64, error) {
	log.Info("call is data exists")
	session, err := GetSession(r.mgoEndpoint)
	if err != nil {
		log.Info("err")
		return 0, errors.New("Get session faild" + r.mgoEndpoint)
	}
	//defer session.Close()
	var mesgs []DataMessage
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"method": "datafilereg", "param.filehash": merkleroothash}).All(&mesgs)
	}
	session.SetCollection("pre_datafilereg", query)
	var reqs uint64
	if mesgs != nil {
		reqs = 1
	}
	return reqs, err
}

// CallDataSliceIPRequest on server
func (r *MongoRepository) CallDataSliceIPRequest(guid string) (*util.DataDBInfo, error) {
	log.Info("call datasliceip")
	session, err := GetSession(r.mgoEndpoint)
	if err != nil {
		log.Info(err)
		return nil, errors.New("Get session faild" + r.mgoEndpoint)
	}
	//defer session.Close()
	var mesgs DataMessage
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"method": "datafilereg", "param.filehash": guid}).One(&mesgs)
	}
	session.SetCollection("pre_datafilereg", query)
	reqs := &util.DataDBInfo{
		Filehash: mesgs.Param.Filehash,
		Username:mesgs.Param.Info.Username,
		Filename:mesgs.Param.Info.Filename,
		Filesize:mesgs.Param.Info.Filesize,
		Filepolicy:mesgs.Param.Info.Filepolicy,
		Filenumber:mesgs.Param.Info.Filenumber,
		Simorass:mesgs.Param.Info.Simorass,
		Optype:mesgs.Param.Info.Optype,
		Storeaddr:mesgs.Param.Info.Storeaddr,
	}

	return reqs, err
}
