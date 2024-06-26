package routes

import (
	controller "task1/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/employee", controller.CreateRecord).Methods("POST")
	router.HandleFunc("/employees", controller.ReadAllRecords).Methods("GET")
	router.HandleFunc("/update/{id}", controller.UpdateRecords).Methods("POST")
	router.HandleFunc("/delete/{id}", controller.DeleteOne).Methods("DELETE")

	return router
}
