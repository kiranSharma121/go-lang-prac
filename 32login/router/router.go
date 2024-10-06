package router

import (
	"github.com/gorilla/mux"
	"github.com/kiransharma121/login/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controller.ServeHome).Methods("GET")
	router.HandleFunc("/insertone", controller.Insertoneuser).Methods("POST")
	router.HandleFunc("/login", controller.Loginhandler)
	router.HandleFunc("/getalluser", controller.Getalluser).Methods("GET")
	router.HandleFunc("/User/{id}", controller.Updateone).Methods("PUT")
	router.HandleFunc("/user/{id}", controller.Deleteoneuser).Methods("DELETE")
	return router

}
