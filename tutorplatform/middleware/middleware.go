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
	}
	role, id, err := controller.Verifyjwttoken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized token",
		})
		return
	}
	c.Set("role", role)
	c.Set("id", id)
	c.Next()
}
