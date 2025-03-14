package controllers

import (
	"Golang-authorization/config"
	"Golang-authorization/models"
	"Golang-authorization/utils"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginResponse struct {
	Username     string `json:"username"`
	AccessToken  string `json:"access_token"`
	JobTitle     string `json:"job_title"`
	RefreshToken string `json:"refresh_token"`
}

func LoginUser(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user by email
	var user models.User
	if err := config.DB.Preload("Role").Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email ID"})
		return
	}

	// Validate password (compare directly as per your request)
	if user.Password != input.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
		return
	}

	// Generate JWT Tokens
	accessToken, refreshToken, err := utils.GenerateToken(user.ID, user.Role.EmployeePosition)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// Save Refresh Token in DB
	user.RefreshToken = refreshToken
	config.DB.Save(&user)

	// Set Refresh Token as HttpOnly Cookie (secure, not accessible by JavaScript)
	c.SetCookie("refresh_token", refreshToken, 7*24*60*60, "/", "", false, true)

	// Send JSON response to frontend with only necessary data
	c.JSON(http.StatusOK, gin.H{
		"username":     user.Username,
		"access_token": accessToken, // Store only this in localStorage
		"job_title":    user.Role.EmployeePosition,
	})
}



func RefreshToken(c *gin.Context) {

	refreshToken ,err := c.Cookie("refresh_token")

	if err !=nil{
		c.JSON(http.StatusUnauthorized,gin.H{"error":"No refresh token is stored in cookies"})
		return
	}
// validating refresh token
	claims,err :=utils.ValidateRefreshToken(refreshToken)

	if err !=nil{
		c.JSON(http.StatusUnauthorized,gin.H{"error" : "Invalid refresh token to Validate"})
	}
//extracting user ID
	userID,ok := claims["user_id"].(float64)

	if ! ok{
		c.JSON(http.StatusUnauthorized,gin.H{"error":"Invalid token claims from refresh token"})
		return
	}

	var user models.User

	if err :=config.DB.Preload("Role").Where("id= ?",uint(userID)).First(&user).Error;

	err !=nil{
		c.JSON(http.StatusUnauthorized,gin.H{"error":"User not found"})
		return
	}
fmt.Println("refresh token stored :",user.RefreshToken)
	// Check if the refresh token matches the stored token
	if user.RefreshToken != refreshToken {
		c.JSON(http.StatusUnauthorized, gin.H{"stored token":user.RefreshToken})
		return
	}

	acesssToken, newRefreshToken ,err :=utils.GenerateToken(user.ID,user.Role.EmployeePosition)

	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Could  not generate a new token"})
	}

	user.RefreshToken=newRefreshToken
	config.DB.Save(&user)

	c.SetCookie("refresh_token",newRefreshToken,7*24*60*60,"/","",false ,true)

	c.JSON(http.StatusOK,gin.H {"new Access Token": acesssToken})
}


// func getRoleName(roleID uint) string {
// 	var role models.Role
// 	if err := config.DB.Where("roll_no = ?", roleID).First(&role).Error; err != nil {
// 		return "Unknown"
// 	}
// 	return role.EmployeePosition
// }