package main

import (
	log "github.com/cihub/seelog"
	"github.com/micro/go-micro"
	requirement_proto "github.com/bottos-project/magiccube/service/requirement/proto"
	"golang.org/x/net/context"
	"github.com/bottos-project/magiccube/tools/db/mongodb"
	"gopkg.in/mgo.v2/bson"
	"github.com/bottos-project/magiccube/service/common/bean"
	"github.com/bottos-project/magiccube/config"
	"os"
	"github.com/bottos-project/magiccube/service/common/data"
)


type Requirement struct {}

func (u *Requirement) Publish(ctx context.Context, req *requirement_proto.PublishRequest, rsp *requirement_proto.PublishResponse) error {
	i, err := data.PushTransaction(req)
	if err != nil {
		rsp.Code = 3001
		rsp.Msg = err.Error()
	}
	log.Info(i)
	return nil
}

func (u *Requirement) Query(ctx context.Context, req *requirement_proto.QueryRequest, rsp *requirement_proto.QueryResponse) error {

	var pageNum, pageSize, skip int= 1, 20, 0
	if req.PageNum > 0 {
		pageNum = int(req.PageNum)
	}

	if req.PageSize > 0 && req.PageSize < 20{
		pageSize = int(req.PageSize)
	}

	skip = (pageNum - 1) *  pageSize

	var where interface{}
	where = bson.M{"param.info.optype": bson.M{"$in": []int32{1,2}}}
	if len(req.Username) > 0{
		where = bson.M{"param.info.optype": bson.M{"$in": []uint32{1,2}}, "param.info.username": req.Username}
	}

	if len(req.ReqId) > 0 {
		where = bson.M{"param.info.optype": bson.M{"$in": []uint32{1,2}}, "param.datareqid": req.ReqId}
	}

	if req.ReqType > 0 {
		where = bson.M{"param.info.optype": bson.M{"$in": []uint32{1,2}}, "param.info.reqtype": req.ReqType}
	}

	if len(req.Username) > 0 && req.ReqType > 0 {
		where = bson.M{"param.info.optype": bson.M{"$in": []uint32{1,2}}, "param.info.username": req.Username, "param.info.reqtype": req.ReqType}
	}

	var ret []bean.Requirement
	var mgo = mgo.Session()
	defer mgo.Close()
	count, err:= mgo.DB(config.DB_NAME).C("pre_datareqreg").Find(where).Count()
	log.Info(count)
	if err != nil {
		log.Error(err)
	}
	mgo.DB(config.DB_NAME).C("pre_datareqreg").Find(where).Sort("-_id").Skip(skip).Limit(pageSize).All(&ret)

	var rows = []*requirement_proto.RequirementData{}
	for _, v := range ret {

		rows = append(rows, &requirement_proto.RequirementData{
			RequirementId : v.Param.DataReqId,
			Username : v.Param.Info.Username,
			RequirementName : v.Param.Info.Reqname,
			ReqType:v.Param.Info.Reqtype,
			FeatureTag : v.Param.Info.Featuretag,
			SampleHash : v.Param.Info.Samplehash,
			ExpireTime : v.Param.Info.Expiretime,
			Price : v.Param.Info.Price,
			Description : v.Param.Info.Description,
			PublishDate : uint64(v.CreateTime.Unix()),
		})
	}

	var data = &requirement_proto.QueryData{
		RowCount: uint32(count),
		PageNum: uint32(pageNum),
		Row:rows,
	}
	log.Info(data)
	rsp.Data = data
	return nil
}

func init() {
	logger, err := log.LoggerFromConfigAsFile("./config/req-log.xml")
	if err != nil{
		log.Error(err)
		panic(err)
	}
	defer logger.Flush()
	log.ReplaceLogger(logger)
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.v3.requirement"),
		micro.Version("3.0.0"),
	)

	service.Init()

	requirement_proto.RegisterRequirementHandler(service.Server(), new(Requirement))

	if err := service.Run(); err != nil {
		os.Exit(1)
	}
}
