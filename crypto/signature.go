package crypto

import(
	"crypto/ecdsa"
	"github.com/bottos-project/bottos/crypto/secp256k1"
	"crypto/rand"
	"crypto/elliptic"
	"math/big"
)

const (

	// number of bits in a big.Word
	wordBits = 32 << (uint64(^big.Word(0)) >> 63)
	// number of bytes in a big.Word
	wordBytes = wordBits / 8
)


func GenerateKey() (pubkey, seckey []byte) {
	key, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	return elliptic.Marshal(secp256k1.S256(), key.X, key.Y), PaddedBigBytes(key.D, 32)
}

func Sign(msg, seckey []byte) ([]byte, error){
	sign, err := secp256k1.Sign(msg, seckey)
	return sign[:len(sign)-1], err
}

func VerifySign(pubkey, msg, sign []byte) bool {
	return secp256k1.VerifySignature(pubkey, msg, sign)
}

func PaddedBigBytes(bigint *big.Int, n int) []byte {
	if bigint.BitLen()/8 >= n {
		return bigint.Bytes()
	}
	ret := make([]byte, n)
	ReadBits(bigint, ret)
	return ret
}

func ReadBits(bigint *big.Int, buf []byte) {
	i := len(buf)
	for _, d := range bigint.Bits() {
		for j := 0; j < wordBytes && i > 0; j++ {
			i--
			buf[i] = byte(d)
			d >>= 8
		}
	}
}
