package controllers

import (
	"net/http"

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

func Root(w http.ResponseWriter, r *http.Request) {
	util.WriteJSON(w, responses.Response{Message: "root"})
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	util.WriteJSON(w, responses.Response{Message: "Still alive"})
}

func TestRoute(w http.ResponseWriter, r *http.Request) {
	util.WriteJSON(w, responses.Response{Message: "Root"})
}
