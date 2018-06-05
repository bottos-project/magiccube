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
	"github.com/bottos-project/magiccube/config"
	"github.com/bottos-project/magiccube/service/common/bean"
	"github.com/bottos-project/magiccube/service/query"
	"github.com/bottos-project/magiccube/tools/db/mongodb"
	log "github.com/cihub/seelog"
	"github.com/robfig/cron"
	"time"
)

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
	//DiurnalStatis()
	c := cron.New()
	spec := "0, 5, 0, *, *, *" //每天凌晨0:05执行一次
	c.AddFunc(spec, DiurnalStatis)
	c.Start()
	select {} //阻塞主线程不退出
}

func DiurnalStatis() {
	log.Info("Execution of tasks!!!")
	min, max := query.YesterdayTimeSlot()
	var d bean.RecordNum
	d.TxNum = query.TxNum(min, max)
	d.TxAmount = query.TxAmount(min, max)
	d.RequirementNum = query.RequirementNum(min, max)
	d.AssetNum = query.AssetNum(min, max)
	d.AccountNum = query.AccountNum(min, max)
	d.Date = time.Unix(min, 0).Format("2006-01-02")
	d.Timestamp = int(min)
	d.CreateTime = time.Now()

	var mgo = mgo.Session()
	defer mgo.Close()
	err := mgo.DB(config.DB_NAME).C("rec_num").Insert(d)
	if err != nil {
		log.Error(err.Error())
	}
}
