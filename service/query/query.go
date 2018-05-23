package query

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"github.com/bottos-project/bottos/tools/db/mongodb"
	log "github.com/cihub/seelog"
	"github.com/bottos-project/bottos/service/common/bean"
	"github.com/bottos-project/bottos/config"
)

func TimestampToUTC(timestamp int64) time.Time {
	date := time.Unix(timestamp, 0)
	local1, _ := time.LoadLocation("UTC")
	return date.In(local1)
}

func TxNum(min int64, max int64) int {
	var mgo = mgo.Session()
	defer mgo.Close()
	count, err := mgo.DB(config.DB_NAME).C("Transactions").Find(bson.M{"method": "buydata", "create_time": bson.M{"$gte": TimestampToUTC(min), "$lte": TimestampToUTC(max)}}).Count()
	if err != nil {
		log.Error(err.Error())
	}
	return count
}

func TxAmount(min int64, max int64) uint64 {
	var amount uint64 = 0
	var ret []bean.Tx


	var mgo = mgo.Session()
	defer mgo.Close()

	err :=mgo.DB(config.DB_NAME).C("Transactions").Find(bson.M{"method": "buydata", "create_time": bson.M{"$gte": TimestampToUTC(min), "$lte": TimestampToUTC(max)}}).All(&ret)
	if err!= nil {
		log.Error(err)
	}
	log.Info(ret)

	for _, v := range ret {
		var ret2 bean.AssetBean
		log.Info(v.Param.Info.AssetId)
		mgo.DB(config.DB_NAME).C("pre_assetreg").Find(bson.M{"param.assetid":v.Param.Info.AssetId, "create_time": bson.M{"$lt": v.CreateTime}}).Sort("-create_time").Limit(1).One(&ret2)
		log.Info(ret2)
		amount +=  ret2.Param.Info.Price
	}

	return amount
}

func RequirementNum(min int64, max int64) int {
	var mgo = mgo.Session()
	defer mgo.Close()
	var count = 0;
	count, err := mgo.DB(config.DB_NAME).C("pre_datareqreg").Find(bson.M{"create_time": bson.M{"$gte": TimestampToUTC(min), "$lt": TimestampToUTC(max)}}).Count()
	if err!= nil {
		log.Error(err)
	}
	return count
}

func AssetNum(min int64, max int64) int {
	var mgo = mgo.Session()
	defer mgo.Close()
	var count = 0;
	count, err :=mgo.DB(config.DB_NAME).C("pre_assetreg").Find(bson.M{"create_time": bson.M{"$gte": TimestampToUTC(min), "$lte": TimestampToUTC(max)}}).Count()
	if err!= nil {
		log.Error(err)
	}
	return count
}

func AccountNum(min int64, max int64) int {
	var mgo = mgo.Session()
	defer mgo.Close()
	var count = 0;
	count, err :=mgo.DB(config.DB_NAME).C("Accounts").Find(bson.M{"create_time": bson.M{"$gte": TimestampToUTC(min), "$lte": TimestampToUTC(max)}}).Count()
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
