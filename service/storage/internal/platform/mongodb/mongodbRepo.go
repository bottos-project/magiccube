package mongodb

import (
	"encoding/json"
	"errors"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoRepository struct {
	mgoEndpoint string
}

//NewMongoRepository creates a new MongoRepository
func NewMongoRepository(endpoint string) *MongoRepository {
	return &MongoRepository{mgoEndpoint: endpoint}
}

type MongoContext struct {
	mgoSession *mgo.Session
}

func GetSession(url string) (*MongoContext, error) {
	if url == "" {
		return nil, errors.New("invalid para url")
	}
	var err error
	mgoSession, err := mgo.Dial(url)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Dial faild" + url)
	}
	return &MongoContext{mgoSession.Clone()}, nil
}
func (c *MongoContext) GetCollection(db string, collection string) *mgo.Collection {
	session := c.mgoSession
	defer session.Close()
	collects := session.DB("bottos").C(collection)
	return collects
}
func (c *MongoContext) SetCollection(collection string, s func(*mgo.Collection) error) error {
	session := c.mgoSession
	defer session.Close()
	collects := session.DB("bottos").C(collection)
	return s(collects)
}

func (c *MongoContext) SetCollectionCount(collection string, s func(*mgo.Collection) (int, error)) (int, error) {
	session := c.mgoSession
	defer session.Close()
	collects := session.DB("bottos").C(collection)
	return s(collects)
}
func (c *MongoContext) SetCollectionByDB(db string, collection string, s func(*mgo.Collection) error) error {
	session := c.mgoSession
	defer session.Close()
	collects := session.DB(db).C(collection)
	return s(collects)
}

// CollectionExists returns true if the collection name exists in the specified database.
func (c *MongoContext) isCollectionExists(useCollection string) bool {
	session := c.mgoSession
	database := session.DB("bottos")
	collections, err := database.CollectionNames()

	if err != nil {
		return false
	}

	for _, collection := range collections {
		if collection == useCollection {
			return true
		}
	}

	return false
}

// ToString converts the quer map to a string.
func ToString(queryMap interface{}) string {
	json, err := json.Marshal(queryMap)
	if err != nil {
		return ""
	}

	return string(json)
}

// ToStringD converts bson.D to a string.
func ToStringD(queryMap bson.D) string {
	json, err := json.Marshal(queryMap)
	if err != nil {
		return ""
	}

	return string(json)
}
