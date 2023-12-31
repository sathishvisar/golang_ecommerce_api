package common

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// GetPageAndLimit extracts page and limit from URL query parameters and provides default values if not present.
func GetPageAndLimit(r *http.Request) (int, int) {
	// Parse the URL query parameters
	queryParams := r.URL.Query()

	// Get a specific parameter value by key
	page := queryParams.Get("page")
	limit := queryParams.Get("limit")

	// Perform pagination query
	pageNumber := 1
	perPage := 10

	if page != "" {
		pageNumber, _ = strconv.Atoi(page)
	}

	if limit != "" {
		perPage, _ = strconv.Atoi(limit)
	}

	return pageNumber, perPage
}

// QueryDatabase executes the provided query on the database and returns the result rows.
func QueryDatabase(db *sql.DB, query string) (*sql.Rows, error) {
	// Execute the query on the database
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Defer closing the rows to ensure it happens after the query
	defer rows.Close()

	return rows, nil
}

// RespondWithJSON marshals the provided data to JSON and writes it to the response.
func RespondWithJSON(w http.ResponseWriter, data interface{}) {
	// Marshal the data to JSON
	responseJSON, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the client
	w.Write(responseJSON)
}
