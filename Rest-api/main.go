package main

import (
	"github.com/gin-gonic/gin"
	"github.com/restapi/database"
	"github.com/restapi/routers"
)

func main() {
	database.InitDB()
	server := gin.Default()
	routers.Router(server)
	server.Run(":8080")
}
