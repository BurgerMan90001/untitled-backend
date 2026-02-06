package config

import (
	"context"
	"database/sql"
	"io"
	"log"
	"strconv"
	"time"

	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
)

//"postgres://username:password@localhost:5432/database_name?sslmode=mode"


func RunDatabase() {
	cfg, err := pgConfigFromEnv()

	if err != nil {
		log.Fatalf("Postgres configuration error %v", err)
	}
	
	embedCfg := embedBuildConfig(cfg)

	embedDb := embeddedpostgres.NewDatabase(embedCfg)


	if err := embedDb.Start(); err != nil {
		panic(err)
	}

	log.Printf("Postgres running on %s\n", embedCfg.GetConnectionURL())


	defer embedDb.Stop()


	DatabaseConnect(cfg.string())

}
func DatabaseConnect(url string) {
	const timeout = 5*time.Second

	db, err := sql.Open("postgres",url)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()

    if err := db.PingContext(ctx); err != nil {
        panic(err)
    }

    log.Println("Ping successful")
}

func embedBuildConfig(cfg pgconfig) embeddedpostgres.Config {
	portNum, err := strconv.Atoi(cfg.port)
	if err != nil {
		panic(err)
	}
	return embeddedpostgres.DefaultConfig().
		Username(cfg.username).
		Password(cfg.password).
		Database(cfg.database).
		Port(uint32(portNum)).
		Logger(io.Discard)

}