package database

import (
	"io"
	"log"
	"strconv"

	embpg "github.com/fergusstrange/embedded-postgres"
)

func CreateEmbedDatabase() (*embpg.EmbeddedPostgres, string) {
	cfg := pgConfigFromEnv()
	embedCfg := embedBuildConfig(cfg)
	embedDb := embpg.NewDatabase(embedCfg)

	startEmbedDatabase(embedDb)
	
	log.Printf("Postgres running on %s\n", embedCfg.GetConnectionURL())
	
	return embedDb, embedCfg.GetConnectionURL()
}
func startEmbedDatabase(embedDb *embpg.EmbeddedPostgres) {
	if err := embedDb.Start(); err != nil {
		panic(err)
	}
}

func embedBuildConfig(cfg pgconfig) embpg.Config {
	portNum, err := strconv.Atoi(cfg.port)
	if err != nil {
		panic(err)
	}
	return embpg.DefaultConfig().
		Username(cfg.username).
		Password(cfg.password).
		Database(cfg.database).
		Port(uint32(portNum)).
		Logger(io.Discard)
}