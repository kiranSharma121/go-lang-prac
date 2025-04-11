package event

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sqlite/models"
)

func SignUp(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to bind the json",
		})
		return
	}
	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "couldn't authenticate user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Data stored in the database successfully",
	})

}
func Login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to bind the data",
		})
	}
	err = user.ValidateCredentials()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid Credentials",
		})
		return
	}
	token, err := models.GenerateJwtToken(user.Email, user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "couldn't authenticate the user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Login sucessfully", "token": token,
	})

}
