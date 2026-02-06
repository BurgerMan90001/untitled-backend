package config

import (
	"fmt"
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



func pgConfigFromEnv() (pgconfig, error) {
    var missing []string

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
    switch cfg.sslMode {
    case "", "disable", "allow", "require", "verify-ca", "verify-full":
        // valid sslmode
    default:
        return cfg, fmt.Errorf(`invalid sslmode "%s": expected one of "", "disable", "allow", "require", "verify-ca", or "verify-full"`, cfg.sslMode)
    }

    if len(missing) > 0 {
        sort.Strings(missing) // sort for consistency in error message
        return cfg, fmt.Errorf("missing required environment variables: %v", missing)
    }
    return cfg, nil
}


//"postgres://username:password@localhost:5432/database_name?sslmode=mode"
func (pg pgconfig) string() string {
	s := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", pg.username, pg.password, pg.host, pg.port, pg.database)
	if pg.sslMode != "" {
		s += "?sslmode=" + pg.sslMode
	}
	return s
}
