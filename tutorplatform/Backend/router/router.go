package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tutorplatform/controller"
	"github.com/tutorplatform/middleware"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.POST("/signup", controller.Signup)
	router.POST("/login", controller.Login)
	router.POST("/webhook", controller.Webhook)
	studentGroup := router.Group("/student")
	studentGroup.Use(middleware.Authentication, middleware.StudentOnly())
	{
		studentGroup.GET("/dashboard", controller.StudentDashboard)
		studentGroup.GET("/courses", controller.GetCourse)
		studentGroup.POST("/enroll/:courseID", controller.EnrollInCourse)
		studentGroup.GET("/my-courses", controller.EnrollCourse)
		studentGroup.POST("/pay/:id", controller.Payment)
	}
	TeacherGroup := router.Group("/tutor")
	TeacherGroup.Use(middleware.Authentication, middleware.TutorOnly())
	{
		TeacherGroup.GET("/dashboard", controller.TutorDashboard)
		TeacherGroup.POST("/course", controller.CreateCourse)
		TeacherGroup.DELETE("/course/:id", controller.DeleteCourse)
		TeacherGroup.PUT("/course/:id", controller.UpdateCourses)
		TeacherGroup.GET("/:id/students", controller.EnrollStudent)
	}
	return router
}
