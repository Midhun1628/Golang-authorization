package routes

import (
	"Golang-authorization/controllers"
	
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/register", controllers.CreateUser)
	router.POST("/login", controllers.LoginUser)

	// Apply AuthMiddleware to protected routes
	// protected := router.Group("/")
	// protected.Use(middleware.AuthMiddleware()) // Only authenticated users can access these
	// {
		router.GET("/users", controllers.GetUsers)
		router.PUT("/users/:id", controllers.UpdateUser)
		router.DELETE("/users/:id", controllers.DeleteUser)
	// }
}
