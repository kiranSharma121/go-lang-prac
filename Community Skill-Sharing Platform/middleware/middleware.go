package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/skillplatform/controller"
	"github.com/skillplatform/database"
	"github.com/skillplatform/model"
)

func Authentication(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "missing jwt token",
		})
		c.Abort()
		return
	}
	role, id, err := controller.VerifyJwtToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized token",
			"error":   err.Error(),
		})
		c.Abort()
		return
	}
	c.Set("role", role)
	c.Set("id", uint(id))
	c.Next()
}
func LearnerOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != "learner" {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Access forbidden...Learner only",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
func MentorOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exits := c.Get("role")
		if !exits || role != "mentor" {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Access Forbidden...only for the mentor",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exits := c.Get("role")
		if !exits || role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Access Forbiden.. AdminOnly",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
func EnsureEnrollment(c *gin.Context) {
	userID := c.GetUint("id")

	otherIDStr := c.Query("receiver_id")
	otherIDUint64, err := strconv.ParseUint(otherIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid receiver_id"})
		c.Abort()
		return
	}
	otherID := uint(otherIDUint64)
	var enrollment model.Enrollment
	err = database.DB.
		Joins("JOIN skills ON skills.id = enrollments.skill_id").
		Where("(enrollments.learner_id = ? AND skills.user_id = ?) OR (enrollments.learner_id = ? AND skills.user_id = ?)",
			userID, otherID, otherID, userID).
		First(&enrollment).Error

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": "chat not allowed, no valid enrollment"})
		c.Abort()
		return
	}
	c.Next()
}
