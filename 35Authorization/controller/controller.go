package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kiransharma121/Authorization/database"
	"github.com/kiransharma121/Authorization/middleware"
	"github.com/kiransharma121/Authorization/model"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func ServeHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "welcome to the serve home",
	})

}
func hasedPassword(password string) string {
	hasedString, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error in generating the hased password")
	}
	return string(hasedString)
}
func signup(user model.LearningSystem) error {
	hasedstring := hasedPassword(user.Password)
	user.Password = hasedstring
	inserted, err := database.Collection.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Println("Error in inserting the data in the database")
	}
	fmt.Println("Inserted with the insert id:", inserted.InsertedID)
	return nil
}
func Signup(c *gin.Context) {
	var userinfo model.LearningSystem
	err := c.ShouldBindJSON(&userinfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Invalid information",
		})
	}
	if userinfo.Username == "" || userinfo.Email == "" || userinfo.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "fields cant be empty",
		})
		return
	}
	var existinguse model.LearningSystem
	err = database.Collection.FindOne(context.TODO(), bson.M{"email": userinfo.Email}).Decode(&existinguse)
	if err == nil {
		c.JSON(http.StatusAlreadyReported, gin.H{
			"message": "Account is already register",
		})
		return
	}
	err = signup(userinfo)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Error in signup",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Data is stored in the database",
	})
}
func verifypassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func Login(c *gin.Context) {
	var user model.LearningSystem
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Invalid userinformation",
		})
		return
	}
	var registeruser model.LearningSystem
	filter := bson.M{"email": user.Email}
	err = database.Collection.FindOne(context.TODO(), filter).Decode(&registeruser)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Email is invalid",
		})
		return
	}
	if verifypassword(registeruser.Password, user.Password) {
		tokenString, err := middleware.GenerateToken(user.Username, user.Email, user.Role)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"message": "Error in generating the token",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "Login Successfully",
				"Token":   tokenString,
			})
		}

	}

}
