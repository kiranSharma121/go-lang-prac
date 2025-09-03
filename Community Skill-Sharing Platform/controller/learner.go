package controller

import (
	"net/http"
	"strconv"

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
func EnrollinCourse(c *gin.Context) {
	skillIDStr := c.Param("id")
	skillID, err := strconv.Atoi(skillIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid skill id",
			"error":   err.Error(),
		})
		return
	}
	learnerID, exist := c.Get("id")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "invalid learner id",
		})
	}
	var existing model.Enrollment
	err = database.DB.Where("skill_id=? AND learner_id=?", skillID, learnerID).First(&existing).Error
	if err == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Already enroll in the skill",
		})
		return
	}
	enrollment := model.Enrollment{
		SkillID:   uint(skillID),
		LearnerID: uint(learnerID.(int)),
	}
	err = database.DB.Create(&enrollment).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to enroll in the skill",
			"error":   err.Error(),
		})
		return
	}
	database.DB.Preload("Skill").Preload("Learner").First(&enrollment, enrollment.ID)
	c.JSON(http.StatusOK, enrollment)
}
