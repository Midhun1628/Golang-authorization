package routes

import (
	"Golang-authorization/controllers"
	"Golang-authorization/middleware"

	
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/login", controllers.LoginUser)
	router.POST("/refresh",controllers.RefreshToken)
	router.GET("/roles", controllers.GetRoles)
	
	auth := router.Group("/")
	auth.Use(middleware.AuthMiddleware()) // Protect all routes with authentication

	{
		auth.POST("/register", middleware.PermissionMiddleware("Create"), controllers.CreateUser)
		auth.GET("/users", middleware.PermissionMiddleware("Read"), controllers.GetUsers)
		auth.PUT("/users/:id", middleware.PermissionMiddleware("Update"), controllers.UpdateUser)
		auth.DELETE("/users/:id", middleware.PermissionMiddleware("Delete"), controllers.DeleteUser)
	}
}