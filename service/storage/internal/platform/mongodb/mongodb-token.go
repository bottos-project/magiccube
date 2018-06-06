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

//UserToken is to user token
type UserToken struct {
	ID        bson.ObjectId `bson:"_id"`
	Username  string        `bson:"user_name"`
	Token     string        `bson:"token"`
	CreatedAt time.Time     `bson:"createdAt"`
}

//CallInsertUserToken is to insert user token
func (r *MongoRepository) CallInsertUserToken(username string, token string) (uint32, error) {
	session, err := GetSession(r.mgoEndpoint)
	if err != nil {
		fmt.Println(err)
		return 0, errors.New("Get session faild" + r.mgoEndpoint)
	}

	record := &UserToken{
		ID:       bson.NewObjectId(),
		Username: username,
		Token:    token}
	insert := func(c *mgo.Collection) error {
		return c.Insert(record)
	}
	err = session.SetCollectionByDB("local", "usertoken", insert)
	fmt.Println(err)
	return 1, nil
}

//CallGetUserToken is to get user token
func (r *MongoRepository) CallGetUserToken(username string, token string) (*util.TokenDBInfo, error) {
	session, err := GetSession(r.mgoEndpoint)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Get session faild" + r.mgoEndpoint)
	}
	var mesgs UserToken
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"$or": []bson.M{{"user_name": username}, {"token": token}}}).One(&mesgs)
	}
	session.SetCollectionByDB("local", "usertoken", query)
	tokeninfo := &util.TokenDBInfo{
		Username:   mesgs.Username,
		Token:      mesgs.Token,
		InsertTime: int64(mesgs.CreatedAt.Nanosecond())}

	return tokeninfo, nil
}

//CallDelUserToken is to delete user token
func (r *MongoRepository) CallDelUserToken(username string, token string) (uint32, error) {
	session, err := GetSession(r.mgoEndpoint)
	if err != nil {
		fmt.Println(err)
		return 0, errors.New("Get session faild" + r.mgoEndpoint)
	}

	remove := func(c *mgo.Collection) error {
		return c.Remove(bson.M{"user_name": username, "token": token})
	}
	session.SetCollectionByDB("local", "usertoken", remove)
	return 1, nil
}
