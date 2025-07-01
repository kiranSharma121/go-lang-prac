package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tutorplatform/controller"
	"github.com/tutorplatform/middleware"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.POST("/signup", controller.Signup)
	router.GET("/getuser", controller.GetUser)
	router.POST("/login", controller.Login)
	studentGroup := router.Group("/student")
	studentGroup.Use(middleware.Authentication, middleware.StudentOnly())
	{
		studentGroup.GET("/dashboard", controller.StudentDashboard)
	}
	TeacherGroup := router.Group("/tutor")
	TeacherGroup.Use(middleware.Authentication, middleware.TutorOnly())
	{
		TeacherGroup.GET("/dashboard", controller.TutorDashboard)
		TeacherGroup.POST("/course", controller.CreateCourse)
	}

	return router
}
