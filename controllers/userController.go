package controllers

import (
	"net/http"
	"Golang-authorization/config"
	"Golang-authorization/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Preload("Role").Find(&users)
	c.JSON(http.StatusOK, users)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	config.DB.First(&user, id)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.User{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}	