package config

import (
	"fmt"
	"log"
	"os"
	"sort"
)

type pgconfig struct {
	username string
	password string
	database string
	host string
	port string
	sslMode string
}


func PgConfigFromEnv() pgconfig {
    var missing []string
    var err error = nil

    get := func(key string) string {
        val := os.Getenv(key)
        if val == "" {
            missing = append(missing, key)
        }
        return val
    }
    cfg := pgconfig{
        username: get("PG_USER"),
        database: get("PG_DATABASE"),
        host:     get("PG_HOST"),
        password: get("PG_PASSWORD"),
        port:     get("PG_PORT"),
        sslMode:  os.Getenv("PG_SSLMODE"), // optional, so we don't add it to missing
    }

    err = checkSslMode(cfg)

    err = checkMissingVars(missing)

    if err != nil {
		log.Fatalf("Postgres configuration error %v", err)
	}
    return cfg
}
func checkSslMode(cfg pgconfig) (error) {
    switch cfg.sslMode {
    // valid sslmode
    case "", "disable", "allow", "require", "verify-ca", "verify-full":
        return nil
    default:
        return fmt.Errorf(`invalid sslmode "%s": expected one of "", "disable", "allow", "require", "verify-ca", or "verify-full"`, cfg.sslMode)
    }
}

func checkMissingVars(missing []string) error {
    if len(missing) > 0 {
        sort.Strings(missing) // sort for consistency in error message
        return fmt.Errorf("missing required environment variables: %v", missing)
    }
    return nil
}
//"postgres://username:password@localhost:5432/database_name?sslmode=mode"
func (pg pgconfig) String() string {
	s := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", pg.username, pg.password, pg.host, pg.port, pg.database)
	if pg.sslMode != "" {
		s += "?sslmode=" + pg.sslMode
	}
	return s
}
