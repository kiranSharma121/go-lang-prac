package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/movie/models"
	"github.com/movie/utils"
)

func Signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to bind the json",
			"error":   err.Error(),
		})
		return
	}
	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "couldn't able to save user in the database",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "sucessfully stored user in the database",
	})

}
func Login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to bind the json",
			"error":   err.Error(),
		})
		return
	}
	err = user.Validatecredentials()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to validate user",
			"error":   err.Error(),
		})
		return
	}
	token, err := utils.GeneratejwtToken(user.Userid, user.UserName, user.Email, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error in generating the token",
			"error":   err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"token":   token,
	})
}
