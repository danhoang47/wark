package hasher

import "encoding/base64"

type Hasher interface {
	Hash([]byte) string
}

type hasher struct{}

func New() Hasher { return &hasher{} }

func (*hasher) Hash(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
