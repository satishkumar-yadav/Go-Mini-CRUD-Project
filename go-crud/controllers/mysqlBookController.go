// ==============================
// 📁 File: controllers/book.go (MySQL CRUD with Detailed Explanation)
// ✅ Beginner-Friendly Full Explanation of Every Line and Concept
// ==============================

package controllers

// Import standard and external packages we need
import (
	"encoding/json" // to convert Go structs to JSON and vice versa
	"errors"        // to define and return custom error messages
	"fmt"
	"net/http" // to build HTTP server and handle requests
	"strconv"  // to convert string to integer and vice versa
	"strings"  // to perform string operations like trimming spaces

	"github.com/satishkumar-yadav/Go-Mini-CRUD-Project/go-crud/db"     // our custom package for MySQL DB connection
	"github.com/satishkumar-yadav/Go-Mini-CRUD-Project/go-crud/models" // our custom package for the Book struct/model

	"github.com/gorilla/mux" // external package to build clean, RESTful route handling
)

// ✅ Helper Function: To validate incoming book data before inserting or updating
func validateBook(b models.Book) error {
	if strings.TrimSpace(b.Title) == "" ||
		strings.TrimSpace(b.Author) == "" ||
		b.Rating < 1 || b.Rating > 5 {
		return errors.New("Invalid input: title and author are required, and rating must be between 1 and 5")
	}
	return nil
}

// ✅ GET /books - Fetch all books from database
func GetBooks(w http.ResponseWriter, r *http.Request) {
	// Query to get all books
	rows, err := db.DB.Query("SELECT * FROM books")
	if err != nil {
		http.Error(w, "Database query error", http.StatusInternalServerError) //  http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close() // always close the rows after looping

	var books []models.Book
	for rows.Next() {
		var b models.Book
		rows.Scan(&b.ID, &b.Title, &b.Author, &b.Rating) // read columns into struct fields
		books = append(books, b)
	}

	w.Header().Set("Content-Type", "application/json") // tell client we are returning JSON
	json.NewEncoder(w).Encode(books)                   // send all books as JSON
	fmt.Println("Id : ", books[0], "Rating : ", books[3], "Title : ", books[1], "Author : ", books[2])
}

// ✅ POST /books - Insert a new book into database
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var b models.Book
	err := json.NewDecoder(r.Body).Decode(&b) // parse JSON body into struct
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if err := validateBook(b); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Prepare query to insert into MySQL
	result, err := db.DB.Exec("INSERT INTO books(title, author, rating) VALUES (?, ?, ?)", b.Title, b.Author, b.Rating)
	if err != nil {
		http.Error(w, "Database insert error", http.StatusInternalServerError)
		return
	}

	// Get the last inserted ID and attach it to the struct
	lastID, _ := result.LastInsertId()
	b.ID = int(lastID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(b) // return the newly created book as JSON
}

// ✅ PUT /books/{id} - Update a book's details
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"] // get ID from route param
	var b models.Book

	err := json.NewDecoder(r.Body).Decode(&b) // parse JSON into struct
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if err := validateBook(b); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := db.DB.Exec("UPDATE books SET title=?, author=?, rating=? WHERE id=?", b.Title, b.Author, b.Rating, id)
	if err != nil {
		http.Error(w, "Database update error", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	b.ID, _ = strconv.Atoi(id) // convert id back to int and assign
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(b)
}

// ✅ DELETE /books/{id} - Delete a book by ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"] // get ID from URL param

	res, err := db.DB.Exec("DELETE FROM books WHERE id=?", id)
	if err != nil {
		http.Error(w, "Database delete error", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Book deleted successfully"})
}

// ==============================
// 📌 Summary for Beginners:
// ------------------------------
// - "mux.Vars" gets URL params (like {id})
// - "json.NewDecoder().Decode" reads JSON body
// - "json.NewEncoder().Encode" writes JSON response
// - "db.DB.Exec" executes SQL commands
// - "rows.Scan" reads result from DB into variables
// - Every handler sets content type and writes JSON
// ==============================
