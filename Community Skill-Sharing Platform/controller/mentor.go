package controller

import (
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
