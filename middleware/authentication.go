package middlewares

import (
    "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
    "github.com/golang-jwt/jwt/v5"
    "net/http"
	"os"
    "strings"
)

func init(){
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	godotenv.Load()
}

var SecretKey =os.Getenv("JWT_SECRET")

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenHeader := c.GetHeader("Authorization")
        if tokenHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
            c.Abort()
            return
        }

        tokenString := strings.Split(tokenHeader, "Bearer ")[1]
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte(SecretKey), nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Set("user", token.Claims.(jwt.MapClaims)) // Store claims in context
        c.Next()
    }
}
