package routes

import (
	"Golang-authorization/controllers"
	"Golang-authorization/middleware"

	
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/login", controllers.LoginUser)
	
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		    protected.POST("/register", controllers.CreateUser)
			protected.GET("/users", controllers.GetUsers)
			protected.PUT("/users/:id", controllers.UpdateUser)
			protected.DELETE("/users/:id", controllers.DeleteUser)
	}
	
}