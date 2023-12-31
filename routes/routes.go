package routes

import (
	"ecommerce-api/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {

	router := mux.NewRouter()

	// Products API routes
	router.HandleFunc("/api/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/api/products/{id}", controllers.GetProduct).Methods("GET")
	router.HandleFunc("/api/categories", controllers.GetCategories).Methods("GET")

	return router
}

func StartServer() {
	http.Handle("/", SetupRouter())
	http.ListenAndServe(":8080", nil)
}
