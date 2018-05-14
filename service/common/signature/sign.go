package signature

import (
	sign_proto "github.com/bottos-project/bottos/service/common/signature/proto"
	"github.com/golang/protobuf/proto"
	"crypto/sha256"
	"github.com/bottos-project/crypto-go/crypto"
	"encoding/hex"
	//"github.com/ethereum/go-ethereum/common/hexutil"
	log "github.com/cihub/seelog"
	"encoding/json"
)

func VerifySignBot(pubkeyStr string, jsonstr string) (bool, error) {
	var req sign_proto.Transaction
	json.Unmarshal([]byte(jsonstr), &req)
	log.Info(req)

	dataByte, err := hex.DecodeString(req.Param)
	log.Info(dataByte)
	if err != nil {
		return false, err
	}
	msg := &sign_proto.BasicTransaction{
		Version:     req.Version,
		CursorNum:   req.CursorNum,
		CursorLabel: req.CursorLabel,
		Lifetime:    req.Lifetime,
		Sender:      req.Sender,
		Contract:    req.Contract,
		Method:      req.Method,
		Param:       dataByte,
		SigAlg:      req.SigAlg,
	}
	log.Info("testMsg:", msg)
	//data serialization
	data, err := proto.Marshal(msg)
	log.Info("data:", data)
	if err != nil {
		return false, err
	}

	//generate Hash
	h := sha256.New()
	h.Write([]byte(hex.EncodeToString(data)))
	bs := h.Sum(nil)
	log.Info("bs:", bs)

	sign, err := hex.DecodeString(req.Signature)
	if err != nil {
		return false, err
	}
	//hex string to byte[]
	pub_key, err := hex.DecodeString(pubkeyStr)
	log.Info(pub_key)
	if err != nil {
		return false, err
	}

	return crypto.VerifySign(pub_key, bs, sign), err
}
