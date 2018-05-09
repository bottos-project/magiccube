package mongodb

import (
	"errors"
	"fmt"

	"time"

	"github.com/bottos-project/bottos/service/storage/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserToken struct {
	ID        bson.ObjectId `bson:"_id"`
	Username  string        `bson:"user_name"`
	Token     string        `bson:"token"`
	CreatedAt time.Time     `bson:"createdAt"`
}

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
func (r *MongoRepository) CallGetUserToken(username string, token string) (*util.TokenDBInfo, error) {
	session, err := GetSession(r.mgoEndpoint)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Get session faild" + r.mgoEndpoint)
	}
	var mesgs UserToken
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"$or": []bson.M{bson.M{"user_name": username}, bson.M{"token": token}}}).One(&mesgs)
	}
	session.SetCollectionByDB("local", "usertoken", query)
	tokeninfo := &util.TokenDBInfo{
		mesgs.Username,
		mesgs.Token,
		int64(mesgs.CreatedAt.Nanosecond())}

	return tokeninfo, nil
}

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
