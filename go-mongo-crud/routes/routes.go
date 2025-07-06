package routes

import (
	"go-mongo-crud/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter() // mux Router for defining HTTP routes

	r.HandleFunc("/book", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")

	return r
}
