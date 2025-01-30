package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// InitDB initializes the database connection
func InitDB() {
	// Get the environment variables for DB connection
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Check for missing environment variables
	if dbHost == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		log.Fatal("Missing one or more environment variables for DB connection")
	}

	// Format connection string
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName,
	)

	var err error
	// Open a database connection (this does not establish the connection yet)
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database connection: ", err)
	}

	// Ping the database to make sure it is reachable
	if err = DB.Ping(); err != nil {
		log.Fatal("Database unreachable: ", err)
	}

	// Set connection pool parameters
	DB.SetMaxOpenConns(25)   // Set max open connections
	DB.SetMaxIdleConns(25)   // Set max idle connections
	DB.SetConnMaxLifetime(5) // Set connection lifetime (in minutes)

	log.Println("Database connected successfully")
}

// CloseDB closes the database connection
func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Fatal("Error closing database connection: ", err)
	}
	log.Println("Database connection closed successfully")
}


