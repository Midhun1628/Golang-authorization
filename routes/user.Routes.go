package routes

import (
	"github.com/gin-gonic/gin"
	"Golang-authorization/controllers"
	"Golang-authorization/middleware"
)

func SetupRoutes(router *gin.Engine) {
	// Auth Route
	router.POST("/login", controllers.Login)

	userGroup := router.Group("/users")
	userGroup.Use(middleware.AuthMiddleware()) // Apply JWT middleware globally
	{
		userGroup.GET("/", controllers.GetUsers)
		userGroup.POST("/", middleware.RoleMiddleware("Super Admin", "Admin"), controllers.CreateUser)
		userGroup.PUT("/:id", middleware.RoleMiddleware("Super Admin", "Admin"), controllers.UpdateUser)
		userGroup.DELETE("/:id", middleware.RoleMiddleware("Super Admin"), controllers.DeleteUser)
	}
}



