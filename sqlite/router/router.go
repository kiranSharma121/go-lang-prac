package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sqlite/event"
	"github.com/sqlite/middlewares"
)

func Routers() *gin.Engine {
	server := gin.Default()
	server.GET("/events", event.GetEvents)
	server.POST("/events", middlewares.Authenicate, event.CreateEvent)
	server.GET("/events/:id", event.GetEvent)
	server.PUT("/events/:id", middlewares.Authenicate, event.UpDateEvent)
	server.DELETE("/events/:id", middlewares.Authenicate, event.DeleteEvent)
	server.POST("/signup", event.SignUp)
	server.POST("/login", event.Login)
	return server
}
