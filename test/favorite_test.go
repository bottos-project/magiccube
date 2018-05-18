package test

import (
	"testing"
	"encoding/hex"
	push_sign "github.com/bottos-project/bottos/service/common/signature/push"
	"github.com/bottos-project/bottos/service/common/util"
	"github.com/protobuf/proto"
	"github.com/bottos-project/crypto-go/crypto"
	pack "github.com/bottos-project/bottos/core/contract/msgpack"
)

type Favorite struct {
	Username string
	GoodsId string
	GoodsType string //[asset, requirement]
	OpType uint32

}

func TestFavoriteSignature(t *testing.T){
	data := Favorite{
		Username:"ttsss1s1s",
		GoodsId:"1",
		GoodsType:"asset",
		OpType:1,
	}

	param,_ := pack.Marshal(data)
	t.Log(hex.EncodeToString(param))
	data1 := &push_sign.TransactionSign{
		Version: 1,
		CursorNum: 17,
		CursorLabel: 1798372187,
		Lifetime: 1524802582,
		Sender: "ttsss1s1s",
		Contract: "favoritemng",
		Method: "favoritepro",
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