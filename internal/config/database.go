package config

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
)

//"postgres://username:password@localhost:5432/database_name?sslmode=mode"
type pgconfig struct {
	user string
	password string
	database string
	host string
	sslmode string
	port string
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func readEnvFile() {
	path := filepath.Join(os.TempDir(), "dat")
    dat, err := os.ReadFile(path)
    checkErr(err)
    fmt.Print(string(dat))
}
func connect() {
	/*
	var pgconfig = getConfig()
	//config, err := pgconfig{}.getConfig()

	if err != nil {
		panic(err)
	}
	*/
}

/* For testing */
func runEmbedded() {
	postgres := embeddedpostgres.NewDatabase()
}	
func createConfig() (pgconfig, error) {
	var missing []string
	
	getConfigVal := func(key string) string {
		val := os.Getenv(key)
		if val == "" {
			missing = append(missing, key)
		}
		return val
	}

	cfg := pgconfig{
		user:     getConfigVal("PG_USER"),
        database: getConfigVal("PG_DATABASE"),
        host:     getConfigVal("PG_HOST"),
        password: getConfigVal("PG_PASSWORD"),
        port:     getConfigVal("PG_PORT"),
	}

	if len(missing) > 0 {
		sort.Strings(missing)
		return cfg, fmt.Errorf("missing config variables %v", missing)
	}
	return cfg, nil
	
}
func (p pgconfig) getConfigString() string {
	return ""
}
