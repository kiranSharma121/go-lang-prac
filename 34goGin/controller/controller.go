package controller

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/kiransharma121/gin/database"
	"github.com/kiransharma121/gin/model"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func ServeHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to the serveHom",
	})
}
func hasedPassword(password string) string {
	hasedString, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error in hasing the password")
	}
	return string(hasedString)
}
func signUp(user model.GinAuth) error {
	hasedstring := hasedPassword(user.Password)
	user.Password = hasedstring
	inserted, err := database.Collection.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Println("Error in inserting the data in the database")
	}
	fmt.Println("Inserted with the insert id:", inserted.InsertedID)
	return err
}
func Signup(c *gin.Context) {
	var userinfo model.GinAuth
	err := c.ShouldBindJSON(&userinfo)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Invalid information",
		})
	}
	if userinfo.UserName == "" || userinfo.Email == "" || userinfo.Password == "" {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "username,email and password shouldn't be empty",
		})
		return

	}
	var existinguser model.GinAuth
	err = database.Collection.FindOne(context.TODO(), bson.M{"email": userinfo.Email}).Decode(&existinguser)
	if err == nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"Message": "Registered already",
		})
		return
	}
	err = signUp(userinfo)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Error in signup",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Data stored in the database",
	})

}
func verifyPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil

}
func Login(c *gin.Context) {
	var user struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Invalid information",
		})
	}
	var dbuser model.GinAuth
	err = database.Collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&dbuser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Email",
		})
		return
	}
	if verifyPassword(dbuser.Password, user.Password) {
		token, err := GenerateToken(user.Email, user.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Could not generate token",
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "Login successful",
				"token":   token,
			})
		}

	}
}

var secretKey = []byte("student_secrete_key")

type claims struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(email, username string) (string, error) {
	expireTime := time.Now().Add(24 * time.Hour)
	claims := &claims{
		Email:    email,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    username,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", nil
	}
	return tokenString, nil
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
