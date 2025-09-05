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
func DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	var user model.User
	err := database.DB.Find(&user, userID).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to find the user with that id",
		})
		return
	}
	err = database.DB.Delete(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to delete the user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted user successfully",
	})

}
