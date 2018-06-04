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
	query_sign "github.com/bottos-project/magiccube/service/common/signature/query"
	"github.com/bottos-project/magiccube/service/common/util"
	"github.com/protobuf/proto"
	"github.com/bottos-project/crypto-go/crypto"
	pack "github.com/bottos-project/magiccube/core/contract/msgpack"
)

type Favorite struct {
	Username string
	OpType uint32
	GoodsType string //[asset, requirement]
	GoodsId string
}

func TestFavoriteSignature(t *testing.T){
	//dc0004da000474657374ce00000001da000b726571756972656d656e74da000131
	//dc0004da000474657374ce00000001da000b726571756972656d656e74da000131
	data := Favorite{
		Username:"test",
		OpType:1,
		GoodsId:"1",
		GoodsType:"requirement",
	}

	param,_ := pack.Marshal(data)
	t.Log(hex.EncodeToString(param))
	data1 := &push_sign.TransactionSign{
		Version: 1,
		CursorNum: 17,
		CursorLabel: 1798372187,
		Lifetime: 1524802582,
		Sender: "test",
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

func TestGetFavoriteSignature(t *testing.T){
	data := &query_sign.QuerySign{
		Username: "test",
		Random:"1",
	}


	msg, _ := proto.Marshal(data)
	seckey,_ := hex.DecodeString("e4877f7665e3c22d4e5acb1a24a2fc3554ceaa575e2a3a9e794a98d9c4c3940f")

	t.Log(msg)
	t.Log(seckey)
	sign, _ := crypto.Sign(util.Sha256(msg), seckey)

	t.Log(hex.EncodeToString(sign))
}