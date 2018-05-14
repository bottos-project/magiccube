package test

import(
	"testing"
	sign "github.com/bottos-project/bottos/service/common/signature/proto"
	"github.com/bottos-project/bottos/crypto"
	"encoding/hex"
	"github.com/bottos-project/bottos/service/common/util"
	"github.com/protobuf/proto"
	pack "github.com/bottos-project/bottos/core/contract/msgpack"
	"github.com/bottos-project/bottos/service/common/bean"
)

func TestRegitser(t *testing.T) {

}


func TestSignature(t *testing.T){
	dids := bean.Did{
		"test",
		"test",
	}
	t.Log(dids)
	param,_ := pack.Marshal(dids)
	t.Log(hex.EncodeToString(param))
	data := &sign.BasicTransaction{
		Version: 1,
		CursorNum: 5134,
		CursorLabel: 1754826169,
		Lifetime: 1526281264,
		Sender: "bottos",
		Contract: "usermng",
		Method: "reguser",
		Param: param,
		SigAlg:1,
	}
	msg, _ := proto.Marshal(data)
	seckey,_ := hex.DecodeString("e4877f7665e3c22d4e5acb1a24a2fc3554ceaa575e2a3a9e794a98d9c4c3940f")
	t.Log(msg)
	t.Log(seckey)
	sign, _ := crypto.Sign(util.Sha256(msg), seckey)

	t.Log(hex.EncodeToString(sign))
}


func TestNewUser(t *testing.T){
	dids := bean.Did{
		"test",
		"test",
	}
	t.Log(dids)
	param,_ := pack.Marshal(dids)
	t.Log(param)
}


func TestGenKey(t *testing.T) {
	/**
		一
		04eb2a646ccd798646d02a0c0b17a9627bd32a0825b5393e6bbf8737090d8996ee786ee4ea2676c6e0736376b8cfa363ae985d6eb587220f773b511b8db64301ea
		e4877f7665e3c22d4e5acb1a24a2fc3554ceaa575e2a3a9e794a98d9c4c3940f

		二
		0401787e34de40f3aeb4c28259637e8c9e84b5a58f57b3c23f010f4dc7230dffced4976238196bd32cd90569d66f747525b194ca83146965df092d2585b975d0d3
		81407d25285450184d29247b5f06408a763f3057cba6db467ff999710aeecf8e
	*/
	pubkey, seckey := crypto.GenerateKey()

	t.Log(hex.EncodeToString(pubkey))
	t.Log(hex.EncodeToString(seckey))
}