package controllers

import (
	"database/sql"
	"ecommerce-api/common"
	"ecommerce-api/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Product represents a product in your application
type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category int     `json:"category_id"`
}

type CreateProductData struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category int     `json:"category_id"`
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	// Read the body data
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Unmarshal the JSON data into a struct
	var requestData CreateProductData
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		http.Error(w, "Error decoding JSON data", http.StatusBadRequest)
		return
	}

	// Prepare the SQL statement
	stmt, err := db.DB.Prepare("INSERT INTO products(name, price, category_id) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Execute the SQL statement
	result, err := stmt.Exec(requestData.Name, requestData.Price, requestData.Category)
	if err != nil {
		log.Fatal(err)
	}

	// Get the ID of the inserted row
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lastInsertID)

	// Respond to the client
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Product created successfully"))
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get the user ID from the request parameters
	params := mux.Vars(r)
	productID := params["id"]

	// Perform SELECT query with pagination
	query := fmt.Sprintf("SELECT * FROM products WHERE id = %s", productID)

	// Use the shared function to query the database
	row := db.DB.QueryRow(query)

	// Create a Product variable to store the result
	var product Product

	// Scan the result into the Product variable
	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Category)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			http.Error(w, "Product not found", http.StatusInternalServerError)
			return
		}
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Respond with the fetched product in JSON format
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, `{"id":%d,"name":"%s","price":%f,"category_id": %d}`, product.ID, product.Name, product.Price, product.Category)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	// Use the shared function to get page and limit
	pageNumber, perPage := common.GetPageAndLimit(r)

	// Calculate OFFSET based on pageNumber and perPage
	offset := (pageNumber - 1) * perPage

	// Perform SELECT query with pagination
	query := fmt.Sprintf("SELECT * FROM products LIMIT %d OFFSET %d", perPage, offset)

	// Use the shared function to query the database
	rows, err := db.DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Create a slice to store the products
	var products []Product

	// Iterate through the rows and populate the products slice
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Category)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}

	// Use the shared function to respond with JSON
	common.RespondWithJSON(w, products)
}
