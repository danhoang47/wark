package jwt

import (
	"testing"
	"time"
)

func TestGenerate(t *testing.T) {
	jwtProvider := New("ABCD")
	token := jwtProvider.Generate("asdasdasdasdasdad", time.Second*15)

	if token == "" {
		t.Fail()
		t.Logf("Cannot generate token with %v", token)
	}
}
