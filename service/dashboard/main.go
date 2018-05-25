package main

import (
	log "github.com/cihub/seelog"
	"github.com/micro/go-micro"
	dashboard_proto "github.com/bottos-project/magiccube/service/dashboard/proto"
	"golang.org/x/net/context"
	"github.com/bottos-project/magiccube/tools/db/mongodb"
	"gopkg.in/mgo.v2/bson"
	"github.com/bottos-project/magiccube/service/common/bean"
	"time"
	"github.com/bottos-project/magiccube/config"
	"github.com/bottos-project/magiccube/service/query"
	"os"
	"encoding/json"
)

type Dashboard struct {}

func (u *Dashboard) GetNodeInfos(ctx context.Context, req *dashboard_proto.GetNodeInfosRequest, rsp *dashboard_proto.GetNodeInfosResponse) error {
	//var ret []bean.NodeBean
	//var mgo = mgo.Session()
	//defer mgo.Close()
	//mgo.DB(config.DB_NAME).C("Messages").Find(&bson.M{"type": "nodeinforeg"}).All(&ret)
	//var rows = []*dashboard_proto.NodeInfoData{}
	//for _, v := range ret {
	//	rows = append(rows, &dashboard_proto.NodeInfoData{
	//		Ip : v.Data.BasicInfo.NodeIP,
	//		Port : v.Data.BasicInfo.NodePort,
	//		Address : v.Data.BasicInfo.NodeAddress,
	//	})
	//}
	//
	//log.Info(rows)
	//rsp.Code = 0
	//rsp.Data = rows
	//rsp.Msg = "OK"
	return nil
}

func (u *Dashboard) GetTxList(ctx context.Context, req *dashboard_proto.GetTxListRequest, rsp *dashboard_proto.GetTxListResponse) error {
	var pageNum, pageSize, skip int= 1, 20, 0
	if req.PageNum > 0 {
		pageNum = int(req.PageNum)
	}

	if req.PageSize > 0 && req.PageSize <= 50{
		pageSize = int(req.PageSize)
	}

	skip = (pageNum - 1) *  pageSize

	var sort string= "-create_time"
	if req.Sort == "asc" {
		sort = "create_time"
	}
	var ret []bean.Tx
	var mgo = mgo.Session()
	defer mgo.Close()
	count, err:=mgo.DB(config.DB_NAME).C("Transactions").Find(&bson.M{"method": "buydata"}).Count()
	if err != nil {
		log.Error(err)
	}
	mgo.DB(config.DB_NAME).C("Transactions").Find(&bson.M{"method": "buydata"}).Sort(sort).Limit(pageSize).Skip(skip).All(&ret)
	log.Info(ret)
	var rows = []*dashboard_proto.Tx{}
	var ret2 = bean.AssetBean{}
	for _, v := range ret {
		err := mgo.DB(config.DB_NAME).C("pre_assetreg").Find(&bson.M{"param.assetid": v.Param.Info.AssetId, "create_time": bson.M{"$lt": v.CreateTime}}).Sort("-create_time").Limit(1).One(&ret2)
		if err != nil {
			log.Error(err)
		}

		rows = append(rows, &dashboard_proto.Tx{
			TransactionId: v.TransactionId,
			From: v.Param.Info.Username,
			To: ret2.Param.Info.UserName,
			Price: ret2.Param.Info.Price,
			AssetType: ret2.Param.Info.AssetType,
			Timestamp: uint64(v.CreateTime.Unix()),
			BlockNumber: v.BlockNumber,
		})
	}

	var data = &dashboard_proto.TxListData{
		PageNum: uint32(pageNum),
		RowCount: uint32(count),
		Row:rows,
	}

	rsp.Data = data
	return nil
}

func (u *Dashboard) GetBlockList(ctx context.Context, req *dashboard_proto.GetBlockListRequest, rsp *dashboard_proto.GetBlockListResponse) error {
	var pageNum, pageSize, skip int= 1, 15, 0
	if req.PageNum > 0 {
		pageNum = int(req.PageNum)
	}

	if req.PageSize > 0 && req.PageSize <= 20 {
		pageSize = int(req.PageSize)
	}

	skip = (pageNum - 1) *  pageSize

	var sort string= "-block_number"
	if req.Sort == "asc" {
		sort = "block_number"
	}

	var mgo = mgo.Session()
	defer mgo.Close()
	count, err:= mgo.DB(config.DB_NAME).C("Blocks").Find(nil).Count()
	log.Info(count)
	if err != nil {
		log.Error(err)
	}
	var ret []bean.Block

	mgo.DB(config.DB_NAME).C("Blocks").Find(nil).Sort(sort).Skip(skip).Limit(pageSize).All(&ret)

	var rows = []*dashboard_proto.Block{}

	for _, v := range ret {
		rows = append(rows, &dashboard_proto.Block{
			BlockNumber: v.BlockNumber,
			BlockHash: v.BlockHash,
			PrevBlockHash:v.PrevBlockHash,
			MerkleRoot:v.MerkleRoot,
			Delegate:v.Delegate,
			Timestamp:v.Timestamp,
			TxNum:uint32(len(v.Transactions)),
		})
	}
	var data = &dashboard_proto.BlockData{
		PageNum: uint64(pageNum),
		RowCount:uint64(count),
		Row:rows,
	}

	rsp.Data = data
	return nil
}

func (u *Dashboard) GetBlockInfo(ctx context.Context, req *dashboard_proto.GetBlockInfoRequest, rsp *dashboard_proto.GetBlockInfoResponse) error {
	var mgo = mgo.Session()
	defer mgo.Close()

	var ret *bean.Block
	mgo.DB(config.DB_NAME).C("Blocks").Find(bson.M{"block_number":req.BlockNumber}).One(&ret)
	var block = &dashboard_proto.Block{
		BlockNumber:ret.BlockNumber,
		BlockHash:ret.BlockHash,
		MerkleRoot:ret.MerkleRoot,
		PrevBlockHash:ret.PrevBlockHash,
		TxNum: uint32(len(ret.Transactions)),
		Delegate:ret.Delegate,
		Timestamp:ret.Timestamp,
	}


	var ret2 []*bean.Transaction
	mgo.DB(config.DB_NAME).C("Transactions").Find(bson.M{"block_number":req.BlockNumber}).All(&ret2)

	var rows []*dashboard_proto.TxList

	for _, v := range ret2 {
		json_buf,_ := json.Marshal(v.Param)
		rows = append(rows, &dashboard_proto.TxList{
			TransactionId: v.TransactionId,
			SequenceNum: v.SequenceNum,
			CursorNum:v.CursorNum,
			CursorLabel:v.CursorLabel,
			Lifetime:v.Lifetime,
			Sender:v.Sender,
			Contract:v.Contract,
			Method:v.Method,
			Param: string(json_buf),
			Signature:v.Signature,
			CreateTime: uint64(v.CreateTime.Unix()),
		})
	}

	var data = &dashboard_proto.BlockInfoData{
		Block: block,
		TxList: rows,
	}

	rsp.Data = data
	return nil
}

func (u *Dashboard) GetRequirementNumByDay(ctx context.Context, req *dashboard_proto.GetRequirementNumByDayRequest, rsp *dashboard_proto.GetRequirementNumByDayResponse) error {
	timeSlice := getRecent7DayTimeSlice()
	log.Info(timeSlice)
	var data []*dashboard_proto.RequirementNumByDayData
	var mgo = mgo.Session()
	defer mgo.Close()
	for i:=0; i<len(timeSlice)-1; i++ {
		count, err := mgo.DB(config.DB_NAME).C("pre_datareqreg").Find(bson.M{"param.info.optype": bson.M{"$in": []int32{1,2}},"create_time": bson.M{"$gte": query.TimestampToUTC(int64(timeSlice[i])), "$lt": query.TimestampToUTC(int64(timeSlice[i+1]))}}).Count()
		if err!= nil {
			log.Error(err)
		}
		log.Info(count)
		data = append(data, &dashboard_proto.RequirementNumByDayData{
			Time:int64(timeSlice[i]),
			Count:int64(count),
		})
	}
	rsp.Data = data
	return nil
}

func (u *Dashboard) GetAssetNumByDay(ctx context.Context, req *dashboard_proto.GetAssetNumByDayRequest, rsp *dashboard_proto.GetAssetNumByDayResponse) error {
	timeSlice :=getRecent7DayTimeSlice()
	log.Info(timeSlice)
	var data []*dashboard_proto.AssetNumByDayData
	var mgo = mgo.Session()
	defer mgo.Close()
	for i:=0; i<len(timeSlice)-1; i++ {
		count, err := mgo.DB(config.DB_NAME).C("pre_assetreg").Find(bson.M{"param.info.optype": bson.M{"$in": []int32{1,2}},"create_time": bson.M{"$gte": query.TimestampToUTC(int64(timeSlice[i])), "$lt": query.TimestampToUTC(int64(timeSlice[i+1]))}}).Count()
		if err!= nil {
			log.Error(err)
		}

		data = append(data, &dashboard_proto.AssetNumByDayData{
			Time:int64(timeSlice[i]),
			Count:int64(count),
		})
	}
	rsp.Data = data
	return nil
}

func (u *Dashboard) GetAccountNumByDay(ctx context.Context, req *dashboard_proto.GetAccountNumByDayRequest, rsp *dashboard_proto.GetAccountNumByDayResponse) error {
	timeSlice :=getRecent7DayTimeSlice()
	log.Info(timeSlice)
	var data []*dashboard_proto.AccountNumByDayData
	var mgo = mgo.Session()
	defer mgo.Close()
	for i:=0; i<len(timeSlice)-1; i++ {
		count, err :=mgo.DB(config.DB_NAME).C("Accounts").Find(bson.M{"create_time": bson.M{"$gte": query.TimestampToUTC(int64(timeSlice[i])), "$lt": query.TimestampToUTC(int64(timeSlice[i+1]))}}).Count()
		if err!= nil {
			log.Error(err)
		}

		data = append(data, &dashboard_proto.AccountNumByDayData{
			Time: int64(timeSlice[i]),
			Count:int64(count),
		})
	}

	rsp.Data = data
	return nil
}

func (u *Dashboard) GetTxNumByDay(ctx context.Context, req *dashboard_proto.GetTxNumByDayRequest, rsp *dashboard_proto.GetTxNumByDayResponse) error {
	timeSlice :=getRecent7DayTimeSlice()
	log.Info(timeSlice)
	var data []*dashboard_proto.TxNumByDayData
	var mgo = mgo.Session()
	defer mgo.Close()
	for i:=0; i<len(timeSlice)-1; i++ {
		count, err := mgo.DB(config.DB_NAME).C("Transactions").Find(bson.M{"method": "buydata", "create_time": bson.M{"$gte": query.TimestampToUTC(int64(timeSlice[i])), "$lt": query.TimestampToUTC(int64(timeSlice[i+1]))}}).Count()
		if err!= nil {
			log.Error(err)
		}

		data = append(data, &dashboard_proto.TxNumByDayData{
			Time:int64(timeSlice[i]),
			Count:int64(count),
		})
	}

	rsp.Data = data
	return nil
}

func (u *Dashboard) GetTxAmountByDay(ctx context.Context, req *dashboard_proto.GetTxAmountByDayRequest, rsp *dashboard_proto.GetTxAmountByDayResponse) error {
	timeSlice :=getRecent7DayTimeSlice()
	log.Info(timeSlice)
	var ret []bean.Tx

	var data []*dashboard_proto.TxAmountByDay
	var mgo = mgo.Session()
	defer mgo.Close()

	for i:=0; i<len(timeSlice)-1; i++ {
		var ret2 bean.AssetBean
		var amount uint64 = 0

		err :=mgo.DB(config.DB_NAME).C("Transactions").Find(bson.M{"method": "buydata", "create_time": bson.M{"$gte": query.TimestampToUTC(int64(timeSlice[i])), "$lt": query.TimestampToUTC(int64(timeSlice[i+1]))}}).All(&ret)
		if err!= nil {
			log.Error(err)
		}

		for _, v := range ret {
			mgo.DB(config.DB_NAME).C("pre_assetreg").Find(bson.M{"data.asset_id":v.Param.Info.AssetId, "create_time": bson.M{"$lt": v.CreateTime}}).Sort("-create_time").Limit(1).One(&ret2)
			amount += ret2.Param.Info.Price
		}

		data = append(data, &dashboard_proto.TxAmountByDay{
			Time:int64(timeSlice[i]),
			Count:int64(amount),
		})
	}

	rsp.Data = data
	return nil
}

func (u *Dashboard) GetTxNum(ctx context.Context, req *dashboard_proto.GetTxNumRequest, rsp *dashboard_proto.GetTxNumResponse) error {
	var mgo = mgo.Session()
	defer mgo.Close()

	count, err := mgo.DB(config.DB_NAME).C("Transactions").Find(bson.M{"method": "buydata"}).Count()
	if err != nil {
		log.Info(err)
	}

	rsp.Data = &dashboard_proto.Num{
		Num: int64(count),
	}
	return nil
}

func (u *Dashboard) GetTxAmount(ctx context.Context, req *dashboard_proto.GetTxAmountRequest, rsp *dashboard_proto.GetTxAmountResponse) error {
	var ret []bean.Tx
	var mgo = mgo.Session()
	defer mgo.Close()
	mgo.DB(config.DB_NAME).C("Transactions").Find(&bson.M{"method": "buydata"}).All(&ret)
	log.Info(ret)
	var ret2 bean.AssetBean
	var amount uint64 = 0
	for _, v := range ret {
		mgo.DB(config.DB_NAME).C("pre_assetreg").Find(bson.M{"data.asset_id":v.Param.Info.AssetId, "create_time": bson.M{"$lt": v.CreateTime}}).Sort("-create_time").Limit(1).One(&ret2)
		amount += ret2.Param.Info.Price
	}

	rsp.Code = 1
	rsp.Data = &dashboard_proto.Num{
		Num: int64(amount),
	}
	return nil
}

func (u *Dashboard) GetAllTypeTotal(ctx context.Context, req *dashboard_proto.GetAllTypeTotalRequest, rsp *dashboard_proto.GetAllTypeTotalResponse) error {
	var accountNum, assetNum, requirementNum, txNum int= 0, 0, 0, 0
	var txAmount uint64 = 0
	min, max := query.TodayTimeSolt()
	accountNum = accountNum + query.AccountNum(min, max)
	assetNum = assetNum + query.AssetNum(min, max)
	requirementNum = requirementNum + query.RequirementNum(min, max)
	txAmount = txAmount + query.TxAmount(min, max)
	txNum = txNum + query.TxNum(min, max)
	log.Info(accountNum, assetNum, requirementNum, txAmount, txNum)
	var mgo = mgo.Session()
	defer mgo.Close()

	var ret []bean.RecordNum
	err := mgo.DB(config.DB_NAME).C("rec_num").Find(nil).All(&ret);
	if err != nil {
		log.Error(err.Error())
	}

	//job := &mgo2.MapReduce{
	//	Map:    "function() { emit(this.id_, this.asset_num) }",
	//	Reduce: "function(key, values) { return Array.sum(values) }",
	//}
	//var result []struct {
	//	Id    int "_id"
	//	Value int
	//}
	//_, err = mgo.DB(config.DB_NAME).C("record_num_log").Find(nil).MapReduce(job, &result)
	//if err != nil {
	//	panic(err)
	//}
	//log.Info(result)

	for _, v := range ret {
		accountNum = accountNum + v.AccountNum
		assetNum = assetNum + v.AssetNum
		requirementNum = requirementNum + v.RequirementNum
		txAmount = txAmount + v.TxAmount
		txNum = txNum + v.TxNum
	}

	var data []*dashboard_proto.AllTypeData

	data = append(data,&dashboard_proto.AllTypeData{
		Type:"AccountNum",
		Total:int64(accountNum),
	})
	data = append(data,&dashboard_proto.AllTypeData{
		Type:"AssetNum",
		Total:int64(assetNum),
	})
	data = append(data,&dashboard_proto.AllTypeData{
		Type:"RequirementNum",
		Total:int64(requirementNum),
	})
	data = append(data,&dashboard_proto.AllTypeData{
		Type:"TxAmount",
		Total:int64(txAmount),
	})
	data = append(data,&dashboard_proto.AllTypeData{
		Type:"TxNum",
		Total:int64(txNum),
	})

	rsp.Code = 0
	rsp.Data = data
	return nil
}


func getRecent7DayTimeSlice() []int {
	t := time.Now()
	loc, _ := time.LoadLocation("Local")
	tm := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, loc)
	unixTime := int(tm.Unix()) + 1
	//tm := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc)
	//unixTime := int(tm.Unix())
	var timeSlice []int;
	for i:=7; i>=0; i-- {
		timeSlice = append(timeSlice, unixTime - i*24*60*60)
	}
	return timeSlice
}

func init() {
	logger, err := log.LoggerFromConfigAsFile("./config/dash-log.xml")
	if err != nil{
		log.Error(err)
		panic(err)
	}
	defer logger.Flush()
	log.ReplaceLogger(logger)
}

func main() {

	service := micro.NewService(
		micro.Name("go.micro.srv.v3.dashboard"),
		micro.Version("3.0.0"),
	)

	service.Init()

	dashboard_proto.RegisterDashboardHandler(service.Server(), new(Dashboard))

	if err := service.Run(); err != nil {
		os.Exit(1)
	}
}
