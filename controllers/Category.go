package controllers

import (
	"ecommerce-api/common"
	"ecommerce-api/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateCategoryData struct {
	Name string `json:"name"`
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
	}

	var requestData CreateCategoryData
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		http.Error(w, "Error decoding JSON data", http.StatusBadRequest)
		return
	}

	stmt, err := db.DB.Prepare("INSERT INTO categories(name) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(&requestData.Name)
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
	w.Write([]byte("Category created successfully"))

}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	// Use the shared function to get page and limit
	pageNumber, perPage := common.GetPageAndLimit(r)

	// Calculate OFFSET based on pageNumber and perPage
	offset := (pageNumber - 1) * perPage

	// Perform SELECT query with pagination
	query := fmt.Sprintf("SELECT * FROM categories LIMIT %d OFFSET %d", perPage, offset)

	// Use the shared function to query the database
	rows, err := db.DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Create a slice to store the Categories
	var categories []Category

	// Iterate through the rows and populate the Categories slice
	for rows.Next() {
		var category Category
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		categories = append(categories, category)
	}

	// Use the shared function to respond with JSON
	common.RespondWithJSON(w, categories)
}
