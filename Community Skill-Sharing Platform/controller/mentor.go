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
	MentorID := c.GetUint("id")
	skill.UserID = MentorID
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
	mentorID := c.GetUint("id")
	skillID := c.Param("id")
	var skill model.Skill
	err := database.DB.First(&skill, skillID).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't find the skill with that id",
		})
		return
	}
	if skill.UserID != mentorID {
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
	userID := c.GetUint("id")
	var skill model.Skill
	err := database.DB.Find(&skill, skillID).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to find the skill with that id",
		})
		return
	}
	if skill.UserID != userID {
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
	mentorId := c.GetUint("id")
	id := c.Param("id")

	var session model.Session
	if err := database.DB.First(&session, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "session not found"})
		return
	}

	var input struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request", "error": err.Error()})
		return
	}

	allowed := map[string]bool{"pending": true, "accepted": true, "rejected": true, "completed": true}
	if !allowed[input.Status] {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid status"})
		return
	}

	err := database.DB.Save(&session).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "unable to update session"})
		return
	}

	notification := model.Notification{
		UserID:  session.LearnerID,
		Type:    "session_status",
		Content: fmt.Sprintf("Your session with mentor %d is now %s", mentorId, input.Status),
	}
	if err := database.DB.Create(&notification).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "unable to create notification"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "session updated",
		"session": session,
	})
}
