package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tutorplatform/database"
	"github.com/tutorplatform/model"
)

func StudentDashboard(c *gin.Context) {
	id, _ := c.Get("id")
	role, _ := c.Get("role")

	var user model.User
	database.DB.First(&user, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "welcome to the student dashboard",
		"user_id": id,
		"role":    role,
		"name":    user.Name,
		"Options": []string{
			"Book a class",
			"View upcoming sessions",
			"chat with tutors",
		},
	})
}

func TutorDashboard(c *gin.Context) {
	id, _ := c.Get("id")
	role, _ := c.Get("role")
	c.JSON(http.StatusOK, gin.H{
		"message": "welcome to the tutor dashboard",
		"user_id": id,
		"role":    role,
		"options": []string{
			"book a class",
			"View upcoming sessions",
			"chat with tutors",
		},
	})
}
