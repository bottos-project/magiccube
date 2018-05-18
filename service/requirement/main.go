package main

import (
	log "github.com/cihub/seelog"
	"github.com/micro/go-micro"
	requirement_proto "github.com/bottos-project/bottos/service/requirement/proto"
	"golang.org/x/net/context"
	"github.com/bottos-project/bottos/tools/db/mongodb"
	"gopkg.in/mgo.v2/bson"
	"github.com/bottos-project/bottos/service/bean"
	"github.com/bottos-project/bottos/config"
	"os"
	"github.com/bottos-project/bottos/service/common/data"
)


type Requirement struct {}

func (u *Requirement) Publish(ctx context.Context, req *requirement_proto.PublishRequest, rsp *requirement_proto.PublishResponse) error {
	i, err := data.PushTransaction(req)
	if err != nil {
		rsp.Code = 3000
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

	if req.PageSize > 0 {
		pageSize = int(req.PageSize)
	}

	skip = (pageNum - 1) *  pageSize

	var where interface{}
	where = &bson.M{"type": "datareqreg"}
	log.Info(req.Username)
	if req.Username != ""{
		where = &bson.M{"type": "datareqreg","data.basic_info.user_name": req.Username}
	}

	log.Info(where)

	var ret []bean.RequirementBean

	var mgo = mgo.Session()
	defer mgo.Close()
	count, err:= mgo.DB(config.DB_NAME).C("Messages").Find(where).Count()
	log.Info(count)
	if err != nil {
		log.Error(err)
	}
	mgo.DB(config.DB_NAME).C("Messages").Find(where).Sort("-_id").Skip(skip).Limit(pageSize).All(&ret)

	var rows = []*requirement_proto.RequirementData{}
	for _, v := range ret {

		rows = append(rows, &requirement_proto.RequirementData{
			RequirementId : v.Data.DataReqID,
			Username : v.Data.BasicInfo.UserName,
			RequirementName : v.Data.BasicInfo.RequirementName,
			FeatureTag : v.Data.BasicInfo.FeatureTag,
			SamplePath : v.Data.BasicInfo.SamplePath,
			SampleHash : v.Data.BasicInfo.SampleHash,
			ExpireTime : v.Data.BasicInfo.ExpireTime,
			Price : v.Data.BasicInfo.Price,
			Description : v.Data.BasicInfo.Description,
			PublishDate : uint32(v.CreatedAt.Unix()),

		})
	}

	var data = &requirement_proto.QueryData{
		RowCount: uint32(count),
		PageNum: uint32(pageNum),
		Row:rows,
	}
	log.Info(data)
	rsp.Code = 0
	rsp.Data = data
	rsp.Msg = "OK"

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
		micro.Name("bottos.srv.requirement"),
		micro.Version("3.0.0"),
	)

	service.Init()

	requirement_proto.RegisterRequirementHandler(service.Server(), new(Requirement))

	if err := service.Run(); err != nil {
		os.Exit(1)
	}
}
