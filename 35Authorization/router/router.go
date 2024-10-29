package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kiransharma121/Authorization/controller"
	"github.com/kiransharma121/Authorization/middleware"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.GET("/", controller.ServeHome)
	router.POST("/signup", controller.Signup)
	router.POST("/login", controller.Login)
	router.GET("/dashboard", middleware.AuthMiddleware(), middleware.Dashboard)
	return router
}
