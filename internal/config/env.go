package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not load env %v.\n", err)
	}
}

func GetServerURL() string {
	return fmt.Sprintf("localhost:%s", os.Getenv("SERVER_PORT"))
}
