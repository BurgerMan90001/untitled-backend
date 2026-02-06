package server

import (
	"log"
	"net/http"

	// side effects
	_ "github.com/jackc/pgx/v5"

	"github.com/BurgerMan90001/untitled-backend/internal/config"
	"github.com/BurgerMan90001/untitled-backend/internal/model"
	"github.com/BurgerMan90001/untitled-backend/pkg/util"
)

type Server struct {}

var user = model.User {
    Id: "123",
    Firstname: "dasdasd",
    Lastname: "Wwww",
    Age: 32,
}


func NewServer() *Server {
	s := Server{}
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


	mux := http.NewServeMux()
	server := NewServer()

	mux.HandleFunc("GET /", server.createUser)
	
	log.Printf("Server listening at %s", url)

	config.RunDatabase()
	//tests.Test()

    log.Fatal(http.ListenAndServe(url, mux))
}