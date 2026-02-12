package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/BurgerMan90001/untitled-backend/internal/model/responses"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"name"`
	Email    string `json:"email"`
}

func QueryUserById(db *sql.DB, id int) any {
	var (
		username string
		email    string
		response any
	)

	// query the user
	query := "SELECT username, email FROM users WHERE id=$1"
	err := db.QueryRowContext(context.TODO(), query, id).Scan(&username, &email)

	switch {
	case err == sql.ErrNoRows:
		message := fmt.Sprintf("No users found with id %d", id)
		response = responses.Response{Message: message}
	case err != nil:
		errorMessage := fmt.Sprintf("Query error: %v", err)
		response = responses.Error{Error: errorMessage}
	default:
		response = User{
			Id:       id,
			Username: username,
			Email:    email,
		}
	}
	return response
}

/* Create a user example
var user = User {
    Id: "123",
    Firstname: "dasdasd",
    Lastname: "Wwww",
    Age: 32,
}
*/
