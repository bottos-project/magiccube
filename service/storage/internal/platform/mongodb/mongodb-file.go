/*Copyright 2017~2022 The Bottos Authors
  This file is part of the Bottos Data Exchange Client
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
	"fmt"

	"github.com/bottos-project/magiccube/service/storage/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// FileMessage is definition of file msg
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

// CallGetUserFileList is to get user file list
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
			FileHash:          mesgs[i].Data.FileHash,
			Username:          mesgs[i].Data.BasicInfo.UserName,
			FileName:          mesgs[i].Data.BasicInfo.FileName,
			FileSize:          mesgs[i].Data.BasicInfo.FileSize,
			FileNumber:        mesgs[i].Data.BasicInfo.FileNumber,
			FilePolicy:        mesgs[i].Data.BasicInfo.FilePolicy,
			AuthorizedStorage: mesgs[i].Data.BasicInfo.AuthPath}
		reqs = append(reqs, dbtag)
	}

	fmt.Println(reqs)
	return reqs, nil
}
