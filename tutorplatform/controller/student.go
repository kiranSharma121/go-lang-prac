package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tutorplatform/database"
	"github.com/tutorplatform/model"
)

func GetCourse(c *gin.Context) {
	var course []model.Course
	err := database.DB.Find(&course).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "couldn't fetch the course",
		})
		return
	}
	c.JSON(http.StatusOK, course)
}
func EnrollInCourse(c *gin.Context) {
	courseIDstr := c.Param("courseID")
	courseID, err := strconv.Atoi(courseIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}
	studentID, _ := c.Get("id")
	var existing model.Enrollment
	err = database.DB.Where("course_id=? AND student_id=?", courseID, studentID).First(&existing).Error
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Already enroll in the course",
		})
		return
	}
	entrollment := model.Enrollment{
		CourseID:  uint(courseID),
		StudentID: uint(studentID.(int)),
	}
	err = database.DB.Create(&entrollment).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "enrollment failed",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Enrollment in the course successfully",
	})
}
