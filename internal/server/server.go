package server

import (
	"database/sql"
	"log"
	"net/http"

	// side effects
	_ "github.com/jackc/pgx/v5"

	"github.com/BurgerMan90001/untitled-backend/internal/config"
	"github.com/BurgerMan90001/untitled-backend/internal/middleware"
	"github.com/BurgerMan90001/untitled-backend/internal/model"
	"github.com/BurgerMan90001/untitled-backend/internal/util"
)

type Server struct {
	db *sql.DB
}


func NewServer(db *sql.DB) *Server {
	s := Server{db}
	return &s;
}

func (s Server) createUser(w http.ResponseWriter, r *http.Request) {
	var user = model.User {
		Id: "123",
		Firstname: "dasdasd",
		Lastname: "Wwww",
		Age: 32,
	}

	util.WriteJSON(w, user)
}

func Run() {
	const url = "localhost:8080"

	var databaseUrl string

	
	// setup database
	cfg := config.PgConfigFromEnv()
	databaseUrl = cfg.String()
	db := config.DatabaseConnect(databaseUrl)

	// setup server
	mux := http.NewServeMux()
	server := NewServer(db)
	

	// setup routes
	mux.HandleFunc("GET /", server.createUser)
	

	// add middleware
	handler := middleware.Logger(mux)

	
    log.Fatal(http.ListenAndServe(url, handler))
	log.Printf("Server listening at %s", url)
}