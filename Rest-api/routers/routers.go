package routers

import "github.com/gin-gonic/gin"

func Router(server *gin.Engine) {
	server.GET("/events", GetEvents)
	server.POST("/events", CreateEvents)
	server.GET("/events/:id", GetEvent)
	server.PUT("/events/:id", UpDateEvent)
	server.DELETE("/events/:id", DeleteEvent)
}
