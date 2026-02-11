package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/BurgerMan90001/untitled-backend/internal/model"
	"github.com/BurgerMan90001/untitled-backend/internal/model/responses"
	"github.com/BurgerMan90001/untitled-backend/internal/util"
)

/*
func queryUserById(db *sql.DB, id string) model.User {
	rows, err := db.QueryContext(context.TODO(), "SELECT * FROM users WHERE id=?", id)
	if err != nil {
		log.Fatal(err)
	}
	var user model.User
	var (
		username string
		email string
	)
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&username, &email); err != nil {
			log.Fatal(err)

		}
	}
	user = model.User{
		Id: id, Username: username, Email: email,
	}
	return user
}
*/
func root(w http.ResponseWriter, r *http.Request) {
	util.WriteJSON(w, responses.Response{Message: "root"})
}
func healthCheck(w http.ResponseWriter, r *http.Request) {
	util.WriteJSON(w, responses.Response{Message: "Still alive"})
}
func getUserById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		// validate request
		

		// query the user
		rows, err := db.Query("SELECT * FROM users WHERE id=?", id)

		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close()

		var (
			username string
			email string
		)
		
		for rows.Next() {
			if err := rows.Scan(&username, &email); err != nil {
				log.Fatal(err)
				
			}
		}
		user := model.User{
			Id: id, Username: username, Email: email,
		}
	
		util.WriteJSON(w, user)
	}
}

func createUserById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		//r.BasicAuth()
		util.WriteJSON(w, nil)
	}
}

func deleteUserById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//id := r.PathValue("id")

		// validatie 

		// query
		//user := queryUser(db, id)


		util.WriteJSON(w, responses.DeleteResponse{Message: "Successfully deleted user"})
	}
}

func testRoute(w http.ResponseWriter, r *http.Request) {
	util.WriteJSON(w, responses.Response{Message: "Root"})
}