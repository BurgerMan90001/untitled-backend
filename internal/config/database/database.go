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

	createDatabase(db)

	return db
}

func openDatabase(url string) *sql.DB {
	fmt.Println("Connecting to postgres database...")
	db, err := sql.Open("postgres", url)

	if err != nil {
		panic(err)
	}
	return db
}

func createDatabase(db *sql.DB) {

}

func pingDatabase(db *sql.DB) {
	const timeout = 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		panic(err)
	}

	log.Println("Ping successful")
}
