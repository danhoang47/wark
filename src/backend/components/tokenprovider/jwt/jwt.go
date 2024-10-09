package jwt

import (
	"errors"
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
		return "", fmt.Errorf("cannot generate token for id: %s", id)
	}

	return tokenString, nil
}

func (j *jwtProvider) Verify(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	switch {
	case token.Valid:
		claims, ok := token.Claims.(jwt.RegisteredClaims)

		if !ok {
			log.Fatalln("Claims is not of type jwt.RegisteredClaims")
		}

		return claims.Subject, nil
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		panic(jwt.ErrTokenExpired)
	}

	return "", nil
}
