package common

import (
	"crypto/rand"
	"log"
	"math/big"
	"strings"
)

const MAX_SALT_CHAR = 12

var letters []rune = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
var maxIndex *big.Int = big.NewInt(MAX_SALT_CHAR)

func GetSalt() string {
	salt := &strings.Builder{}

	for range MAX_SALT_CHAR {
		index, err := rand.Int(rand.Reader, maxIndex)

		if err != nil {
			log.Fatalln(err)
		}

		salt.WriteRune(letters[index.Int64()])
	}

	return salt.String()
}
