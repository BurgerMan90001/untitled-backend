package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/BurgerMan90001/untitled-backend/internal/config/database"
	"github.com/BurgerMan90001/untitled-backend/internal/middleware"
	"github.com/BurgerMan90001/untitled-backend/internal/model/responses"
	"github.com/BurgerMan90001/untitled-backend/internal/util"
)

/*
type Server struct {
	db *sql.DB
}


func NewServer(db *sql.DB) *Server {
	s := Server{db}
	return &s;
}
*/

func environmentSetup(environment string) (*sql.DB)  {
	var db *sql.DB
	switch environment {
	case "dev": 
		// create emb
		//embedDb, databaseUrl := database.CreateEmbedDatabase()

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
	//db := environmentSetup(*environment)
	//defer db.Close()

	// setup server
	mux := http.NewServeMux()


	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		util.WriteJSON(w, responses.Response{Message: "aildlasdasd"})
	})
	
	//setupRoutes(mux, db)

	// add middleware
	handler := middleware.Logger(mux)

	// start server
	log.Printf("Server listening at %s", serverUrl)
    log.Fatal(http.ListenAndServe(serverUrl, handler))
}
