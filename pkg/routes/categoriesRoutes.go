package routes

import (
	"budgetManager/pkg/handlers"
	"database/sql"
	"github.com/gorilla/mux"
)

func RegisterCategoryRoutes(r *mux.Router, db *sql.DB) {
	r.HandleFunc("/categories", handlers.GetAllCategoriesHandler(db)).Methods("GET")
	r.HandleFunc("/categories", handlers.CreateCategoryHandler(db)).Methods("POST")
	r.HandleFunc("/categories/{id}", handlers.UpdateCategoryHandler(db)).Methods("PUT")
	r.HandleFunc("/categories/{id}", handlers.DeleteCategoryHandler(db)).Methods("DELETE")
	// Add other category-related routes here
}
