package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("student_secrete_key")

type claims struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization token required",
			})
			c.Abort()
			return
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer")
		token, _ := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		if claims, ok := token.Claims.(*claims); ok && token.Valid {
			c.Set("email", claims.Email)
			c.Set("username", claims.Username)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token or expired token",
			})
			c.Abort()
		}

	}
}
func Dashboard(c *gin.Context) {
	username, _ := c.Get("username")
	email, _ := c.Get("email")
	c.JSON(http.StatusOK, gin.H{
		"message":  "access granted to the student",
		"username": username,
		"email":    email,
	})
}
