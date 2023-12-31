package main

import (
	"ecommerce-api/db"
	"ecommerce-api/routes"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Initialize the database connection
	db.InitDB()

	// Defer the closure of the database connection when the application exits
	defer db.CloseDB()

	// Start the server
	routes.StartServer()
}
