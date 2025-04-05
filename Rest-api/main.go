package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/restapi/models"
)

func main() {
	server := gin.Default()
	server.GET("/events", GetEvents)
	server.POST("/events", CreateEvents)
	server.Run(":8080")

}
func GetEvents(c *gin.Context) {
	events := models.GetAllEvents()
	c.JSON(http.StatusOK, events)

}
func CreateEvents(c *gin.Context) {
	var events models.Events
	err := c.ShouldBindJSON(&events)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "couldn't bind to json",
		})
		return
	}
	events.Id = rand.Int() % 100
	events.UserId = rand.Int() % 100
	events.Time = time.Now()
	events.Save()
	c.JSON(http.StatusOK, gin.H{
		"message": "created the event successfully",
	})

}
