package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StudentDashboard(c *gin.Context) {
	id, _ := c.Get("id")
	role, _ := c.Get("role")
	c.JSON(http.StatusOK, gin.H{
		"message": "welcome to the student dashboard",
		"user_id": id,
		"role":    role,
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
