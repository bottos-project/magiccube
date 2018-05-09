package main

import (
	log "github.com/jeanphorn/log4go"
	"github.com/micro/go-micro"
	requirement_proto "github.com/bottos-project/bottos/service/requirement/proto"
	"golang.org/x/net/context"
	"github.com/bottos-project/bottos/tools/db/mongodb"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"bytes"
	"github.com/bottos-project/bottos/service/bean"
	"github.com/bottos-project/bottos/config"
)

const (
	BASE_URL              	= config.BASE_CHAIN_URL
	PUSH_TRANSACTION_URL 	= BASE_URL + "v1/chain/push_transaction"
)

type Requirement struct {}

func (u *Requirement) Publish(ctx context.Context, req *requirement_proto.PublishRequest, rsp *requirement_proto.PublishResponse) error {
	log.Info(req.Body)
	is_true := requirementPublish(req.Body)
	if is_true {
		rsp.Code = 0
		rsp.Msg = "Publish success"
	}else{
		rsp.Code = 3001
		rsp.Msg = "Publish failure"
	}
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
	if req.Username != "" && req.FeatureTag > 0 {
		where = &bson.M{"type": "datareqreg","data.basic_info.user_name": req.Username,"data.basic_info.feature_tag": req.FeatureTag}
	}else{
		if req.Username != "" {
			where = &bson.M{"type": "datareqreg","data.basic_info.user_name": req.Username}
		}
		if req.FeatureTag > 0 {
			where = &bson.M{"type": "datareqreg","data.basic_info.feature_tag": req.FeatureTag}
		}
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

	var rows = []*requirement_proto.QueryRow{}
	for _, v := range ret {

		rows = append(rows, &requirement_proto.QueryRow{
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
		RowCount: uint64(count),
		PageNum: uint64(pageNum),
		Row:rows,
	}
	log.Info(data)
	rsp.Code = 0
	rsp.Data = data
	rsp.Msg = "OK"

	return nil
}

func requirementPublish(info string) bool {
	req, err := http.NewRequest("POST", PUSH_TRANSACTION_URL, bytes.NewBuffer([]byte(info)))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode/100 == 2 {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Info("---Body---", string(body))
		return true
	}
	return false
}

func main() {
	log.LoadConfiguration(config.BASE_LOG_CONF)
	defer log.Close()
	log.LOGGER("requirement.srv")

	service := micro.NewService(
		micro.Name("go.micro.srv.requirement"),
		micro.Version("2.0.0"),
	)

	service.Init()

	requirement_proto.RegisterRequirementHandler(service.Server(), new(Requirement))

	if err := service.Run(); err != nil {
		log.Exit(err)
	}
}
