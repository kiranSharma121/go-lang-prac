package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/movie/models"
)

func CreateMovies(c *gin.Context) {
	var movies models.Movie
	err := c.ShouldBindJSON(&movies)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to bind the json",
			"error":   err.Error(),
		})
		return
	}
	role := c.GetString("role")
	userid := c.GetInt64("userid")

	fmt.Println()

	fmt.Println("role is ", role, " and userId is ", userid)

	movies.Userid = userid
	fmt.Println("----------------------")
	fmt.Println(movies.Userid)
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "only admin can create movies",
		})
		return
	}

	err = movies.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to save data in the movies table",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "stored data in the movie table successfully",
	})

}

func UpDateMovies(c *gin.Context) {
	movieid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "invalid params id",
			"error":   err.Error(),
		})
		return
	}
	userid := c.GetInt64("userid")
	role := c.GetString("role")
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "only admin can update movies",
		})
		return
	}
	movie, err := models.GetMoviesById(int(movieid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't get movie with that id",
			"error":   err.Error(),
		})
		return
	}
	if (movie.Movieid) != int64(userid) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized to update the movies",
		})
		return
	}
	var Updatemovie models.Movie
	err = c.ShouldBindJSON(&Updatemovie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to bind the json",
			"error":   err.Error(),
		})
		return
	}
	Updatemovie.Userid = userid
	Updatemovie.Movieid = movieid
	err = Updatemovie.Updatemovie()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to update the movie",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "update the movies sucessfully",
	})
}
func DeleteMovies(c *gin.Context) {
	movieid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "invalid params id",
			"error":   err.Error(),
		})
		return
	}
	userid := c.GetInt64("userid")
	role := c.GetString("role")
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "only admin can create movies",
		})
		return
	}
	movie, err := models.GetMoviesById(int(movieid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to get movies form the id",
			"error":   err.Error(),
		})
		return
	}
	if int(movie.Movieid) != int(userid) {
		fmt.Printf("%T", movie.Movieid)
		fmt.Printf("%T", userid)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized user",
		})
		return
	}
	err = movie.Deletemovies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to delete the movie",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted movie successfully",
	})

}
func CreatShowTime(c *gin.Context) {
	var shows models.Showtime
	err := c.ShouldBindJSON(&shows)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to bind the json",
		})
		return
	}
	shows.StartAT = time.Now()
	userid := c.GetInt64("userid")
	role := c.GetString("role")
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "only admin can create movies",
		})
		return
	}
	shows.Userid = userid

	err = shows.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to store data in the database",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "showtable is created...",
	})

}
func CreateSeats(c *gin.Context) {
	var seats models.Seat
	err := c.ShouldBindJSON(&seats)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "couldn't bind the json",
			"error":   err.Error(),
		})
		return
	}
	role := c.GetString("role")
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "only admin can create movies",
		})
		return
	}
	err = seats.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "couldn't save in the database",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "stored the data in the seats",
	})
}
func UpDateShowTime(c *gin.Context) {
	showtimeid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to get id form params",
			"error":   err.Error(),
		})
		return
	}
	userid := c.GetInt64("userid")
	showtime, err := models.Getshowsbyid(int(showtimeid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to get showtime form the id",
			"error":   err.Error(),
		})
		return
	}
	var show models.Showtime
	err = c.ShouldBindJSON(&show)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to bind the json",
		})
		return
	}
	showtime.Showtimeid = showtimeid
	showtime.Movieid = show.Movieid
	showtime.Userid = userid
	showtime.StartAT = show.StartAT
	showtime.Capacity = show.Capacity
	err = showtime.Updateshowtime()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error in updateshowtime in database",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "showtime is updated successfully",
	})
}
