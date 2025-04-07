package routers

import (
	"math/rand/v2"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/restapi/models"
)

func GetEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error in fatching data from the database",
		})
	}
	c.JSON(http.StatusOK, events)

}
func GetEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "unable to find the event id",
		})
		return
	}
	event, err := models.GetEventById(int(eventId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not find the event",
		})
		return
	}
	c.JSON(http.StatusOK, event)
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
	events.UserId = rand.Int() % 100
	events.Time = time.Now()
	err = events.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "couldnot able to create events",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "created the event successfully",
	})

}
func UpDateEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "unable to find the event id",
		})
		return
	}
	_, err = models.GetEventById(int(eventId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not find the event",
		})
		return
	}
	var updatedevent models.Events
	err = c.ShouldBindJSON(&updatedevent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "couldn't bind to json",
		})
		return
	}
	updatedevent.Id = eventId
	err = updatedevent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not find the event",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "updated the event successfully",
	})
}
func DeleteEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event ID",
		})
		return
	}

	event, err := models.GetEventById(int(eventId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Event not found",
		})
		return
	}

	err = event.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete event",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully",
	})
}
