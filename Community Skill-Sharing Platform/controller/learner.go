package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skillplatform/database"
	"github.com/skillplatform/model"
)

func GetSkills(c *gin.Context) {
	var skill model.Skill
	err := database.DB.Find(&skill).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to find the skill",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, skill)
}
