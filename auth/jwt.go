package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecretKey = []byte("your-secret-key")

func GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"user_name": username,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})
}
