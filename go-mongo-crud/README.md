Let’s now move to build the same CRUD API using Go + MongoDB, keeping everything similar to the MySQL version.

---

🧩 MongoDB CRUD Project Overview
We will build the same Book Manager API, but instead of MySQL, we’ll use MongoDB.
MongoDB stores data as flexible JSON-like documents (BSON), and works great with Go using the official MongoDB Go Driver.

---

✅ Step 1: Folder Structure for MongoDB CRUD
We’ll keep a similar structure to your MySQL version:

- go-mongo-crud/
- ├── controllers/ → Business logic
- │ └── book.go
- ├── db/ → MongoDB connection
- │ └── db.go
- ├── models/ → Book model (BSON)
- │ └── book.go
- ├── routes/ → Route definitions
- │ └── routes.go
- ├── .env → MongoDB credentials
- ├── go.mod / go.sum → Go modules
- └── main.go → Entry point

---

✅ Step 2: Create .env File for MongoDB

```env
MONGO_URI=mongodb://localhost:27017
MONGO_DB=bookdb
```

📝 Replace localhost with MongoDB cloud URI if using Atlas.

---

✅ Step 3: Initialize Go Project

```bash
go mod init go-mongo-crud
```

---

✅ Step 4: Install Dependencies

```bash
go get github.com/gorilla/mux
go get github.com/joho/godotenv
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/mongo/options
```

---

✅ Step 5: Create MongoDB Connection (db/db.go)
Here’s a clear and simple MongoDB connection file:

```go
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

func Connect() {
err := godotenv.Load()
if err != nil {
log.Fatal("❌ Error loading .env file")
}

    uri := os.Getenv("MONGO_URI")
    dbName := os.Getenv("MONGO_DB")

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
    	log.Fatal("❌ MongoDB Connection Error:", err)
    }

    Client = client
    BookCollection = client.Database(dbName).Collection("books")

    log.Println("✅ Connected to MongoDB")

}
```

---

✅ Explanation for Beginners:
Line What it does
godotenv.Load() Loads the .env variables
os.Getenv(...) Gets Mongo URI and DB name
context.WithTimeout(...) Limits DB connection time
mongo.Connect(...) Connects Go app to MongoDB
BookCollection = client.Database(...).Collection(...) Saves reference to books collection
log.Println(...) Confirms connection success

---

✅ With this, we’ve connected Go to MongoDB and can now move forward to:

- 1. Create the Book model in BSON
- 2. Write controllers for GET, POST, PUT, DELETE
- 3. Setup routes
- 4. Run and test!

---
