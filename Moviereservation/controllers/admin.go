package controllers

import (
	"fmt"
	"net/http"

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
