﻿package main

import (
	log "github.com/jeanphorn/log4go"
	"github.com/micro/go-micro"
	dashboard_proto "github.com/code/bottos/service/dashboard/proto"
	"golang.org/x/net/context"
	"github.com/code/bottos/tools/db/mongodb"
	"gopkg.in/mgo.v2/bson"
	"github.com/code/bottos/service/bean"
	"time"
	"github.com/code/bottos/config"
)


type Dashboard struct {}

func (u *Dashboard) GetNodeInfos(ctx context.Context, req *dashboard_proto.GetNodeInfosRequest, rsp *dashboard_proto.GetNodeInfosResponse) error {
	var ret []bean.NodeBean
	var mgo = mgo.Session()
	defer mgo.Close()
	mgo.DB("bottos").C("Messages").Find(&bson.M{"type": "nodeinforeg"}).All(&ret)
	var rows = []*dashboard_proto.NodeInfoData{}
	for _, v := range ret {
		rows = append(rows, &dashboard_proto.NodeInfoData{
			Ip : v.Data.BasicInfo.NodeIP,
			Port : v.Data.BasicInfo.NodePort,
			Address : v.Data.BasicInfo.NodeAddress,
		})
	}

	log.Info(rows)
	rsp.Code = 0
	rsp.Data = rows
	rsp.Msg = "OK"
	return nil
}

func (u *Dashboard) GetRecentTxList(ctx context.Context, req *dashboard_proto.GetRecentTxListRequest, rsp *dashboard_proto.GetRecentTxListResponse) error {
	log.Info(req.Limit)
	var limit int = 15
	if req.Limit > 0 {
		limit = int(req.Limit)
	}
	var ret []bean.TxBean
	var mgo = mgo.Session()
	defer mgo.Close()
	mgo.DB("bottos").C("Messages").Find(&bson.M{"type": "datapurchase"}).Sort("-createdAt").Limit(limit).All(&ret)
	log.Info(ret)
	var rows = []*dashboard_proto.RecentTxListData{}
	var ret2 = bean.AssetBean{}
	for _, v := range ret {
		log.Info(v.Data.BasicInfo.AssetID)
		err := mgo.DB("bottos").C("Messages").Find(&bson.M{"type": "assetreg", "data.asset_id": v.Data.BasicInfo.AssetID}).One(&ret2)
		if err != nil {
			log.Error(err)
		}

		rows = append(rows, &dashboard_proto.RecentTxListData{
			TransactionId: v.TransactionID,
			From: ret2.Data.BasicInfo.UserName,
			To: v.Data.BasicInfo.UserName,
			Price: ret2.Data.BasicInfo.Price,
			Type: ret2.Data.BasicInfo.FeatureTag,
			Date: v.CreatedAt.String(),
			BlockId: v.BlockNum,
		})
	}

	rsp.Code = 1
	rsp.Data = rows
	return nil
}

func (u *Dashboard) GetRequirementNumByDay(ctx context.Context, req *dashboard_proto.GetRequirementNumByDayRequest, rsp *dashboard_proto.GetRequirementNumByDayResponse) error {
	timeSlice :=getRecent7DayTimeSlice()
	log.Info(timeSlice)
	var data []*dashboard_proto.RequirementNumByDayData
	var mgo = mgo.Session()
	defer mgo.Close()
	for i:=0; i<len(timeSlice)-1; i++ {
		count, err := mgo.DB("bottos").C("Messages").Find(bson.M{"type": "datareqreg", "createdAt": bson.M{"$gte": time.Unix(int64(timeSlice[i]),0), "$lt": time.Unix(int64(timeSlice[i+1]), 0)}}).Count()
		if err!= nil {
			log.Error(err)
		}
		log.Info(count)
		data = append(data, &dashboard_proto.RequirementNumByDayData{
			Time:int64(timeSlice[i]),
			Count:int64(count),
		})
	}

	rsp.Code = 1
	rsp.Data = data
	return nil
}

func (u *Dashboard) GetAllTxNum(ctx context.Context, req *dashboard_proto.GetAllTxNumRequest, rsp *dashboard_proto.GetAllTxNumResponse) error {
	var mgo = mgo.Session()
	defer mgo.Close()
	count, err := mgo.DB("bottos").C("Transactions").Count()
	if err != nil {
		log.Info(err)
	}

	rsp.Code = 1
	rsp.Data = &dashboard_proto.Num{
		Num: int64(count),
	}
	return nil
}

func (u *Dashboard) GetAssetNumByDay(ctx context.Context, req *dashboard_proto.GetAssetNumByDayRequest, rsp *dashboard_proto.GetAssetNumByDayResponse) error {
	timeSlice :=getRecent7DayTimeSlice()
	log.Info(timeSlice)
	var data []*dashboard_proto.AssetNumByDayData
	var mgo = mgo.Session()
	defer mgo.Close()
	for i:=0; i<len(timeSlice)-1; i++ {
		count, err :=mgo.DB("bottos").C("Messages").Find(bson.M{"type": "assetreg", "createdAt": bson.M{"$gte": time.Unix(int64(timeSlice[i]),0), "$lt": time.Unix(int64(timeSlice[i+1]), 0)}}).Count()
		if err!= nil {
			log.Error(err)
		}

		data = append(data, &dashboard_proto.AssetNumByDayData{
			Time:int64(timeSlice[i]),
			Count:int64(count),
		})
	}

	rsp.Code = 1
	rsp.Data = data
	return nil
}

func (u *Dashboard) GetSumTxAmount(ctx context.Context, req *dashboard_proto.GetSumTxAmountRequest, rsp *dashboard_proto.GetSumTxAmountResponse) error {
	var ret []bean.TxBean
	var mgo = mgo.Session()
	defer mgo.Close()
	mgo.DB("bottos").C("Messages").Find(&bson.M{"type": "datapurchase"}).All(&ret)
	log.Info(ret)
	var ret2 bean.AssetBean
	var amount uint64 = 0
	for _, v := range ret {
		mgo.DB("bottos").C("Messages").Find(bson.M{"type": "assetreg", "data.asset_id":v.Data.BasicInfo.AssetID}).One(&ret2)
		amount += ret2.Data.BasicInfo.Price
	}

	rsp.Code = 1
	rsp.Data = &dashboard_proto.Num{
		Num: int64(amount),
	}
	return nil
}

func (u *Dashboard) GetTxNumByDay(ctx context.Context, req *dashboard_proto.GetTxNumByDayRequest, rsp *dashboard_proto.GetTxNumByDayResponse) error {
	timeSlice :=getRecent7DayTimeSlice()
	log.Info(timeSlice)
	var data []*dashboard_proto.TxNumByDayData
	var mgo = mgo.Session()
	defer mgo.Close()
	for i:=0; i<len(timeSlice)-1; i++ {
		count, err := mgo.DB("bottos").C("Transactions").Find(bson.M{"createdAt": bson.M{"$gte": time.Unix(int64(timeSlice[i]),0), "$lt": time.Unix(int64(timeSlice[i+1]), 0)}}).Count()
		if err!= nil {
			log.Error(err)
		}

		data = append(data, &dashboard_proto.TxNumByDayData{
			Time:int64(timeSlice[i]),
			Count:int64(count),
		})
	}

	rsp.Code = 1
	rsp.Data = data
	return nil
}

func getRecent7DayTimeSlice() []int {
	t := time.Now()
	tm := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
	unixTime := int(tm.Unix()) + 1
	var timeSlice []int;
	for i:=7; i>=0; i-- {
		timeSlice = append(timeSlice, unixTime - i*24*60*60)
	}
	return timeSlice
}

func main() {
	log.LoadConfiguration(config.BASE_LOG_CONF)
	defer log.Close()
	log.LOGGER("dashboard.srv")

	service := micro.NewService(
		micro.Name("go.micro.srv.dashboard"),
		micro.Version("2.0.0"),
	)

	service.Init()

	dashboard_proto.RegisterDashboardHandler(service.Server(), new(Dashboard))

	if err := service.Run(); err != nil {
		log.Exit(err)
	}

}
