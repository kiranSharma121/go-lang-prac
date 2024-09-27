package route

import (
	"github.com/gorilla/mux"
	"github.com/kiransharma121/mangodb/controller"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/api/fullnames", controller.GetAllName).Methods("GET")
	router.HandleFunc("/api/fullname", controller.CreateName).Methods("POST")
	router.HandleFunc("/api/fullname/{id}", controller.UpDateOneName).Methods("PUT")
	router.HandleFunc("/api/deletenames", controller.DeleteAllName).Methods("DELETE")
	router.HandleFunc("/api/deletename/{id}", controller.DeleteOneName).Methods("DELETE")
	return router
}
