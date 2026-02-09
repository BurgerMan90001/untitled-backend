package server

import (
	"log"
	"net/http"

	"github.com/BurgerMan90001/untitled-backend/internal/config"
	"github.com/BurgerMan90001/untitled-backend/internal/middleware"
)

func Run() {
	const url = "localhost:8080"
	
	// setup database
	cfg := config.PgConfigFromEnv()
	databaseUrl := cfg.String()
	db := config.DatabaseConnect(databaseUrl)

	// setup server
	mux := http.NewServeMux()

	// setup routes
	mux.HandleFunc("GET /user/{id}", getUser(db))
	mux.HandleFunc("POST /user/{id}", createUser(db))

	// add middleware
	handler := middleware.Logger(mux)

	// start server
	log.Printf("Server listening at %s", url)
    log.Fatal(http.ListenAndServe(url, handler))
}
