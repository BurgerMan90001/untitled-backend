package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

func DatabaseConnect() *sql.DB {
	cfg := pgConfigFromEnv()
	databaseUrl := cfg.String()
	return DatabaseConnectURL(databaseUrl)
}

func DatabaseConnectURL(url string) *sql.DB {
	const timeout = 5*time.Second

	fmt.Printf("Connecting to %s\n", url)
	db, err := sql.Open("postgres", url)

	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()

    if err := db.PingContext(ctx); err != nil {
        panic(err)
    }

    log.Println("Ping successful")

	

	return db
}
func initDatabase() {}


