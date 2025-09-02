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
