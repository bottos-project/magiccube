package query

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"github.com/code/bottos/tools/db/mongodb"
	log "github.com/jeanphorn/log4go"
	"github.com/code/bottos/service/bean"
	"github.com/code/bottos/config"
)

func TimestampToUTC(timestamp int64) time.Time {
	date := time.Unix(timestamp, 0)
	local1, _ := time.LoadLocation("UTC")
	return date.In(local1)
}

func TxNum(min int64, max int64) int {
	var mgo = mgo.Session()
	defer mgo.Close()
	count, err := mgo.DB("bottos").C("Messages").Find(bson.M{"type": "datapurchase", "createdAt": bson.M{"$gte": TimestampToUTC(min), "$lte": TimestampToUTC(max)}}).Count()
	if err != nil {
		log.Error(err.Error())
	}
	return count
}

func TxAmount(min int64, max int64) int {
	var amount int = 0
	var ret []bean.TxBean
	var ret2 []bean.AssetBean

	var mgo = mgo.Session()
	defer mgo.Close()

	log.Info(TimestampToUTC(min))

	err :=mgo.DB("bottos").C("Messages").Find(bson.M{"type": "datapurchase", "createdAt": bson.M{"$gte": TimestampToUTC(min), "$lte": TimestampToUTC(max)}}).All(&ret)
	if err!= nil {
		log.Error(err)
	}
	var assetIdArr []string;
	for _, v := range ret {
		assetIdArr = append(assetIdArr, v.Data.BasicInfo.AssetID)
	}
	log.Info(len(assetIdArr))
	mgo.DB(config.DB_NAME).C("Messages").Find(bson.M{"type": "assetreg", "data.asset_id": bson.M{"$in": assetIdArr}}).All(&ret2)
	for _, v := range ret2 {
		amount +=  int(v.Data.BasicInfo.Price)
	}
	log.Info(amount)
	log.Info(TimestampToUTC(max))
	return amount
}

func RequirementNum(min int64, max int64) int {
	var mgo = mgo.Session()
	defer mgo.Close()
	var count = 0;
	count, err := mgo.DB(config.DB_NAME).C("Messages").Find(bson.M{"type": "datareqreg", "createdAt": bson.M{"$gte": TimestampToUTC(min), "$lt": TimestampToUTC(max)}}).Count()
	if err!= nil {
		log.Error(err)
	}
	return count
}

func AssetNum(min int64, max int64) int {
	var mgo = mgo.Session()
	defer mgo.Close()
	var count = 0;
	count, err :=mgo.DB(config.DB_NAME).C("Messages").Find(bson.M{"type": "assetreg", "createdAt": bson.M{"$gte": TimestampToUTC(min), "$lte": TimestampToUTC(max)}}).Count()
	if err!= nil {
		log.Error(err)
	}
	return count
}

func AccountNum(min int64, max int64) int {
	var mgo = mgo.Session()
	defer mgo.Close()
	var count = 0;
	count, err :=mgo.DB(config.DB_NAME).C("Messages").Find(bson.M{"type": "newaccount", "createdAt": bson.M{"$gte": TimestampToUTC(min), "$lte": TimestampToUTC(max)}}).Count()
	if err!= nil {
		log.Error(err)
	}
	return count
}

func YesterdayTimeSlot() (int64, int64) {
	timeStr := time.Now().Format("2006-01-02")
	//t, _ := time.Parse("2006-01-02", timeStr)
	loc, _ := time.LoadLocation("Local")                                //重要：获取时区
	theTime, _ := time.ParseInLocation("2006-01-02", timeStr, loc)
	return theTime.Unix()-24*60*60, theTime.Unix()-1
}

func TodayTimeSolt() (int64, int64) {
	timeStr := time.Now().Format("2006-01-02")
	loc, _ := time.LoadLocation("Local")                                //重要：获取时区
	theTime, _ := time.ParseInLocation("2006-01-02", timeStr, loc)
	//t, _ := time.Parse("2006-01-02", timeStr)
	return theTime.Unix(), theTime.Unix()+ 24*60*60 -1
}
