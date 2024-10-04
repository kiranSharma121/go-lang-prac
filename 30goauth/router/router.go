package router

import (
	"github.com/gorilla/mux"
	"github.com/kiransharma121/goauth/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.ServeHome).Methods("GET")
	router.HandleFunc("/register", controllers.RegisterHandler).Methods("POST")
	return router
}
