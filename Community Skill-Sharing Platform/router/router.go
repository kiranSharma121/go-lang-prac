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
	LearnerGroup := router.Group("/learner")
	LearnerGroup.Use(middleware.Authentication, middleware.LearnerOnly())
	{
		LearnerGroup.GET("/skills", controller.GetSkills)
		LearnerGroup.POST("/skill/:id", controller.EnrollinCourse)
		LearnerGroup.GET("/myskills", controller.EnrolledCourse)
	}
	MentorGroup := router.Group("/mentor")
	MentorGroup.Use(middleware.Authentication, middleware.MentorOnly())
	{
		MentorGroup.POST("/skills", controller.CreateSkill)
		MentorGroup.DELETE("/skill/:id", controller.DeleteSkill)
		MentorGroup.POST("/skill/:id", controller.UpdateSkill)
	}
	AdminGroup := router.Group("/admin")
	AdminGroup.Use(middleware.Authentication, middleware.AdminOnly())
	{
		AdminGroup.GET("/users", controller.ListUser)
		AdminGroup.GET("/enrollments", controller.ListEnrollments)
		AdminGroup.DELETE("/users/:id", controller.DeleteUser)

	}

	return router

}
