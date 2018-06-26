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
package main

import (
	"github.com/bottos-project/magiccube/service/common/bean"
	"github.com/bottos-project/magiccube/tools/db/mongodb"
	log "github.com/cihub/seelog"
	"github.com/robfig/cron"
	"gopkg.in/mgo.v2/bson"
)

var branchTable = []string{"favoritepro", "datareqreg", "assetreg", "presale", "datafilereg"}
var prefix = "pre_"

//Favoritepro struct
type Favoritepro struct {
	GoodsId   string `bson:"goodsid"`
	GoodsTyte string `bson:"goodstype"`
	OpTyte    uint32 `bson:"optype"`
}

//Datareqreg struct
type Datareqreg struct {
	DataReqId string `bson:"datareqid"`
	Info      `bson:"info"`
}

//Presale struct
type Presale struct {
	DataPresaleId string `bson:"datapresaleid"`
	Info          `bson:"info"`
}

//Assetreg struct
type Assetreg struct {
	AssetId string `bson:"assetid"`
	Info    `bson:"info"`
}

//File struct
type File struct {
	FileHash string `bson:"filehash"`
	Info     `bson:"info"`
}

//Info struct
type Info struct {
	OpTyte uint32 `bson:"optype"`
}

//RecMessageId struct
type RecMessageId struct {
	MessageID string `bson:"message_id"`
}

func init() {
	logger, err := log.LoggerFromConfigAsFile("./config/task-log.xml")
	if err != nil {
		log.Error(err)
		panic(err)
	}
	defer logger.Flush()
	log.ReplaceLogger(logger)
}

func main() {
	c := cron.New()
	spec := "@every 3s"
	c.AddFunc(spec, BranchTable)
	c.Start()
	select {}
}

//BranchTable is to branch
func BranchTable() {
	log.Info("Execution of tasks!!!")
	var mgo = mgo.Session()
	defer mgo.Close()

	var recMsg RecMessageId
	mgo.DB("bottos").C("rec_id").Find(nil).One(&recMsg)

	var part bean.Transaction
	mgo.DB("bottos").C("Transactions").Find(nil).Sort("-_id").Limit(1).One(&part)
	log.Info("part-last-id:", part.ID)

	if recMsg.MessageID == part.ID.Hex() {
		return
	}

	var where = bson.M{"_id": bson.M{"$lte": bson.ObjectIdHex(part.ID.Hex())}, "method": bson.M{"$in": branchTable}}
	if recMsg.MessageID != "" {
		where = bson.M{"_id": bson.M{"$gt": bson.ObjectIdHex(recMsg.MessageID), "$lte": bson.ObjectIdHex(part.ID.Hex())}, "method": bson.M{"$in": branchTable}}
	}

	var ret []bean.Transaction
	mgo.DB("bottos").C("Transactions").Find(where).All(&ret)

	log.Info(len(ret))
	for _, v := range ret {
		switch v.Method {
		case "favoritepro":
			var favoritepro = &Favoritepro{}
			data, err := bson.Marshal(v.Param)
			if err != nil {
				log.Error(err)
				return
			}
			bson.Unmarshal(data, &favoritepro)

			//if favoritepro.OpTyte == 2 || favoritepro.OpTyte == 3 {
			where = bson.M{"param.goodsid": favoritepro.GoodsId, "param.goodstype": favoritepro.GoodsTyte}
			set := bson.M{"$set": bson.M{"param.optype": 3}}
			mgo.DB("bottos").C(prefix+v.Method).UpdateAll(where, set)
			//}
		case "datareqreg":
			var datareqreg = &Datareqreg{}
			data, err := bson.Marshal(v.Param)
			if err != nil {
				log.Error(err)
				return
			}
			bson.Unmarshal(data, &datareqreg)
			if datareqreg.OpTyte == 2 || datareqreg.OpTyte == 3 {
				where = bson.M{"param.datareqid": datareqreg.DataReqId, "param.info.optype": datareqreg.Info.OpTyte}
				set := bson.M{"$set": bson.M{"param.info.optype": 3}}
				mgo.DB("bottos").C(prefix+v.Method).UpdateAll(where, set)
			}
		case "assetreg":
			var assetreg = &Assetreg{}
			data, err := bson.Marshal(v.Param)
			if err != nil {
				log.Error(err)
				return
			}
			bson.Unmarshal(data, &assetreg)
			if assetreg.OpTyte == 2 || assetreg.OpTyte == 3 {
				where = bson.M{"param.assetid": assetreg.AssetId, "param.info.optype": assetreg.Info.OpTyte}
				set := bson.M{"$set": bson.M{"param.info.optype": 3}}
				mgo.DB("bottos").C(prefix+v.Method).UpdateAll(where, set)
			}
		case "presale":
			var presale = &Presale{}
			data, err := bson.Marshal(v.Param)
			if err != nil {
				log.Error(err)
				return
			}
			bson.Unmarshal(data, &presale)
			mgo.DB("bottos").C(prefix + v.Method).UpdateAll(bson.M{"param.datapresaleid": presale.DataPresaleId}, bson.M{"$inc": bson.M{"param.info.isRead": 0}})
			
			if presale.OpTyte == 2 || presale.OpTyte == 3 {
				where = bson.M{"param.datapresaleid": presale.DataPresaleId, "param.info.optype": presale.Info.OpTyte}
				set := bson.M{"$set": bson.M{"param.info.optype": 3}}
				mgo.DB("bottos").C(prefix+v.Method).UpdateAll(where, set)
			}
		case "datafilereg":
			var file = &File{}
			data, err := bson.Marshal(v.Param)
			if err != nil {
				log.Error(err)
				return
			}
			bson.Unmarshal(data, &file)
			if file.OpTyte == 2 || file.OpTyte == 3 {
				where = bson.M{"param.filehash": file.FileHash, "param.info.optype": file.Info.OpTyte}
				set := bson.M{"$set": bson.M{"param.info.optype": 3}}
				mgo.DB("bottos").C(prefix+v.Method).UpdateAll(where, set)
			}
		}
		mgo.DB("bottos").C(prefix + v.Method).Insert(v)

		//if common_data.Action == "edit" || common_data.Action == "del"{
		//	where = bson.M{"data.id": common_data.Id}
		//	set := bson.M{"$set": bson.M{ "data.action": "del"}}
		//	mgo.DB("zltest").C(v.Type).UpdateAll(where,set);
		//}
		//mgo.DB("zltest").C(v.Type).Insert(v)
	}

	if recMsg.MessageID != "" {
		mgo.DB("bottos").C("rec_id").Update(nil, map[string]interface{}{
			"message_id": part.ID.Hex(),
		})
	} else {
		mgo.DB("bottos").C("rec_id").Insert(map[string]interface{}{
			"message_id": part.ID.Hex(),
		})
	}
}
