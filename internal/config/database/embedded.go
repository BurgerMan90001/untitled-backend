package database

import (
	"io"
	"log"
	"strconv"

	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
)

func RunEmbedDatabase(cfg pgconfig) {

	embedCfg := embedBuildConfig(cfg)

	embedDb := embeddedpostgres.NewDatabase(embedCfg)

	if err := embedDb.Start(); err != nil {
		panic(err)
	}

	log.Printf("Postgres running on %s\n", embedCfg.GetConnectionURL())

	defer embedDb.Stop()
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