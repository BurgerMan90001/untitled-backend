package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/BurgerMan90001/untitled-backend/internal/middleware"
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


		util.WriteJSON(w, model.DeleteResponse{Message: "Successfully deleted user"})
	}
}

func testRoute(w http.ResponseWriter, r *http.Request) {

	util.WriteJSON(w, model.Response{Message: "Root"})
	
}

func Run() {
	const url = "localhost:8080"
	
	// setup database
	
	//db := database.DatabaseConnect()
	//defer db.Close()

	// setup server
	mux := http.NewServeMux()

	// setup routes
	mux.HandleFunc("/", testRoute)
	// mux.HandleFunc("GET /user/{id}", getUserById(db))
	// mux.HandleFunc("POST /user/{id}", createUserById(db))
	// mux.HandleFunc("DELETE /user/{id}", deleteUserById(db))
	
	// add middleware
	handler := middleware.Logger(mux)

	// start server
	log.Printf("Server listening at %s", url)
    log.Fatal(http.ListenAndServe(url, handler))
}
