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

	id, err := jwtProvider.Verify("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ3YXJrLmNvbSIsInN1YiI6IjAxOTI5MDliLWY2YmQtNzQ3Zi04ZGI3LTQ2OTA0ODZiZTViMiIsImV4cCI6MTcyOTA4NDAxNywiaWF0IjoxNzI5MDgzMTE3fQ.09H15XZXWqETYMUwR0NQyEvzcgaIzjZ_dobPiejsZTg")

	if err != nil {
		t.Fatal(err)
	}

	if id == "" {
		t.Fatal("jwt: expect id is present")
	}

	if id != "0192909b-f6bd-747f-8db7-4690486be5b2" {
		t.Fatal("jwt: expect id to be equal")
	}
}
