package router

import (
	"github.com/gin-gonic/gin"
	"github.com/skillplatform/controller"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.POST("/signup", controller.SignUp)
	return router

}
