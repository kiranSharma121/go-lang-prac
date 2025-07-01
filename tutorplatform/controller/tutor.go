package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tutorplatform/database"
	"github.com/tutorplatform/model"
)

func CreateCourse(c *gin.Context) {
	var course model.Course
	err := c.ShouldBindJSON(&course)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't bind the json",
		})
		return
	}
	tutorid, _ := c.Get("id")
	course.TutorID = uint(tutorid.(int))
	err = database.DB.Create(&course).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "couldn't create the course",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, course)
}
func DeleteCourse(c *gin.Context) {
	tutorid, _ := c.Get("id")
	courseId := c.Param("id")
	var course model.Course
	err := database.DB.First(&course, courseId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "can't find the course with that id",
			"error":   err.Error(),
		})
		return
	}
	if course.TutorID != uint(tutorid.(int)) {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Only you can delete the course that you have created",
		})
		return

	}
	err = database.DB.Delete(&course).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to delete the course",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted the course successfully",
	})
}
