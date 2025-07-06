package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/satishkumar-yadav/Go-Mini-CRUD-Project/go-crud/db"
	"github.com/satishkumar-yadav/Go-Mini-CRUD-Project/go-crud/models"

	"github.com/gorilla/mux"
)

// GET /books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT * FROM books")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var b models.Book
		rows.Scan(&b.ID, &b.Title, &b.Author, &b.Rating)
		books = append(books, b)
	}
	json.NewEncoder(w).Encode(books)
	fmt.Println("Id : ", books[0], "Rating : ", books[3], "Title : ", books[1], "Author : ", books[2])
}

// POST /books
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var b models.Book
	json.NewDecoder(r.Body).Decode(&b)

	result, err := db.DB.Exec("INSERT INTO books(title, author, rating) VALUES (?, ?, ?)", b.Title, b.Author, b.Rating)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lastID, _ := result.LastInsertId()
	b.ID = int(lastID)
	json.NewEncoder(w).Encode(b)
}

// PUT /books/{id}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var b models.Book
	json.NewDecoder(r.Body).Decode(&b)

	_, err := db.DB.Exec("UPDATE books SET title=?, author=?, rating=? WHERE id=?", b.Title, b.Author, b.Rating, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b.ID, _ = strconv.Atoi(id)
	json.NewEncoder(w).Encode(b)
}

// DELETE /books/{id}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	_, err := db.DB.Exec("DELETE FROM books WHERE id=?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Book deleted"})
}
