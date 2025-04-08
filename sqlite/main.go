package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sqlite/database"
	"github.com/sqlite/models"
)

func main() {
	database.InitDB()
	server := gin.Default()
	server.GET("/events", GetEvents)
	server.POST("/events", CreateEvent)
	server.Run()
}
func GetEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "couldn't fetch the event from the database",
		})
		return
	}
	c.JSON(http.StatusOK, events)

}
func CreateEvent(c *gin.Context) {
	var event models.Events
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "unable to bind the request",
		})
		return
	}
	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "couldn't save the data in the database",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "inserted data in the database successfully",
	})

}
