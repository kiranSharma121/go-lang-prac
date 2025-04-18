package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/movie/controllers"
)

func Router() *gin.Engine {
	server := gin.Default()
	server.POST("/signup", controllers.Signup)
	server.POST("/login", controllers.Login)
	return server
}
