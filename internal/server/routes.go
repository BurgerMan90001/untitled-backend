package server

import (
	"database/sql"
	"net/http"

	"github.com/BurgerMan90001/untitled-backend/internal/controllers"
)

func setupRoutes(mux *http.ServeMux, db *sql.DB) {
	// user routes
	mux.HandleFunc("/user/{id}", controllers.UserController(db))



	// static files
	//mux.HandleFunc("GET /static", nil)

	//TODO auth routes
	//mux.HandleFunc("/login/oauth", middleware.)
	mux.HandleFunc("POST /auth/login", controllers.LoginHandler)


	mux.HandleFunc("/secret", controllers.Secret)
	mux.Handle("/static/", http.FileServer(http.Dir("public")))
}
