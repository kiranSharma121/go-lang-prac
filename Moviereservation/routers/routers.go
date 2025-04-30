package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/movie/controllers"
	"github.com/movie/middleware"
)

func Router() *gin.Engine {
	server := gin.Default()
	server.POST("/signup", controllers.Signup)
	server.POST("/login", controllers.Login)
	server.POST("/movies", middleware.Authentication, controllers.CreateMovies)
	server.GET("/movies", controllers.GetAllMovies)
	server.GET("/movies/:id", controllers.Getmoviebyid)
	server.PUT("/movies/:id", middleware.Authentication, controllers.UpDateMovies)
	server.DELETE("/movies/:id", middleware.Authentication, controllers.DeleteMovies)
	server.POST("/shows", middleware.Authentication, controllers.CreatShowTime)
	server.POST("/seats", middleware.Authentication, controllers.CreateSeats)
	server.GET("/seats", controllers.GetAllSeats)
	return server
}
