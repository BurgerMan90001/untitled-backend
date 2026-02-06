package main

import (
	"fmt"
	"os"

	"github.com/BurgerMan90001/untitled-backend/internal/server"
	"github.com/joho/godotenv"
)


func main() {

    loadEnv()

    server.Run()
}


func loadEnv() {
    err := godotenv.Load()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not load env %v.\n", err)
	}
    
}