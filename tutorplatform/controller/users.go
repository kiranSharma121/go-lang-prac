package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tutorplatform/database"
	"github.com/tutorplatform/model"
)

func CreateUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "couldn't bind the json",
			"error":   err.Error(),
		})
		return
	}
	database.DB.Create(&user)
	c.JSON(http.StatusOK, user)

}
func GetUser(c *gin.Context) {
	var user []model.User
	database.DB.Find(&user)
	c.JSON(http.StatusOK, user)
}
