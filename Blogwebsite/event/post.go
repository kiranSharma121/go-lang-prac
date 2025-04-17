package event

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goVendor/models"
)

func CreatePost(c *gin.Context) {
	var post models.Post
	err := c.ShouldBindJSON(&post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to bind the json",
		})
		return
	}
	userid := c.GetInt64("userId")
	post.Authorid = int(userid)

	err = post.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "couldn't able to store data in the database",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "stored data in the database",
	})

}
func GetPosts(c *gin.Context) {
	post, err := models.GetAllPost()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to get post from the database",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, post)
}
func UpDatePosts(c *gin.Context) {
	postid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid params id",
			"error":   err.Error(),
		})
		return
	}
	userid := c.GetInt64("userId")
	post, err := models.GetPostById(int(postid))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "unable to fetch post form the given id",
			"error":   err.Error(),
		})
		return
	}
	if int(post.Postid) != int(userid) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized to update the posts",
		})
		return
	}
	var upDatePost models.Post
	err = c.ShouldBindJSON(&upDatePost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to bind the json ",
			"error":   err.Error(),
		})
		return
	}
	upDatePost.Postid = postid
	err = upDatePost.UpDatePost()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to update post in the database",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "updated post in the database sucessfully",
	})
}
