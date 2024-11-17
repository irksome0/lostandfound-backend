package utils

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func getSecret() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env file!")
	}

	SecretKey := os.Getenv("SECRET_KEY")
	return SecretKey
}

func GenerateJwt(issuer string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	return claims.SignedString([]byte(getSecret()))
}

func ParseJwt(cookie string) (string, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(getSecret()), nil
	})
	if err != nil || !token.Valid {
		return "", err
	}
	claims := token.Claims.(*jwt.StandardClaims)
	return claims.Issuer, nil
}

func ConvertTime(value string) (time.Time, error) {
	layout := time.RFC3339
	parsedTime, err := time.Parse(layout, value)
	if err != nil {
		return time.Now(), err
	}
	return parsedTime, nil
}
