package router

import (
	"github.com/gorilla/mux"
	"github.com/kiransharma121/auth/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controller.ServeHome).Methods("GET")
	router.HandleFunc("/getalluser", controller.Getallusers).Methods("GET")
	router.HandleFunc("/insertoneuser", controller.InsertOneUser).Methods("POST")
	return router

}
