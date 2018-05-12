package signature

import (
	sign_proto "./proto"
	"github.com/golang/protobuf/proto"
	"crypto/sha256"
	"github.com/bottos-project/crypto-go/crypto"
	"encoding/hex"
	//"github.com/ethereum/go-ethereum/common/hexutil"
	log "github.com/cihub/seelog"
	"encoding/json"
)

func VerifySignBot(pubkeyStr string, jsonstr string) bool {
	var req sign_proto.Test
	json.Unmarshal([]byte(jsonstr), &req)
	log.Info(req)

	dataByte, _ := hex.DecodeString(req.Param)
	log.Info(dataByte)
	testMsg := &sign_proto.BasicTest{
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
	log.Info("testMsg:", testMsg)
	//data serialization
	data, _ := proto.Marshal(testMsg)
	log.Info("data:", data)

	//generate Hash
	h := sha256.New()
	h.Write([]byte(hex.EncodeToString(data)))
	bs := h.Sum(nil)
	log.Info("bs:", bs)

	sign, _ := hex.DecodeString(req.Signature)
	//hex string to byte[]
	pub_key, _ := hex.DecodeString(pubkeyStr)
	log.Info(pub_key)

	return crypto.VerifySign(pub_key, bs, sign)
}
