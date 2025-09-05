package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

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
func EnrolledCourse(c *gin.Context) {
	learnerID, _ := c.Get("id")
	var enrollment []model.Enrollment
	err := database.DB.Preload("Skill").Where("learner_id=?", learnerID).Find(&enrollment).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to fetch the data",
		})
		return
	}
	var skills []model.Skill
	for _, e := range enrollment {
		skills = append(skills, e.Skill)
	}
	c.JSON(http.StatusOK, skills)
}
func CreateSession(c *gin.Context) {
	learnerID, exist := c.Get("id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized user,user not exist",
		})
		return
	}
	var input struct {
		MentorID uint      `json:"mentor_id"`
		Time     time.Time `json:"time"`
	}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to bind the json",
		})
		return
	}
	session := model.Session{
		LearnerID: uint(learnerID.(int)),
		MentorID:  input.MentorID,
		Time:      input.Time,
		Status:    "pending",
	}
	err = database.DB.Create(&session).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to create session table in database",
		})
		return
	}
	notification := model.Notification{
		UserID:  input.MentorID,
		Type:    "session_request",
		Content: fmt.Sprintf("You have a new session request from learner %d", learnerID),
	}
	err = database.DB.Create(&notification).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to create the table notification in the database",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "session has been created successfully",
		"session": session,
	})

}
