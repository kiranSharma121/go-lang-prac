package router

import (
	"github.com/gorilla/mux"
	"github.com/kiransharma121/registration/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controller.ServeHome).Methods("GET")
	router.HandleFunc("/signup", controller.Signup).Methods("POST")
	router.HandleFunc("/login", controller.Login).Methods("POST")
	return router
}
