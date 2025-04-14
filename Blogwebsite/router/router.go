package router

import (
	"github.com/gin-gonic/gin"
	"github.com/goVendor/event"
)

func Router() *gin.Engine {
	server := gin.Default()
	server.POST("/signup", event.SignUp)
	server.POST("/login", event.Login)
	return server
}
