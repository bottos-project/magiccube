package signature

import (
	sign_proto "github.com/bottos-project/bottos/service/common/signature/proto"
	"github.com/golang/protobuf/proto"
	"github.com/bottos-project/crypto-go/crypto"
	"encoding/hex"
	//"github.com/ethereum/go-ethereum/common/hexutil"
	log "github.com/cihub/seelog"
	"encoding/json"
	"github.com/bottos-project/bottos/service/common/util"
	"github.com/smartwalle/errors"
)

func VerifySignBot(pubkeyStr string, jsonstr string) (bool, error) {
	var tx sign_proto.Transaction
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
	msg := &sign_proto.BasicTransaction{
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

	pub_key, err := hex.DecodeString(pubkeyStr)
	if err != nil {
		log.Error(err)
		return false, err
	}
	return crypto.VerifySign(pub_key, util.Sha256(seri_data), sign), errors.New("")
}
