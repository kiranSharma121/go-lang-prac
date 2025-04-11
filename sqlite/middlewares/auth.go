package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sqlite/models"
)

func Authenicate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "not authorize",
		})
		return
	}
	userid, err := models.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid or expired token",
		})
		return
	}
	c.Set("userId", userid)
	c.Next()
}
