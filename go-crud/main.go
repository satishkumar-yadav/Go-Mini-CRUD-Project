package main

import (
	"log"
	"net/http"

	"github.com/satishkumar-yadav/Go-Mini-CRUD-Project/go-crud/db"
	"github.com/satishkumar-yadav/Go-Mini-CRUD-Project/go-crud/routes"
)

func main() {
	db.ConnectMySql()

	r := routes.SetupRoutes()

	log.Println("🚀 Server running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
