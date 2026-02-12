package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

/* Connects to database from env */
func DatabaseConnectEnv() *sql.DB {
	databaseUrl := GetConnectURL()

	db := DatabaseConnectURL(databaseUrl)

	return db
}
/* Connects to database from specified string */
func DatabaseConnectURL(url string) *sql.DB {
	
	db := openDatabase(url)
	// test database connection
	pingDatabase(db)
	
	return db
}

func openDatabase(url string) *sql.DB {
	fmt.Printf("Connecting to %s\n", url)
	db, err := sql.Open("postgres", url)

	if err != nil {
		panic(err)
	}
	return db
}

func pingDatabase(db *sql.DB) {
	const timeout = 5*time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()

    if err := db.PingContext(ctx); err != nil {
        panic(err)
    }

    log.Println("Ping successful")
}


