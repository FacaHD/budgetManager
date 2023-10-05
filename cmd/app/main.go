package main

import (
	"budgetManager/pkg/routes"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/budgetManager")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := mux.NewRouter()

	routes.RegisterRoutes(r, db)

	log.Fatal(http.ListenAndServe(":8080", r))
}
