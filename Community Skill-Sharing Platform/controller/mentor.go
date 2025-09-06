package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skillplatform/database"
	"github.com/skillplatform/model"
)

func CreateSkill(c *gin.Context) {
	var skill model.Skill
	err := c.ShouldBindJSON(&skill)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to bind the json",
		})
		return
	}
	MentorID, _ := c.Get("id")
	skill.UserID = uint(MentorID.(int))
	err = database.DB.Create(&skill).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to create the skill table",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "created database successfully...",
		"skills":  skill,
	})
}
func DeleteSkill(c *gin.Context) {
	mentorID, _ := c.Get("id")
	skillID := c.Param("id")
	var skill model.Skill
	err := database.DB.First(&skill, skillID).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't find the skill with that id",
		})
		return
	}
	if skill.UserID != uint(mentorID.(int)) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "only the creater of the course can delete the course",
		})
		return
	}
	err = database.DB.Delete(&skill).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "fail to delete the skill from the database",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted skill successfully",
	})
}
func UpdateSkill(c *gin.Context) {
	skillID := c.Param("id")
	userID, _ := c.Get("id")
	var skill model.Skill
	err := database.DB.First(&skill, skillID).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to find the skill with that id",
		})
		return
	}
	if skill.UserID != uint(userID.(int)) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "only the author can edit the skill field",
		})
		return
	}
	var input struct {
		Title       string ` json:"title"`
		Description string `json:"description"`
		Category    string `json:"category"`
	}
	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to bind the json",
		})
		return
	}
	skill.Title = input.Title
	skill.Description = input.Description
	skill.Category = input.Category
	c.JSON(http.StatusOK, gin.H{
		"message": "skill field updated successfully",
		"skill":   skill,
	})
}
func UpdateSession(c *gin.Context) {
	mentorId, _ := c.Get("id")
	id := c.Param("id")
	var session model.Session
	err := database.DB.Find(&session, id).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to find the session with the id",
		})
		return
	}
	var input struct {
		Status string `json:"status"`
	}
	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to bind the json",
		})
		return
	}
	allowed := map[string]bool{"pending": true, "accepted": true, "rejected": true, "completed": true}
	if !allowed[input.Status] {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "invalid status",
		})
		return
	}
	session.Status = input.Status
	database.DB.Create(&session)
	notification := model.Notification{
		UserID:  session.LearnerID,
		Type:    "session_request",
		Content: fmt.Sprintf("You have a new session request from mentor %d", mentorId),
	}
	err = database.DB.Create(&notification).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to create the notification",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "session updated",
		"session": session,
	})
}
