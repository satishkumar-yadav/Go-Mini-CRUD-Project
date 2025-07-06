✅ 1. README.md with Setup Instructions for this Go + MySQL Mini CRUD Project:

# 📘 Go CRUD API with MySQL

This project is a beginner-friendly REST API using **Go**, **Gorilla Mux**, and **MySQL** for managing a simple list of books.

---

## 📁 Folder Structure

- go-crud/
- ├── controllers/ → Business logic
- ├── db/ → MySQL connection
- ├── models/ → Data models (Book struct)
- ├── routes/ → API route setup
- ├── .env → DB credentials
- ├── go.mod / go.sum → Go dependencies
- └── main.go → Entry point

---

## 🛠️ Tech Stack

- **Golang**
- **MySQL**
- **Gorilla Mux (Router)**
- **godotenv (env loader)**
- **MySQL Driver**

---

## Dependencies/ Drivers Required :

- Mux
- GoDotEnv
- MySQL Driver

```
go get github.com/gorilla/mux
go get github.com/joho/godotenv
go get github.com/go-sql-driver/MySQL
```

---

## 🚀 Getting Started

### 1️⃣ Clone and navigate

```bash
git clone https://github.com/satishkumar-yadav/go-crud.git
cd go-crud
```

2️⃣ Set up .env file
Create a .env file:

```env
DB_USER={mysqlusername}
DB_PASSWORD={mysqlpassword}
DB_HOST=localhost
DB_PORT=3306
DB_NAME={databaseName} i.e. bookdb
```

3️⃣ Create the database
Run this SQL:

```mysql
CREATE DATABASE bookdb;

USE bookdb;

CREATE TABLE books (
  id INT AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(100),
  author VARCHAR(100),
  rating INT
);
```

4️⃣ Run the project

```
go mod tidy
go run main.go
```

✅ Server will run at:

```
http://localhost:8080
```

---

📬 API Endpoints

- Method URL Description
- GET /books Get all books - endpoint replaced with /satish for custom path testing
- POST /books Create a book
- PUT /books/{id} Update a book
- DELETE /books/{id} Delete a book - endpoint replaced with /satish/{id}

---

🔧 Test with Postman
POST /books

```json
{
  "title": "Go in Action",
  "author": "William Kennedy",
  "rating": 5
}
```

---

### 🧠 Learnings

- ✅ Project structure
- ✅ Clean MVC pattern
- ✅ REST API in Go
- ✅ MySQL connectivity
- ✅ JSON parsing, routing, and CRUD logic

---

✨ Author
Built with 💙 by Satish Kumar Yadav!

---

## ✅ 2. Postman Testing Guide

To test this API, follow these steps in **Postman** or **Thunder Client**:

---

### ➕ Create a Book (`POST /books`)

- URL: `http://localhost:8080/books`
- Method: `POST`
- Body → `raw JSON`:

```json
{
  "title": "The Go Programming Language",
  "author": "Alan Donovan",
  "rating": 5
}
```

✅ Response should return the created book with id.

---

### 📖 Get All Books (GET /books)

- URL: http://localhost:8080/books
- Method: GET
- ✅ Returns all books as array

---

### 🖊️ Update Book (PUT /books/{id})

- URL: http://localhost:8080/books/1
- Method: PUT
- Body → raw JSON:

```json
{
  "title": "Updated Title",
  "author": "Updated Author",
  "rating": 4
}
```

---

### ❌ Delete Book (DELETE /books/{id})

- URL: http://localhost:8080/books/1
- Method: DELETE
- ✅ Returns { "message": "Book deleted" }

---

### ✅ 4. Add Validation & Better Error Handling

We'll now:

- ✅ Check for empty fields
- ✅ Return 400 Bad Request for invalid input
- ✅ Return 404 if book ID doesn't exist

•
