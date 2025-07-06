package routes

import (
	"github.com/satishkumar-yadav/Go-Mini-CRUD-Project/go-crud/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter() // mux Router for defining HTTP routes

	r.HandleFunc("/satish", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/satish/{id}", controllers.DeleteBook).Methods("DELETE")

	/*
	   ✅ API Endpoints
	      Method	Endpoint	  Description
	      GET	    /books	      Get all books  - endpoint replaced with /satish for custom path testing
	      POST	    /books	      Add a book
	      PUT	    /books/{id}	  Update book by ID
	      DELETE	/books/{id}	  Delete book by ID  - endpoint replaced with /satish/{id}
	*/

	return r
}
