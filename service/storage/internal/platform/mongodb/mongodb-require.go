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
	"time"

	"github.com/bottos-project/magiccube/service/storage/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type RequireMessage struct {
	ID                 bson.ObjectId `bson:"_id,omitempty"`
	MessageID          int           `bson:"message_id"`
	TransactionID      string        `bson:"transaction_id"`
	Authorization      []interface{} `bson:"authorization"`
	HandlerAccountName string        `bson:"handler_account_name"`
	Type               string        `bson:"type"`
	Data               struct {
		DataReqID string `bson:"data_req_id"`
		BasicInfo struct {
			UserName        string `bson:"user_name"`
			SessionID       string `bson:"session_id"`
			RequirementName string `bson:"requirement_name"`
			FeatureTag      uint64 `bson:"feature_tag"`
			SamplePath      string `bson:"sample_path"`
			SampleHash      string `bson:"sample_hash"`
			ExpireTime      uint32 `bson:"expire_time"`
			Price           uint64 `bson:"price"`
			Description     string `bson:"description"`
			PublishDate     uint32 `bson:"publish_date"`
			Signature       string `bson:"signature"`
		} `bson:"basic_info"`
	} `bson:"data"`
	CreatedAt string `bson:"createdAt"`
}

func (r *MongoRepository) CallGetAllRequirementList() ([]*util.RequirementDBInfo, error) {
	session, err := GetSession(r.mgoEndpoint)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Get session faild" + r.mgoEndpoint)
	}
	//defer session.Close()
	fmt.Println(session)
	var mesgs []RequireMessage
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"type": "datareqreg"}).All(&mesgs)
	}
	session.SetCollection("Messages", query)
	fmt.Println(mesgs)
	var reqs = []*util.RequirementDBInfo{}
	for i := 0; i < len(mesgs); i++ {

		dbtag := &util.RequirementDBInfo{
			mesgs[i].Data.DataReqID,
			mesgs[i].Data.BasicInfo.UserName,
			mesgs[i].Data.BasicInfo.RequirementName,
			mesgs[i].Data.BasicInfo.FeatureTag,
			mesgs[i].Data.BasicInfo.SamplePath,
			mesgs[i].Data.BasicInfo.SampleHash,
			mesgs[i].Data.BasicInfo.ExpireTime,
			mesgs[i].Data.BasicInfo.Price,
			mesgs[i].Data.BasicInfo.Description,
			mesgs[i].Data.BasicInfo.PublishDate}
		reqs = append(reqs, dbtag)
	}

	fmt.Println(reqs)
	return reqs, nil
}
func (r *MongoRepository) CallGetUserRequirementList(username string) ([]*util.RequirementDBInfo, error) {
	session, err := GetSession(r.mgoEndpoint)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Get session faild" + r.mgoEndpoint)
	}
	//defer session.Close()
	fmt.Println(session)
	var mesgs []RequireMessage

	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"type": "datareqreg", "data.basic_info.user_name": username}).All(&mesgs)
	}
	session.SetCollection("Messages", query)
	fmt.Println(mesgs)
	var reqs = []*util.RequirementDBInfo{}
	for i := 0; i < len(mesgs); i++ {

		dbtag := &util.RequirementDBInfo{
			mesgs[i].Data.DataReqID,
			mesgs[i].Data.BasicInfo.UserName,
			mesgs[i].Data.BasicInfo.RequirementName,
			mesgs[i].Data.BasicInfo.FeatureTag,
			mesgs[i].Data.BasicInfo.SamplePath,
			mesgs[i].Data.BasicInfo.SampleHash,
			mesgs[i].Data.BasicInfo.ExpireTime,
			mesgs[i].Data.BasicInfo.Price,
			mesgs[i].Data.BasicInfo.Description,
			mesgs[i].Data.BasicInfo.PublishDate}
		reqs = append(reqs, dbtag)
	}

	fmt.Println(reqs)
	return reqs, nil
}

func (r *MongoRepository) CallGetRequirementNumByDay(begin time.Time, end time.Time) (uint64, error) {
	session, err := GetSession(r.mgoEndpoint)
	if err != nil {
		fmt.Println(err)
		return 0, errors.New("Get session faild" + r.mgoEndpoint)
	}

	fmt.Println(begin)
	fmt.Println(end)
	//var mesgs []AssetMessage
	query := func(c *mgo.Collection) (int, error) {
		return c.Find(bson.M{"type": "datareqreg",
			"createdAt": bson.M{"$gt": begin, "$lt": end}}).Count()
	}
	num, err := session.SetCollectionCount("Messages", query)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return uint64(num), nil
}
func (r *MongoRepository) CallGetRequirementListByFeature(featur_tag uint64) ([]*util.RequirementDBInfo, error) {
	session, err := GetSession(r.mgoEndpoint)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Get session faild" + r.mgoEndpoint)
	}
	//defer session.Close()
	fmt.Println(session)
	var mesgs []RequireMessage
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"type": "datareqreg", "data.basic_info.feature_tag": featur_tag}).All(&mesgs)
	}
	session.SetCollection("Messages", query)
	fmt.Println(mesgs)
	var reqs = []*util.RequirementDBInfo{}
	for i := 0; i < len(mesgs); i++ {
		dbtag := &util.RequirementDBInfo{
			mesgs[i].Data.DataReqID,
			mesgs[i].Data.BasicInfo.UserName,
			mesgs[i].Data.BasicInfo.RequirementName,
			mesgs[i].Data.BasicInfo.FeatureTag,
			mesgs[i].Data.BasicInfo.SamplePath,
			mesgs[i].Data.BasicInfo.SampleHash,
			mesgs[i].Data.BasicInfo.ExpireTime,
			mesgs[i].Data.BasicInfo.Price,
			mesgs[i].Data.BasicInfo.Description,
			mesgs[i].Data.BasicInfo.PublishDate}
		reqs = append(reqs, dbtag)
	}

	fmt.Println(reqs)
	return reqs, nil
}
