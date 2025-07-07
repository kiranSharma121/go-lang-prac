package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tutorplatform/controller"
)

func Authentication(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "missing token",
		})
		c.Abort()
		return
	}
	role, id, err := controller.Verifyjwttoken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized token",
			"error":   err.Error(),
		})
		c.Abort()
		return
	}
	c.Set("role", role)
	c.Set("id", id)
	c.Next()
}
func TutorOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exist := c.Get("role")
		if !exist || role != "tutor" {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Access restriced for tutor only ",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
func StudentOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exit := c.Get("role")
		if !exit || role != "student" {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "only for the student",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
