package jwt

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtProvider struct {
	secret string
}

func New(secret string) *jwtProvider { return &jwtProvider{secret} }

func (j *jwtProvider) Generate(id string, expiry time.Duration) (string, error) {
	if j.secret == "" {
		log.Fatalln("No SECRET_KEY provided")
	}

	registerdClaims := jwt.RegisteredClaims{
		Issuer:    "wark.com",
		Subject:   id,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
		IssuedAt:  &jwt.NumericDate{Time: time.Now()},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, registerdClaims)
	tokenString, err := token.SignedString([]byte(j.secret))

	if err != nil {
		return "", fmt.Errorf("%w, cannot generate token for id: %s", err, id)
	}

	return tokenString, nil
}

func (j *jwtProvider) Verify(tokenString string) (string, error) {
	claims := &jwt.RegisteredClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.secret), nil
		})

	if token.Valid {
		return claims.Subject, nil
	}

	return "", err
}
