package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tutorplatform/controller"
	"github.com/tutorplatform/middleware"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.POST("/signup", controller.Signup)
	router.POST("/login", controller.Login)
	studentGroup := router.Group("/student")
	studentGroup.Use(middleware.Authentication, middleware.StudentOnly())
	{
		studentGroup.GET("/dashboard", controller.StudentDashboard)
		studentGroup.GET("/courses", controller.GetCourse)
	}
	TeacherGroup := router.Group("/tutor")
	TeacherGroup.Use(middleware.Authentication, middleware.TutorOnly())
	{
		TeacherGroup.GET("/dashboard", controller.TutorDashboard)
		TeacherGroup.POST("/course", controller.CreateCourse)
		TeacherGroup.DELETE("/course/:id", controller.DeleteCourse)
	}

	return router
}
