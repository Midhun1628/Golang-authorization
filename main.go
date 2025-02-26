package main

import (
    "github.com/gin-gonic/gin"
    "Golang-authorization/config"
    "Golang-authorization/routes"
)

func main() {
    config.ConnectDatabase()
    r := gin.Default()
    routes.SetupRoutes(r)
    r.Run(":8080")
}
