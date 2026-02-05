package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"

	"github.com/BurgerMan90001/untitled-backend/internal/model"
)

type Server struct {
	
}

var user = model.User {
    Id: "123",
    Firstname: "dasdasd",
    Lastname: "Wwww",
    Age: 32,
}


func createServer() *Server {
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

	renderJSON(w, user)
}



func renderJSON(w http.ResponseWriter, v interface{}) {
    w.Header().Set("Content-type", "application/json")
    js, err := json.Marshal(v)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Write(js)
}

func Run() {
	const url = "localhost:8080"
	const urlDatabase = "postgres://username:password@localhost:5432/database_name?sslmode=mode"

	conn, err := pgx.Connect(context.TODO(), urlDatabase)

	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database %v\n", err)
	}

	defer conn.Close(context.TODO())
	


	mux := http.NewServeMux()
	server := createServer()

	mux.HandleFunc("GET /", server.createUser)
	

    err := http.ListenAndServe(url, mux)
    log.Fatal(err)
}