package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secreteKey = []byte("student-secrete-key")

type claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(username, email string) (string, error) {
	expireTime := time.Now().Add(1 * time.Hour)
	claims := claims{
		Username: username,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    username,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secreteKey)
	if err != nil {
		fmt.Println("Error in generating the token")
	}
	return string(tokenString), nil
}
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authentication error",
			})
			c.Abort()
			return
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer")
		token, _ := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (interface{}, error) {
			return secreteKey, nil
		})
		if claims, ok := token.Claims.(*claims); ok && token.Valid {
			c.Set("username", claims.Username)
			c.Set("email", claims.Email)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token or expired token",
			})
		}
		c.Abort()
	}
}
func Dashboard(c *gin.Context) {
	username, _ := c.Get("username")
	email, _ := c.Get("email")
	c.JSON(http.StatusOK, gin.H{
		"message":  "Access granted to the student",
		"username": username,
		"email":    email,
	})
}
