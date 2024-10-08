package jwt

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtProvider struct {
	secret string
}

var (
	SECRET_KEY = os.Getenv("SECRET_KEY")
)

func New(secret string) *jwtProvider { return &jwtProvider{secret} }

func (*jwtProvider) Generate(id string, expiry time.Duration) string {
	registerdClaims := jwt.RegisteredClaims{
		Issuer:    "wark.com",
		Subject:   id,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
		IssuedAt:  &jwt.NumericDate{Time: time.Now()},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, registerdClaims)
	tokenString, err := token.SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Fatal(err)
	}

	return tokenString
}

// func (*jwtProvider) Verify(tokenString string) (string, error) {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		return jwtProvider.secret, nil
// 	})

// 	switch {
// 	case token.Valid:
// 		claims, ok := token.Claims.(jwt.RegisteredClaims)

// 		if !ok {
// 			log.Fatalln("Claims is not of type jwt.RegisteredClaims")
// 		}

// 		return claims.Subject, nil
// 	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
// 		log.Fatalln("Token is not expire")
// 	}

// 	return "", nil
// }
