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
package signature

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/bottos-project/crypto-go/crypto"
	"github.com/bottos-project/magiccube/service/common/bean"
	"github.com/bottos-project/magiccube/service/common/data"
	push_sign "github.com/bottos-project/magiccube/service/common/signature/push"
	query_sign "github.com/bottos-project/magiccube/service/common/signature/query"
	"github.com/bottos-project/magiccube/service/common/util"
	log "github.com/cihub/seelog"
	"github.com/golang/protobuf/proto"
)

// push verify sign
func PushVerifySign(jsonstr string, pubkey ...string) (bool, error) {

	var tx bean.TxBean
	err := json.Unmarshal([]byte(jsonstr), &tx)
	if err != nil {
		log.Error(err)
		return false, err
	}

	log.Info(tx)
	dataByte, err := hex.DecodeString(tx.Param)
	if err != nil {
		log.Error(err)
		return false, err
	}
	msg := &push_sign.TransactionSign{
		Version:     tx.Version,
		CursorNum:   tx.CursorNum,
		CursorLabel: tx.CursorLabel,
		Lifetime:    tx.Lifetime,
		Sender:      tx.Sender,
		Contract:    tx.Contract,
		Method:      tx.Method,
		Param:       dataByte,
		SigAlg:      tx.SigAlg,
	}

	//data serialization
	seri_data, err := proto.Marshal(msg)
	if err != nil {
		log.Error(err)
		return false, err
	}

	sign, err := hex.DecodeString(tx.Signature)
	if err != nil {
		log.Error(err)
		return false, err
	}

	var p_key = ""
	if len(pubkey) < 1 {
		accountInfo, err := data.AccountInfo(tx.Sender)
		if err != nil {
			log.Error(err)
			return false, err
		}
		p_key = accountInfo.Pubkey
	} else {
		p_key = pubkey[0]
	}

	pub_key, err := hex.DecodeString(p_key)
	if err != nil {
		log.Error(err)
		return false, err
	}
	return crypto.VerifySign(pub_key, util.Sha256(seri_data), sign), nil
}

type CommonQuery struct {
	Username  string `json:"username"`
	Random    string `json:"random"`
	Signature string `json:"signature"`
}

// query verify sign
func QueryVerifySign(b string) (bool, error) {

	var commonQuery CommonQuery
	err := json.Unmarshal([]byte(b), &commonQuery)
	if err != nil {
		log.Error(err)
		return false, err
	}
	log.Info(commonQuery)

	if len(commonQuery.Username) < 1 {
		return false, errors.New("The Username can not be empty")
	}

	if len(commonQuery.Signature) < 1 {
		return false, errors.New("The Signature value can not be empty")
	}

	if len(commonQuery.Random) < 1 {
		return false, errors.New("The Random value can not be empty")
	}

	msg := &query_sign.QuerySign{
		Username: commonQuery.Username,
		Random:   commonQuery.Random,
	}
	//data serialization
	seri_data, err := proto.Marshal(msg)
	if err != nil {
		log.Error(err)
		return false, err
	}

	sign, err := hex.DecodeString(commonQuery.Signature)
	if err != nil {
		log.Error(err)
		return false, err
	}

	accountInfo, err := data.AccountInfo(commonQuery.Username)
	if err != nil {
		log.Error(err)
		return false, err
	}
	log.Info("accountInfo:", accountInfo)
	pub_key, err := hex.DecodeString(accountInfo.Pubkey)
	if err != nil {
		log.Error(err)
		return false, err
	}
	return crypto.VerifySign(pub_key, util.Sha256(seri_data), sign), nil
}
