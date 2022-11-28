package router

import (
	"github.com/biswaBandhuAich/mysql/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", controller.Home).Methods("GET")
	r.HandleFunc("/book/{id}", controller.GetOneBook).Methods("GET")
	r.HandleFunc("/books", controller.GetAllBooks).Methods("GET")
	r.HandleFunc("/book", controller.AddOneBook).Methods("POST")
	r.HandleFunc("/books", controller.AddBooks).Methods("POST")

	return r
}
