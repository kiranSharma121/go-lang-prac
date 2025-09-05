package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skillplatform/database"
	"github.com/skillplatform/model"
)

func ListUser(c *gin.Context) {
	var user []model.User
	err := database.DB.Where("role IN ?", []string{"mentor", "learner"}).Find(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't find the users",
		})
		return
	}
	c.JSON(http.StatusOK, user)

}
func ListEnrollments(c *gin.Context) {
	var enrollment []model.Enrollment
	err := database.DB.Preload("Learner").Preload("Skill").Find(&enrollment).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to find the enrollments",
		})
		return
	}
	c.JSON(http.StatusOK, enrollment)
}
func DeleteUser(c *gin.Context){
	
}
