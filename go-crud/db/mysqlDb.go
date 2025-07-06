package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func ConnectMySql() {
	// godotenv -	Load .env file into os.Getenv()
	err := godotenv.Load() // godotenv.Load() - loads .env variables into environment
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	// driver name, dataSourceName
	database, err := sql.Open("mysql", dsn) // sql.Open() prepares connection
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	if err := database.Ping(); err != nil { // Ping() confirms DB is alive/responding
		log.Fatal("Ping error: ", err)
	}

	log.Println("✅ Connected to MySQL")
	DB = database
}
