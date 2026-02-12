package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/BurgerMan90001/untitled-backend/internal/config/database"
	"github.com/BurgerMan90001/untitled-backend/internal/middleware"
)

func environmentSetup(environment string) *sql.DB {
	var db *sql.DB
	switch environment {
	case "dev":
		db = database.DatabaseConnectEnv()

	case "prod":
		// connect to local database
		//db = database.DatabaseConnect()
	}
	return db
}

func Run() {
	const serverUrl = "localhost:8080"

	//environment := config.CreateFlags()

	// setup database
	db := database.DatabaseConnectEnv()

	defer db.Close()

	// setup server
	mux := http.NewServeMux()

	setupRoutes(mux, db)

	// add middleware
	handler := middleware.Logger(mux)

	// start server
	log.Printf("Server listening at %s", serverUrl)
	log.Fatal(http.ListenAndServe(serverUrl, handler))
}
