/*Copyright 2017~2022 The Bottos Authors
  This file is part of the Bottos Service Layer
  Created by Developers Team of Bottos.

  This program is free software: you can distribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with Bottos. If not, see <http://www.gnu.org/licenses/>.
 */
package test

import (
	"testing"
	"encoding/hex"
	push_sign "github.com/bottos-project/magiccube/service/common/signature/push"
	"github.com/bottos-project/crypto-go/crypto"
	"github.com/bottos-project/magiccube/service/common/util"
	"github.com/protobuf/proto"
	pack "github.com/bottos-project/magiccube/core/contract/msgpack"
	"github.com/bottos-project/magiccube/service/common/bean"
	"github.com/bottos-project/magiccube/service/common/data"
)

func TestRegitser(t *testing.T) {

}


func TestSignature(t *testing.T){
	dids := bean.Did{
		"asdsada11",
		`sadsasadasd1a`,
	}

	param,_ := pack.Marshal(dids)
	t.Log(hex.EncodeToString(param))
	data := &push_sign.TransactionSign{
		Version: 1,
		CursorNum: 17,
		CursorLabel: 1798372187,
		Lifetime: 1524802582,
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

func TestGoType(t *testing.T)  {

	ret, err := data.PushTransaction(nil)
	t.Log(err)
	t.Log(ret)
}

func TestGetAccountInfo(t *testing.T)  {
	t.Log(data.AccountInfo("rrrr"))
}