package event

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goVendor/models"
)

func SignUp(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "unable to bind json",
		})
		return
	}
	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to store data in the database",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "stored data in the database successfully...",
	})
}
func Login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to bind json..!!invalid user fields",
		})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "unable to validate user credentials",
			"error":   err.Error(),
		})
		return
	}
	token, err := models.GenerateJwtToken(user.Id, user.UserName, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "couldn't authenticate the user",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful....",
		"Token":   token,
	})

}
