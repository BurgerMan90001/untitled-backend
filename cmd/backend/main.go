package main

import (
	"github.com/BurgerMan90001/untitled-backend/internal/config"
	"github.com/BurgerMan90001/untitled-backend/internal/server"

	_ "github.com/lib/pq"
)


func main() {
	config.LoadEnv()

	server.Run()
}