package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tutorplatform/database"
	"github.com/tutorplatform/model"
)

func Signup(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "couldn't bind the json",
			"error":   err.Error(),
		})
		return
	}
	hashpassword, err := Hashpassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"messsage": "unable to hased the password",
			"error":    err.Error(),
		})
		return
	}
	user.Password = hashpassword
	database.DB.Create(&user)
	c.JSON(http.StatusOK, user)
}
func Login(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to bind the json",
			"error":   err.Error(),
		})
		return
	}
	var retriveinfo model.User
	err = database.DB.Where("email = ?", user.Email).First(&retriveinfo).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email not registered"})
		return
	}
	if !Comparepassword(user.Password, retriveinfo.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid username or password",
		})
		return
	}
token, err := Generatejwttoken(retriveinfo.Id, retriveinfo.Name, retriveinfo.Email, retriveinfo.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to generate jwt token",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "login successfully",
		"token":   token,
	})
}
func GetUser(c *gin.Context) {
	var user []model.User
	database.DB.Find(&user)
	c.JSON(http.StatusOK, user)
}
