package handlers

import (
	"budgetManager/pkg/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetAllCategoriesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		nameFilter := r.URL.Query().Get("name")
		categories, err := models.GetAllCategories(db, nameFilter)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(categories)
	}
}

func CreateCategoryHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Body)
		var c models.Category
		json.NewDecoder(r.Body).Decode(&c)

		fmt.Print(json.NewDecoder(r.Body).Decode(&c))
		if c.Name == "" {
			http.Error(w, "Category name cannot be empty", http.StatusBadRequest)
			return
		}
		err := models.CreateCategory(db, c.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func UpdateCategoryHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c models.Category
		json.NewDecoder(r.Body).Decode(&c)
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])
		err := models.UpdateCategory(db, id, c.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func DeleteCategoryHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])
		err := models.DeleteCategory(db, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
