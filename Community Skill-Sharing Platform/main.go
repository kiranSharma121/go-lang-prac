package main

import (
	"fmt"

	"github.com/skillplatform/database"
	"github.com/skillplatform/router"
)

func main() {
	fmt.Println("community skill share platform")
	database.Connection()
	r := router.Router()
	r.Run(":8080")
}
