package main

import (
	"log"
	"net/http"

	"go-mongo-crud/controllers"
	"go-mongo-crud/db"
	"go-mongo-crud/routes"
)

func main() {

	db.ConnectMongoDb()

	// Now db.BookCollection is initialized, so assign it to controller
	controllers.Collection = db.BookCollection

	r := routes.SetupRoutes()

	log.Println("🚀 Server running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
