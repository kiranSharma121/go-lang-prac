package router

import (
	"github.com/gin-gonic/gin"
	"github.com/goVendor/event"
	"github.com/goVendor/middleware"
)

func Router() *gin.Engine {
	server := gin.Default()
	server.POST("/signup", event.SignUp)
	server.POST("/login", event.Login)
	server.POST("/posts", middleware.Authentication, event.CreatePost)
	return server
}
