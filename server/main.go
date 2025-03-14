package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
	"Golang-authorization/routes"
"Golang-authorization/config"
)

func main() {

	config.ConnectDatabase()
	r := gin.Default()

	// CORS Middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Allow your Vue app
		AllowMethods:     []string{"GET", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true, // Allow cookies and Authorization headers
		MaxAge:           12 * time.Hour,
	}))

	// Routes

	routes.SetupRoutes(r)

	// Start the server
	r.Run(":3000")
}
