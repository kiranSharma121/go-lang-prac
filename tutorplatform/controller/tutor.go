package controller

import (
	"net/http"
	"strconv"

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
func UpdateCourses(c *gin.Context) {
	courseID, _ := strconv.Atoi(c.Param("id"))
	tutorID := c.GetInt("id")
	var course model.Course
	err := database.DB.First(&course, courseID).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "couldn't find the course",
			"error":   err.Error(),
		})
		return
	}
	if course.TutorID != uint(tutorID) {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "you can't update the course",
		})
		return
	}
	var input struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	err = c.ShouldBind(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't bind the json",
		})
	}
	course.Title = input.Title
	course.Content = input.Content
	database.DB.Save(&course)
	c.JSON(http.StatusOK, gin.H{
		"message": "course updated", "course": course,
	})
}
func EnrollStudent(c *gin.Context) {
	courseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "invalid course id",
		})
		return
	}
	tutorID := c.GetInt("id")
	var course model.Course
	err = database.DB.First(&course, courseID).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "course not found",
			"error":   err.Error(),
		})
		return
	}
	if course.TutorID != uint(tutorID) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "you can't access to the course",
		})
		return
	}
	var enrollment []model.Enrollment
	err = database.DB.Preload("Student").Where("course_id=?", courseID).Find(&enrollment).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get student",
		})
		return
	}
	var students []map[string]interface{}
	for _, e := range enrollment {
		student := map[string]interface{}{
			"ID":    e.Student.ID,
			"Name":  e.Student.Name,
			"Email": e.Student.Email,
		}
		students = append(students, student)
	}
	c.JSON(http.StatusOK, students)

}
