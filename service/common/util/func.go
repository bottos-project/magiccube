package util

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(msg []byte) []byte {
	sha := sha256.New()
	sha.Write([]byte(hex.EncodeToString(msg)))
	return sha.Sum(nil)
}
