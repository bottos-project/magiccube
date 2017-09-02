package cli

import (
	"math/rand"
	"time"

	"BOT/common/config"
	"BOT/common/log"
	"BOT/crypto"
)

func init() {
	log.Init()
	crypto.SetAlg(config.Parameters.EncryptAlg)
	//seed transaction nonce
	rand.Seed(time.Now().UnixNano())
}
