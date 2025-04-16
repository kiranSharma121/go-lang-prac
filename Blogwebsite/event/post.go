package event

import (
	"net/http"

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
	}
	c.JSON(http.StatusOK, post)
}
