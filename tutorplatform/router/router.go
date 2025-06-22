package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tutorplatform/controller"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.POST("/signup", controller.Signup)
	router.GET("/getuser", controller.GetUser)
	router.POST("/login", controller.Login)
	return router

}
