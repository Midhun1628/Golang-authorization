package routes

import (
    "github.com/gin-gonic/gin"
   "Golang-authorization/controllers"
    "Golang-authorization/middleware"
)

func SetupRoutes(router *gin.Engine) {
    userGroup := router.Group("/users")
    {
        userGroup.POST("/", middlewares.AuthMiddleware(), middlewares.RoleMiddleware("Admin"), controllers.CreateUser)
        userGroup.PUT("/:id", middlewares.AuthMiddleware(), middlewares.RoleMiddleware("Admin"), controllers.UpdateUser)
        userGroup.DELETE("/:id", middlewares.AuthMiddleware(), middlewares.RoleMiddleware("Super Admin"), controllers.DeleteUser)
        userGroup.GET("/", middlewares.AuthMiddleware(), controllers.GetUsers)
    }
}
