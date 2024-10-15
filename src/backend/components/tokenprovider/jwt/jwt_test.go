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

func TestVerify(t *testing.T) {
	jwtProvider := New("wark")

	id, err := jwtProvider.Verify("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ3YXJrLmNvbSIsInN1YiI6IjAxOTI5MDliLWY2YmQtNzQ3Zi04ZGI3LTQ2OTA0ODZiZTViMiIsImV4cCI6MTcyOTAwNTQ2MiwiaWF0IjoxNzI5MDA0NTYyfQ.f1NdiyKxBoR1njtmhVlXNfbub6nvH2rZq_GmP8YbdX8")

	if err != nil {
		t.Fatal(err)
	}

	if id == "" {
		t.Fatal("jwt: expect id is present")
	}
}
