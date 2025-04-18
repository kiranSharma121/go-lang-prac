package event

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sqlite/models"
)

func GetEvents(c *gin.Context) {
	event, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to fetch data form the database",
		})
		return
	}
	c.JSON(http.StatusOK, event)
}
func CreateEvent(c *gin.Context) {
	var events models.Events
	err := c.ShouldBindJSON(&events)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to bind the json",
		})
		return
	}
	events.Time = time.Now()
	userid := c.GetInt64("userId")
	events.Userid = int(userid)
	err = events.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to save data in the database",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "create event ...",
	})
}
func GetEvent(c *gin.Context) {
	eventid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid params id",
		})
		return
	}
	event, err := models.GetEventByID(int(eventid))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "couldn't find the event",
		})
		return
	}
	c.JSON(http.StatusOK, event)

}
func UpDateEvent(c *gin.Context) {
	eventid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid params",
		})
		return
	}

	userid := c.GetInt64("userId")
	event, err := models.GetEventByID(int(eventid))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "can't fetch data with the given eventid",
		})
		return
	}
	if event.Userid != int(userid) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorized to update event",
		})

	}

	var upDatedEvent models.Events
	err = c.ShouldBindJSON(&upDatedEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid JSON in request body",
		})
		return
	}

	upDatedEvent.Id = eventid
	err = upDatedEvent.UpDateEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to update event in database",
		})
		return
	}

	c.JSON(http.StatusOK, upDatedEvent)
}
func DeleteEvent(c *gin.Context) {
	eventid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid params",
		})
		return
	}
	userid := c.GetInt64("userId")

	event, err := models.GetEventByID(int(eventid))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "can't get event by the given id",
		})
		return
	}
	if event.Userid != int(userid) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorized to update event",
		})

	}
	err = event.Delete()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "can't delete event from the database",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted the event from the database",
	})
}
