package main

import (
	"fmt"
	"net/http"

	"github.com/kiransharma121/mongodb/router"
)

func main() {
	fmt.Println("Welcome to the mongoAPI")
	r := router.Router()
	fmt.Println("Server is starting...")
	http.ListenAndServe(":3200", r)
	fmt.Println("Listing to the port 4000")
}
