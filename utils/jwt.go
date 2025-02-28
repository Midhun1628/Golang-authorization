package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Custom claims struct
type Claims struct {
	Username    string   `json:"username"`

	jwt.RegisteredClaims
}

// GenerateJWT generates a new token for a given username, role, and permissions



var secretKey = []byte(os.Getenv("SECRET_KEY"))

func GenerateToken(userID uint) (string, error) {
    expirationTime := time.Now().Add(1 * time.Minute)
    claims := &jwt.MapClaims{
        "user_id": userID,
        "exp":     expirationTime.Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(secretKey)
}


// ValidateJWT verifies the given token
func ValidateJWT(tokenString string) (*Claims, error) {
token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		return []byte(secretKey), nil
	})

	fmt.Println("token from validate func :", token)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)

fmt.Println("claims from validate func:", claims)

	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}