package main

import (
	"fmt"

	"github.com/tutorplatform/database"
	"github.com/tutorplatform/router"
)

func main() {
	fmt.Println("tutor platform using golang")
	database.Connection()
	r := router.Router()
	r.Run(":8080")
}
