package main

import (
	"github.com/gin-gonic/gin"
	"Golang-authorization/config"
	"Golang-authorization/middleware"
	"Golang-authorization/routes"
)

func main() {
	config.ConnectDatabase()
	r := gin.Default()

	// Apply middleware globally
	r.Use(middleware.AuthMiddleware())

	routes.SetupRoutes(r)
	r.Run(":8080")
}
