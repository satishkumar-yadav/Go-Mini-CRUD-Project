package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// ✅ Book defines the structure of our MongoDB document
type Book struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"  json:"id" `     // MongoDB ObjectId
	Title  string             `bson:"title"          json:"title" `  // Book title
	Author string             `bson:"author"         json:"author" ` // Author name
	Rating int                `bson:"rating"         json:"rating" ` // Rating from 1 to 5
}
