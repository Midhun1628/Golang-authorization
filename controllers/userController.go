package controllers

import (
    "github.com/gin-gonic/gin"
    "Golang-authorization/config"
    "Golang-authorization/models"
    "net/http"
)

func GetUsers(c *gin.Context) {
    var users []models.User
    config.DB.Find(&users)
    c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Create(&user)
    c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func UpdateUser(c *gin.Context) {
    var user models.User
    id := c.Param("id")
    if err := config.DB.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.ShouldBindJSON(&user)
    config.DB.Save(&user)
    c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUser(c *gin.Context) {
    id := c.Param("id")
    config.DB.Delete(&models.User{}, id)
    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
