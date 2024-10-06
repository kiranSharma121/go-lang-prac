package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kiransharma121/auth/router"
)

func main() {
	fmt.Println("welcome to the authentication in golang")
	fmt.Println("server is starting...")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":8080", r))

}
