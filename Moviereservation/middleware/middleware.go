package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/movie/utils"
)

func Authentication(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	fmt.Println("---------------")
	fmt.Print(token)
	fmt.Println("---------")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "missing token",
		})
		c.Abort()
		return
	}
	role, userid, err := utils.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized token",
		})
		return
	}
	c.Set("role", role)
	c.Set("userid", userid)
	c.Next()
}
