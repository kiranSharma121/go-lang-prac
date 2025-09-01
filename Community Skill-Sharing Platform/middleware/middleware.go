package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skillplatform/controller"
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
	c.Set("id", id)
	c.Next()
}
