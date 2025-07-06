package models

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

// ✅ This struct represents a row in the books table.
