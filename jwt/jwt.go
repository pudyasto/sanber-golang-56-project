package jwt

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

// Define a custom claims type
type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.StandardClaims
}

func GenerateToken() (generatedToken string) {
	godotenv.Load("config/.env")
	appkey := os.Getenv("APP_KEY")
	signingKey := []byte(appkey)

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
	appkey := os.Getenv("APP_KEY")
	signingKey := []byte(appkey)
	parsedToken, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if parsedToken.Valid && err == nil {
		return true
	} else {
		return false
	}
}
