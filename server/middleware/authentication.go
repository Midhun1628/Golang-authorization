package middleware

import (
	"Golang-authorization/config"
	"Golang-authorization/models"
	"Golang-authorization/utils"
	
	"net/http"
	"strings"
	"github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
)
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader("Authorization")
		if tokenHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token msg from authMiddleware"})
			c.Abort()
			return
		}

		// Extract Bearer token
		parts := strings.Split(tokenHeader," ")


		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format msg from authMiddleware"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		// fmt.Println("Token received:", tokenString) // Debugging

		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
				c.Abort()
				return
		}
		c.Set("user", claims)
		

		// Store user data in context
		c.Set("user", claims)
		c.Next()
	}
}


func PermissionMiddleware(requiredPermission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token missing or invalid"})
			c.Abort()
			return
		}

		userClaims, ok := claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format from permissionMiddleware"})
			c.Abort()
			return
		}

		userID := uint(userClaims["user_id"].(float64))

		// Fetch user from DB along with their role and permissions
		var user models.User
		if err := config.DB.Preload("Role.Permissions").First(&user, userID).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		// Check if user has the required permission
		for _, permission := range user.Role.Permissions {
			if permission.Permissions == requiredPermission {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Not Authorized"})
		c.Abort()
	}
}