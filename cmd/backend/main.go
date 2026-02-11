package main

import (
	"github.com/BurgerMan90001/untitled-backend/internal/config"
	"github.com/BurgerMan90001/untitled-backend/internal/server"
)


func main() {
	config.LoadEnv()

	//config.CreateFlags()
	// setup database
	// embedDb := database.RunEmbedDatabase()
	// db := database.DatabaseConnectURL(embedDb.GetConnectionURL())
	server.Run()
	//defer embedDb.Stop()

	//defer db.Close()
}