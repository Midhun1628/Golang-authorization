
package middlewares

import (
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "net/http"
    
)


func RoleMiddleware(requiredRole string) gin.HandlerFunc {
    return func(c *gin.Context) {
        claims, exists := c.Get("user")
        if !exists {
            c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        userRole := claims.(jwt.MapClaims)["role"].(string)
        if userRole != requiredRole && userRole != "Super Admin" {
            c.JSON(http.StatusForbidden, gin.H{"error": "Access Denied"})
            c.Abort()
            return
        }
        c.Next()
    }
}
