package routers

import "github.com/gin-gonic/gin"

func Router() *gin.Engine {
	server := gin.Default()
	return server
}
