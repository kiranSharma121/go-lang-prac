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

	}
	MentorGroup := router.Group("/mentor")
	MentorGroup.Use(middleware.Authentication, middleware.MentorOnly())
	{

	}
	AdminGroup := router.Group("/admin")
	AdminGroup.Use(middleware.Authentication, middleware.AdminOnly())
	{

	}

	return router

}
