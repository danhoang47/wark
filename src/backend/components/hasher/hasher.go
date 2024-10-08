package hasher

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
)

type sha256Hash struct{}

func NewSHA256Hash() *sha256Hash { return new(sha256Hash) }

func (*sha256Hash) Hash(data string) string {
	hasher := sha256.New()

	_, err := hasher.Write([]byte(data))

	if err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(hasher.Sum(nil))
}
