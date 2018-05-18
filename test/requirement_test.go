package test

import (
	"testing"
	"encoding/hex"
	push_sign "github.com/bottos-project/bottos/service/common/signature/push"
	"github.com/bottos-project/bottos/crypto"
	"github.com/bottos-project/bottos/service/common/util"
	"github.com/protobuf/proto"
	pack "github.com/bottos-project/bottos/core/contract/msgpack"
)

type Requirement struct {
	RequirementId   string `protobuf:"bytes,2,opt,name=requirement_id,json=requirementId" json:"requirement_id"`
	RequirementData RequirementData
}

type RequirementData struct {
	Username        string `protobuf:"bytes,1,opt,name=username" json:"username"`
	RequirementName string `protobuf:"bytes,3,opt,name=requirement_name,json=requirementName" json:"requirement_name"`
	FeatureTag      uint64 `protobuf:"varint,4,opt,name=feature_tag,json=featureTag" json:"feature_tag"`
	SamplePath      string `protobuf:"bytes,5,opt,name=sample_path,json=samplePath" json:"sample_path"`
	SampleHash      string `protobuf:"bytes,6,opt,name=sample_hash,json=sampleHash" json:"sample_hash"`
	ExpireTime      uint32 `protobuf:"varint,7,opt,name=expire_time,json=expireTime" json:"expire_time"`
	Price           uint64 `protobuf:"varint,8,opt,name=price" json:"price"`
	Description     string `protobuf:"bytes,9,opt,name=description" json:"description"`
	PublishDate     uint32 `protobuf:"varint,10,opt,name=publish_date,json=publishDate" json:"publish_date"`
}

func TestReqSignature(t *testing.T){
	data := Requirement{
		RequirementId:"1",
		RequirementData: RequirementData{
			Username:"tttt",
			RequirementName: "1111",
			FeatureTag:1,
			SamplePath: "./s/asd/sd.png",
			SampleHash:"asdasdsdagkfdjg3",
			ExpireTime:1455379533,
			Price:1000,
			Description:"test",
			PublishDate:1455379533,
		},
	}

	param,_ := pack.Marshal(data)
	t.Log(hex.EncodeToString(param))
	data1 := &push_sign.TransactionSign{
		Version: 1,
		CursorNum: 17,
		CursorLabel: 1798372187,
		Lifetime: 1524802582,
		Sender: "tttt",
		Contract: "datareqmng",
		Method: "datareqreg",
		Param: param,
		SigAlg:1,
	}
	msg, _ := proto.Marshal(data1)
	seckey,_ := hex.DecodeString("e4877f7665e3c22d4e5acb1a24a2fc3554ceaa575e2a3a9e794a98d9c4c3940f")

	t.Log(msg)
	t.Log(seckey)
	sign, _ := crypto.Sign(util.Sha256(msg), seckey)

	t.Log(hex.EncodeToString(sign))
}