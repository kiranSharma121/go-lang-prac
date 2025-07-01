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
	course.TutorId = int(tutorid.(int))
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
