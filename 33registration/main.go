package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kiransharma121/registration/router"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")                   // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS") // Allowed methods
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")       // Allowed headers

		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	fmt.Println("Welcome to the registration page")
	r := router.Router()
	fmt.Println("Server is running in portno:4000")
	log.Fatal(http.ListenAndServe(":4000", corsMiddleware(r)))
}
