package main

import (
	"github.com/sqlite/database"
	"github.com/sqlite/router"
)

func main() {
	database.InitDB()
	server := router.Routers()
	server.Run(":8080")
}
