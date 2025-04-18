package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/movie/models"
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
