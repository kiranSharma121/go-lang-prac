package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tutorplatform/controller"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.POST("/signup", controller.CreateUser)
	router.GET("/getuser", controller.GetUser)
	return router

}
