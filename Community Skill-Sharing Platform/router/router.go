package router

import (
	"github.com/gin-gonic/gin"
	"github.com/skillplatform/controller"
	"github.com/skillplatform/middleware"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.POST("/signup", controller.SignUp)
	router.POST("/login", controller.Login)
	router.GET("/ws/chat", controller.ChatHandler)
	LearnerGroup := router.Group("/learner")
	LearnerGroup.Use(middleware.Authentication, middleware.LearnerOnly())
	{
		LearnerGroup.GET("/skills", controller.GetSkills)
		LearnerGroup.POST("/skill/:id", controller.EnrollinCourse)
		LearnerGroup.GET("/myskills", controller.EnrolledCourse)
		LearnerGroup.POST("/sessions", controller.CreateSession)
		LearnerGroup.GET("/mysessions", controller.ListSession)
	}
	MentorGroup := router.Group("/mentor")
	MentorGroup.Use(middleware.Authentication, middleware.MentorOnly())
	{
		MentorGroup.POST("/skills", controller.CreateSkill)
		MentorGroup.DELETE("/skill/:id", controller.DeleteSkill)
		MentorGroup.POST("/skill/:id", controller.UpdateSkill)
		MentorGroup.POST("/session/:id", controller.UpdateSession)
	}
	AdminGroup := router.Group("/admin")
	AdminGroup.Use(middleware.Authentication, middleware.AdminOnly())
	{
		AdminGroup.GET("/users", controller.ListUser)
		AdminGroup.GET("/enrollments", controller.ListEnrollments)
		AdminGroup.DELETE("/users/:id", controller.DeleteUser)
		AdminGroup.GET("/skills", controller.ListSkills)
		AdminGroup.DELETE("/skill/:id", controller.Deleteskill)

	}

	return router

}
