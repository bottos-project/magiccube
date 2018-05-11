package test

import(
	"testing"
	"github.com/bottos-project/bottos/service/common/data"
	"github.com/bottos-project/bottos/service/user/proto"
	"github.com/bottos-project/bottos/crypto"
	"encoding/hex"
)

func TestRegitser(t *testing.T) {

}


func TestS(t *testing.T){
	t.Log(data.PushTransaction(user.UserInfo{
		Version: 1,
		CursorNum: 28,
		CursorLabel: 3745260307,
		Lifetime: 1524802615,
		Sender: "bottos",
		Contract: "bottos",
		Method: "newaccount",
		Param: "",
		SigAlg:1,
		Signature:"20d8a374aee4099acd776dd59dc54d29effc2f697d4ba70562fd0f1726a31ac5",
	}))
}


func TestGenKey(t *testing.T) {
	/**
		04eb2a646ccd798646d02a0c0b17a9627bd32a0825b5393e6bbf8737090d8996ee786ee4ea2676c6e0736376b8cfa363ae985d6eb587220f773b511b8db64301ea
		e4877f7665e3c22d4e5acb1a24a2fc3554ceaa575e2a3a9e794a98d9c4c3940f
	*/
	pubkey, seckey := crypto.GenerateKey()

	t.Log(hex.EncodeToString(pubkey))
	t.Log(hex.EncodeToString(seckey))
}