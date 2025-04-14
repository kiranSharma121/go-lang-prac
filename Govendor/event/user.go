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
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "stored data in the database successfully...",
	})
}
