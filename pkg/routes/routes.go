package routes

import (
	"budgetManager/pkg/handlers"
	"database/sql"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, db *sql.DB) {
	RegisterCategoryRoutes(r, db)
	// Register other resource routes here

	r.HandleFunc("/health", handlers.HealthCheckHandler).Methods("GET")
}
