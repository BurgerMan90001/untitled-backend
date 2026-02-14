package server

import (
	"database/sql"
	"net/http"

	"github.com/BurgerMan90001/untitled-backend/internal/controllers"
)

func setupRoutes(mux *http.ServeMux, db *sql.DB) {
	// setup routes
	mux.HandleFunc("/", controllers.Root)
	mux.HandleFunc("GET /health", controllers.HealthCheck)

	// user routes
	mux.HandleFunc("/user/{id}", controllers.UserController(db))



	

	// static files
	//mux.HandleFunc("GET /static", nil)

	//TODO auth routes
	//mux.HandleFunc("/login/oauth", middleware.)
	mux.Handle("/static/", http.FileServer(http.Dir("public")))
}
