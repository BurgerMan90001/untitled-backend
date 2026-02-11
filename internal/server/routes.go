package server

import (
	"database/sql"
	"net/http"
)


func setupRoutes(mux *http.ServeMux, db *sql.DB) {
	// setup routes
	mux.HandleFunc("/", root)
	mux.HandleFunc("GET /health", healthCheck)
	mux.HandleFunc("GET /user/{id}", getUserById(db))
	mux.HandleFunc("POST /user/{id}", createUserById(db))
	mux.HandleFunc("DELETE /user/{id}", deleteUserById(db))
}