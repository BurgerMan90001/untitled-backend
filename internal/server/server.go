package server

import (
	"database/sql"
	"net/http"

	"github.com/BurgerMan90001/untitled-backend/internal/model"
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

func getUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		var user = model.User {
			Id: id,
			Firstname: "dasdasd",
			Lastname: "Wwww",
			Age: 32,
		}
		util.WriteJSON(w, user)
	}
}

func createUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.BasicAuth()
	}
}

func deleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}