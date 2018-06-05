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

package query

import (
	"github.com/bottos-project/magiccube/config"
	"github.com/bottos-project/magiccube/service/common/bean"
	"github.com/bottos-project/magiccube/tools/db/mongodb"
	log "github.com/cihub/seelog"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// TimestampToUTC to utc timer
func TimestampToUTC(timestamp int64) time.Time {
	date := time.Unix(timestamp, 0)
	local1, _ := time.LoadLocation("UTC")
	return date.In(local1)
}

// TxNum number
func TxNum(min int64, max int64) int {
	var mgo = mgo.Session()
	defer mgo.Close()
	count, err := mgo.DB(config.DB_NAME).C("Transactions").Find(bson.M{"method": "buydata", "create_time": bson.M{"$gte": TimestampToUTC(min), "$lte": TimestampToUTC(max)}}).Count()
	if err != nil {
		log.Error(err.Error())
	}
	return count
}

//TxAmount tx amount
func TxAmounte(min int64, max int64) uint64 {
	var amount uint64
	var ret []bean.Tx

	var mgo = mgo.Session()
	defer mgo.Close()

	err := mgo.DB(config.DB_NAME).C("Transactions").Find(bson.M{"method": "buydata", "create_time": bson.M{"$gte": TimestampToUTC(min), "$lte": TimestampToUTC(max)}}).All(&ret)
	if err != nil {
		log.Error(err)
	}
	log.Info(ret)

	for _, v := range ret {
		var ret2 bean.AssetBean
		log.Info(v.Param.Info.AssetId)
		mgo.DB(config.DB_NAME).C("pre_assetreg").Find(bson.M{"param.assetid": v.Param.Info.AssetId, "create_time": bson.M{"$lt": v.CreateTime}}).Sort("-create_time").Limit(1).One(&ret2)
		log.Info(ret2)
		amount += ret2.Param.Info.Price
	}

	return amount
}

//RequirementNum reqirement number
func RequirementNum(min int64, max int64) int {
	var mgo = mgo.Session()
	defer mgo.Close()
	count := 0
	count, err := mgo.DB(config.DB_NAME).C("pre_datareqreg").Find(bson.M{"create_time": bson.M{"$gte": TimestampToUTC(min), "$lt": TimestampToUTC(max)}}).Count()
	if err != nil {
		log.Error(err)
	}
	return count
}

//AssetNum asset number
func AssetNum(min int64, max int64) int {
	var mgo = mgo.Session()
	defer mgo.Close()
	count := 0
	count, err := mgo.DB(config.DB_NAME).C("pre_assetreg").Find(bson.M{"create_time": bson.M{"$gte": TimestampToUTC(min), "$lte": TimestampToUTC(max)}}).Count()
	if err != nil {
		log.Error(err)
	}
	return count
}

//AccountNum account number
func AccountNum(min int64, max int64) int {
	var mgo = mgo.Session()
	defer mgo.Close()
	count := 0
	count, err := mgo.DB(config.DB_NAME).C("Accounts").Find(bson.M{"create_time": bson.M{"$gte": TimestampToUTC(min), "$lte": TimestampToUTC(max)}}).Count()
	if err != nil {
		log.Error(err)
	}
	return count
}

// YesterdayTimeSlot yesterday time slot
func YesterdayTimeSlot() (int64, int64) {
	timeStr := time.Now().Format("2006-01-02")
	//t, _ := time.Parse("2006-01-02", timeStr)
	loc, _ := time.LoadLocation("Local") //重要：获取时区
	theTime, _ := time.ParseInLocation("2006-01-02", timeStr, loc)
	return theTime.Unix() - 24*60*60, theTime.Unix() - 1
}
//TodayTimeSolt is to solt time
func TodayTimeSolt() (int64, int64) {
	timeStr := time.Now().Format("2006-01-02")
	loc, _ := time.LoadLocation("Local") //重要：获取时区
	theTime, _ := time.ParseInLocation("2006-01-02", timeStr, loc)
	//t, _ := time.Parse("2006-01-02", timeStr)
	return theTime.Unix(), theTime.Unix() + 24*60*60 - 1
}
