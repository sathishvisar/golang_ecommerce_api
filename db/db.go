package db

import (
	"database/sql"
	"fmt"
	"log"
)

// DB is the package-level variable to store the database connection
var DB *sql.DB

// InitDB initializes the database connection
func InitDB() *sql.DB {
	fmt.Println("Init db....")

	// Database connection parameters
	dbUser := "root"
	dbPassword := "1234"
	dbHost := "127.0.0.1:3306"
	dbName := "ecommerce_store"

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
