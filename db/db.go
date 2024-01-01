package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// DB is the package-level variable to store the database connection
var DB *sql.DB

// InitDB initializes the database connection
func InitDB() *sql.DB {
	// Load variables from the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("Init db....")

	// Database connection parameters
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DATABASE")

	// Create a connection string
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbHost, dbName)

	// Open a database connection
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	// Assign the opened database connection to the package-level DB variable
	DB = db

	fmt.Println("Database connected...")
	return db
}

// CloseDB closes the database connection
func CloseDB() {
	// Close the database connection when the application exits
	if DB != nil {
		defer DB.Close()
	}
}
