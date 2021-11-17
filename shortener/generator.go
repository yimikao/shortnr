package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/itchyny/base58-go"
)

///hash long string to short with Sha256 returning []byte
///Binary-Text Encoding to provide final output with BASE58

func Sha256Of(url string) []byte {
	alg := sha256.New()
	alg.Write([]byte(url))
	return alg.Sum(nil)
}

func Base58Encode(bs []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bs)
	if err != nil {
		panic(fmt.Sprintf("couldn't encode bytes %v", err))
	}
	return string(encoded)
}

func GenerateShortLink(initialLink, userId string) string {
	urlHashBytes := Sha256Of(initialLink)
	genNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := Base58Encode([]byte(fmt.Sprintf("%d", genNumber)))
	return finalString[:8] //first 8
}
