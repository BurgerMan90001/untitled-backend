package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BurgerMan90001/untitled-backend/internal/config/database"
	"github.com/BurgerMan90001/untitled-backend/internal/controllers"
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


func GetServerURL() string {
	return fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
}

func Run() {
	serverUrl := GetServerURL()

	//environment := config.CreateFlags()

	// setup database
	db := database.DatabaseConnectEnv()

	defer db.Close()

	// setup server
	mux := http.NewServeMux()

	// setup routes
	mux.HandleFunc("/", controllers.Root)
	mux.HandleFunc("GET /health", controllers.HealthCheck)


	setupRoutes(mux, db)

	// add middleware
	handler := middleware.Logger(mux)
	//handler = middleware.PanicRecovery(mux)

	// start server
	log.Printf("Server listening at %s", serverUrl)
	log.Fatal(http.ListenAndServe(serverUrl, handler))
}
