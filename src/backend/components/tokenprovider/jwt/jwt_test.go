package jwt

import (
	"testing"
	"time"
)

func TestGenerate(t *testing.T) {
	jwtProvider := New("ABCD")
	token, err := jwtProvider.Generate("asdasdasdasdasdad", time.Second*15)

	if token == "" || err != nil {
		t.Fail()
		t.Logf("Cannot generate token with %v", token)
	}
}
