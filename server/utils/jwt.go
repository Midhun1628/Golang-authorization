package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)
type Claims struct {
	Username    string   `json:"username"`
	
	jwt.RegisteredClaims
}
var secretKey = []byte(os.Getenv("SECRET_KEY"))
var refreshSecretKey=[] byte (os.Getenv("REFRESH_SECRET_KEY"))

func GenerateToken(userID uint, role string) (string,string, error) {

	accessClaims := jwt.MapClaims{
		"user_id": userID,
		"role":    role, 
		"exp":     time.Now().Add(1 * time.Minute).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	signedAccessToken,err := accessToken.SignedString(secretKey)

	if err !=nil{
		return "","",err
	}

	refreshClaims:=jwt.MapClaims{
		"user_id": userID,
		"role" :role,
		"exp":time.Now().Add(24*time.Hour).Unix(),
	
	}

	refreshToken :=jwt.NewWithClaims(jwt.SigningMethodHS256,refreshClaims)
	signedRefreshToken,err :=refreshToken.SignedString(refreshSecretKey)

	if err !=nil{
		return "","",err
	}

    return signedAccessToken,signedRefreshToken,nil

}

func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
	})
	if err != nil || !token.Valid {
			return nil, errors.New("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
			return nil, errors.New("invalid token claims")
	}
	return claims, nil
}

func ValidateRefreshToken(refreshtoken string)(jwt.MapClaims,error){
	token,err :=jwt.Parse(refreshtoken,func(token *jwt.Token)(interface{},error){
		return refreshSecretKey,nil
	})

	if err !=nil{
		return nil,errors.New("invalid refresh token")
	}

	claims, ok :=token.Claims.(jwt.MapClaims)

	if !ok{
		return nil,errors.New("invalid refresh token claims")
	}
	return claims,nil
}