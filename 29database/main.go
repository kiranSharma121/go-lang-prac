package main

import (
	"fmt"
	"net/http"

	"github.com/kiransharma121/mangodb/route"
)

func main() {
	fmt.Println("Welcome to data base")
	r := route.Router()
	fmt.Println("The router is begin started...")
	http.ListenAndServe(":4000", r)
}
