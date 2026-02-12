package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/BurgerMan90001/untitled-backend/internal/model"
	"github.com/BurgerMan90001/untitled-backend/internal/model/responses"
	"github.com/BurgerMan90001/untitled-backend/internal/util"
)

func UserController(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getUserById(db, w, r)
		//case "POST":
		case "UPDATE":
		case "DELETE":
			deleteUserById(db, w, r)
		}
	}
}

func getUserById(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	// validation
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Fatal(err)
	}

	response := model.QueryUserById(db, id)

	util.WriteJSON(w, response)
}

func createUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	//r.BasicAuth()
	query := "INSERT INTO users($1,$2,$3) VALUES(2,'soscannill1','rfarmery1@alibaba.com')"
	_, err := db.Exec(query)

	switch {
	case err == sql.ErrNoRows:
	case err != nil:
	default:

	}
	//rows, err := db.Query(query)
	util.WriteJSON(w, nil)

}

func deleteUserById(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	//id := r.PathValue("id")

	// validatie

	// query
	//user := queryUser(db, id)

	util.WriteJSON(w, responses.DeleteResponse{Message: "Successfully deleted user"})

}

func updateUserById(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	//id := r.PathValue("id")

	// validatie

	// query
	//user := queryUser(db, id)

	util.WriteJSON(w, responses.DeleteResponse{Message: "Successfully updated user"})

}

func getUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// var (
		// 	username string
		// 	email string
		// )

		// for rows.Next() {
		// 	if err := rows.Scan(&username, &email); err != nil {
		// 		log.Fatal(err)

		// 	}
		// }
		// user := model.User{
		// 	Id: id, Username: username, Email: email,
		// }
		util.WriteJSON(w, nil)
	}
}
