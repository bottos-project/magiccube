package signature

import (
	query_sign "github.com/bottos-project/bottos/service/common/signature/query"
	push_sign "github.com/bottos-project/bottos/service/common/signature/push"
	"github.com/golang/protobuf/proto"
	"github.com/bottos-project/crypto-go/crypto"
	"encoding/hex"
	log "github.com/cihub/seelog"
	"encoding/json"
	"github.com/bottos-project/bottos/service/common/util"
	"errors"
	"github.com/bottos-project/bottos/service/common/bean"
	"github.com/bottos-project/bottos/service/common/data"
)

func PushVerifySign(jsonstr string, pubkey ...string) (bool, error) {

	var tx bean.TxBean
	err:=json.Unmarshal([]byte(jsonstr), &tx)
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
	return crypto.VerifySign(pub_key, util.Sha256(seri_data), sign), errors.New("")
}

type CommonQuery struct{
	Username string `json:"username"`
	Random string `json:"random"`
	Signature string `json:"signature"`
}

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
		Username:commonQuery.Username,
		Random:commonQuery.Random,
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
	return crypto.VerifySign(pub_key, util.Sha256(seri_data), sign), errors.New("")
}

