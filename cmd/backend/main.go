package main

import (
	"github.com/BurgerMan90001/untitled-backend/internal/config"
	"github.com/BurgerMan90001/untitled-backend/internal/server"
)


func main() {
	config.LoadEnv()

    server.Run()
}


