package test

import (
	"testing"
	"encoding/hex"
	push_sign "github.com/bottos-project/magiccube/service/common/signature/push"
	"github.com/bottos-project/magiccube/service/common/util"
	"github.com/protobuf/proto"
	"github.com/bottos-project/crypto-go/crypto"
	pack "github.com/bottos-project/magiccube/core/contract/msgpack"
)

type Requirement struct {
	RequirementId   string `protobuf:"bytes,2,opt,name=requirement_id,json=requirementId" json:"requirement_id"`
	RequirementData RequirementData
}

type RequirementData struct {
	Username        string `protobuf:"bytes,1,opt,name=username" json:"username"`
	RequirementName string `protobuf:"bytes,3,opt,name=requirement_name,json=requirementName" json:"requirement_name"`
	RequirementType uint64
	FeatureTag      uint64 `protobuf:"varint,4,opt,name=feature_tag,json=featureTag" json:"feature_tag"`
	SampleHash      string `protobuf:"bytes,6,opt,name=sample_hash,json=sampleHash" json:"sample_hash"`
	ExpireTime      uint64 `protobuf:"varint,7,opt,name=expire_time,json=expireTime" json:"expire_time"`
	OpType 			uint32
	Price           uint64 `protobuf:"varint,8,opt,name=price" json:"price"`
	FavoriFlag      uint32
	Description     string `protobuf:"bytes,9,opt,name=description" json:"description"`
}

func TestReqSignature(t *testing.T){
	data := Requirement{
		RequirementId:"2",
		RequirementData: RequirementData{
			Username:"test",
			RequirementName: "1",
			RequirementType: 1,
			FeatureTag:1,
			SampleHash:"1",
			ExpireTime:1455379533,
			Price:1000,
			Description:"1",
			FavoriFlag:1,
			OpType:2,
		},
	}

	param,_ := pack.Marshal(data)
	t.Log(hex.EncodeToString(param))
	data1 := &push_sign.TransactionSign{
		Version: 1,
		CursorNum: 17,
		CursorLabel: 1798372187,
		Lifetime: 1524802582,
		Sender: "test",
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