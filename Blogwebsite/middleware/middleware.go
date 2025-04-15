package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goVendor/models"
)

func Authentication(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized user",
		})
		return
	}
	userid, err := models.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "token is not verified",
		})
		return
	}
	c.Set("userid", userid)
	c.Next()

}
