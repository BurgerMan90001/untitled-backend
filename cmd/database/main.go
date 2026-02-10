package main

import (
	"github.com/BurgerMan90001/untitled-backend/internal/config/database"
)


func main() {
	// setup database
	db := database.DatabaseConnect()

	defer db.Close()
}