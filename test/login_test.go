package test

import (
	"testing"
	"encoding/hex"
	query_sign "github.com/bottos-project/magiccube/service/common/signature/query"
	"github.com/bottos-project/crypto-go/crypto"
	"github.com/bottos-project/magiccube/service/common/util"
	"github.com/protobuf/proto"
	"github.com/bottos-project/magiccube/service/common/data"
)



func TestLoginSignature(t *testing.T){
	data := &query_sign.QuerySign{
		Username:"tssd1111",
		Random:"11111111",
	}


	msg, _ := proto.Marshal(data)
	t.Log(hex.EncodeToString(msg))
	seckey,_ := hex.DecodeString("e4877f7665e3c22d4e5acb1a24a2fc3554ceaa575e2a3a9e794a98d9c4c3940f")

	t.Log(hex.EncodeToString(util.Sha256(msg)))
	sign, _ := crypto.Sign(util.Sha256(msg), seckey)

	t.Log(hex.EncodeToString(sign))
}

func TestQueryUser(t *testing.T){
	d, _:= data.AccountInfo("tssd11")
	t.Log(d)
}
