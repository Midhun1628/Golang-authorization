package services

import (
	"os"
	"time"
"github.com/golang-jwt/jwt/v5"	

)

// GenerateJWT generates a new token for a given username and role
func GenerateJWT(username string, role string) (string, error) {
	secretKey := os.Getenv("SECRET_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	return token.SignedString([]byte(secretKey))
}

// ValidateJWT verifies the given token
func ValidateJWT(tokenString string) (*jwt.Token, error) {
	secretKey := os.Getenv("SECRET_KEY")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	return token, err
}
