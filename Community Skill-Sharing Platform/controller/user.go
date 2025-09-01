package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skillplatform/database"
	"github.com/skillplatform/model"
)

func SignUp(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to bind the json",
		})
		return
	}

	hashPassword, err := HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to generate hash passowrd",
		})
		return
	}
	user.Password = hashPassword
	database.DB.Create(&user)
	// implemend the token jwt token return , the user must see the token
	c.JSON(http.StatusOK, gin.H{
		"message": "signup successfully...",
		"user":    user,
	})
}
func Login(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to bind the json",
		})
		return
	}
	var retriveinfo model.User
	err = database.DB.Where("email=?", user.Email).First(&retriveinfo).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "email is not register",
		})
		return
	}
	if !CompareHashPassword(user.Password, retriveinfo.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid username or password",
		})
		return
	}
	token, err := GenerateJwtToken(retriveinfo.ID, retriveinfo.Name, retriveinfo.Email, retriveinfo.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to generate jwt token",
			"error":   err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "login successfully...",
		"token":   token,
	})

}
