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
	return server
}
