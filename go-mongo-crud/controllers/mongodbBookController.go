package controllers

import (
	"context"
	"encoding/json"
	"go-mongo-crud/models"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var Collection *mongo.Collection // Instead of assigning early
var collection = Collection

//var collection = db.BookCollection

// ====== Helper: validate input ======
func validateBook(book models.Book) string {
	if strings.TrimSpace(book.Title) == "" {
		return "Title cannot be empty"
	}
	if strings.TrimSpace(book.Author) == "" {
		return "Author cannot be empty"
	}
	if book.Rating < 1 || book.Rating > 5 {
		return "Rating must be between 1 and 5"
	}
	return ""
}

// ====== GET /books ======
func GetBooks(w http.ResponseWriter, r *http.Request) {
	if Collection == nil {
		http.Error(w, "Book Collection is not initialized", http.StatusInternalServerError)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //context.TODO()
	defer cancel()

	cursor, err := Collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Failed to fetch books", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var books []models.Book
	for cursor.Next(ctx) {
		var book models.Book
		cursor.Decode(&book)
		books = append(books, book)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// ====== POST /books ======
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)

	if msg := validateBook(book); msg != "" {
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	book.ID = primitive.NewObjectID()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := Collection.InsertOne(ctx, book)
	if err != nil {
		http.Error(w, "Failed to insert book", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// ====== PUT /books/{id} ======
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	objectId, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)

	if msg := validateBook(book); msg != "" {
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": bson.M{
		"title":  book.Title,
		"author": book.Author,
		"rating": book.Rating,
	}}

	result, err := Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		http.Error(w, "Update failed", http.StatusInternalServerError)
		return
	}
	if result.MatchedCount == 0 {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	book.ID = objectId
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// ====== DELETE /books/{id} ======
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	objectId, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := Collection.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		http.Error(w, "Delete failed", http.StatusInternalServerError)
		return
	}
	if result.DeletedCount == 0 {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Book deleted"})
}
