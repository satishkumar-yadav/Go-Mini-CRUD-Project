package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var BookCollection *mongo.Collection

func ConnectMongoDb() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	uri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB")

	/*
		client, err := mongo.NewClient(options.Client().ApplyURI(uri))
		if err != nil {
			log.Fatal(err)
		} */

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("❌ MongoDB Connection Error:", err)
	}

	/*
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}  */

	// Test mongodb connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Could not connect to MongoDb", err)
	}
	//

	Client = client
	BookCollection = client.Database(dbName).Collection("books")
	// assign collection to controller
	// controllers.BookCollection = bookCollection

	log.Println("✅ Connected to MongoDB")
}
