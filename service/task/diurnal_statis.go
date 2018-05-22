package main

import (
	"github.com/robfig/cron"
	"time"
	"github.com/bottos-project/bottos/tools/db/mongodb"
	log "github.com/cihub/seelog"
	"github.com/bottos-project/bottos/service/query"
	"github.com/bottos-project/bottos/service/bean"
	"github.com/bottos-project/bottos/config"
)

func init() {
	logger, err := log.LoggerFromConfigAsFile("./config/task-log.xml")
	if err != nil{
		log.Error(err)
		panic(err)
	}
	defer logger.Flush()
	log.ReplaceLogger(logger)
}

func main() {
	DiurnalStatis()
	c := cron.New()
	spec := "0, 5, 0, *, *, *" //每天凌晨0:05执行一次
	c.AddFunc(spec, DiurnalStatis)
	c.Start()
	select{} //阻塞主线程不退出
}

func DiurnalStatis() {
	log.Info("Execution of tasks!!!")
	min, max := query.YesterdayTimeSlot()
	var d bean.RecordNumLog
	d.TxNum = query.TxNum(min, max)
	d.TxAmount = query.TxAmount(min, max)
	d.RequirementNum = query.RequirementNum(min, max)
	d.AssetNum = query.AssetNum(min, max)
	d.AccountNum = query.AccountNum(min, max)
	d.Date = time.Unix(min, 0).Format("2006-01-02")
	d.Timestamp = int(min)
	d.CreatedAt = time.Now()

	var mgo = mgo.Session()
	defer mgo.Close()
	err := mgo.DB(config.DB_NAME).C("rec_num").Insert(d)
	if err != nil {
		log.Error(err.Error())
	}
}


