package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/movie/models"
	"github.com/movie/utils"
)

func Signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to bind the json",
			"error":   err.Error(),
		})
		return
	}
	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "couldn't able to save user in the database",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "sucessfully stored user in the database",
	})

}
func Login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to bind the json",
			"error":   err.Error(),
		})
		return
	}
	err = user.Validatecredentials()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to validate user",
			"error":   err.Error(),
		})
		return
	}
	token, err := utils.GeneratejwtToken(user.Userid, user.UserName, user.Email, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error in generating the token",
			"error":   err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"token":   token,
	})
}
func GetAllMovies(c *gin.Context) {
	movies, err := models.Getallmovies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to get movies form database",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, movies)

}
func GetAllShowTime(c *gin.Context) {
	shows, err := models.Getallshowtime()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't get showtime form shows",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, shows)
}
func Getmoviebyid(c *gin.Context) {
	movieid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "invalid params id",
			"error":   err.Error(),
		})
		return
	}
	movie, err := models.GetMoviesById(int(movieid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "couldn't get movie by id",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, movie)
}
func GetAllSeats(c *gin.Context) {
	seats, err := models.Allseats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to get seats from the database",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, seats)
}
func BookingSeats(c *gin.Context) {
	var booking models.Booking
	err := c.ShouldBindJSON(&booking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "couldn't bind the json",
			"error":   err.Error(),
		})
		return
	}
	booking.BookedAt=time.Now();
	role := c.GetString("role")
	if role != "user" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "user can only book the movie",
		})
		return
	}
	userid := c.GetInt64("userid")
	if booking.UserID != userid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "user is not authorized",
		})
		return
	}
	err = booking.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to store data in the bookings",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "booked sucessfully",
	})

}
