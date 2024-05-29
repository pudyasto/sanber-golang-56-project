package jwt

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Define a custom claims type
type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.StandardClaims
}

func GenerateToken() (generatedToken string) {
	signingKey := []byte("Ed1bBrCKct06xSi2ii6wR8B3IoXbqnvg")

	claims := MyCustomClaims{
		"bar",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		log.Fatalf("Error signing the token: %v", err)
	}
	return tokenString

}

func CheckToken(tokenString string) bool {
	signingKey := []byte("Ed1bBrCKct06xSi2ii6wR8B3IoXbqnvg")
	parsedToken, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if parsedToken.Valid && err == nil {
		return true
	} else {
		return false
	}
}
