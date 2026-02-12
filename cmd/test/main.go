package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"github.com/BurgerMan90001/untitled-backend/internal/config" // embedded postgres server.
	_ "github.com/lib/pq"
)


func main() {
    config.LoadEnv()
    timeout := flag.Duration("timeout", 5*time.Second, "timeout for connecting to postgres")
    flag.Parse()

    cfg, err := pgConfigFromEnv() // defined below
    if err != nil {
        log.Fatalf("postgres configuration error: %v", err)
    }


    log.Printf("postgres is running on: %s\n", cfg.String())

    // ---- connect to postgres ----

    db, err := sql.Open(
        "postgres",
        cfg.String(), // defined below
    )
    if err != nil {
        panic(err)
    }
    defer db.Close() // always close the database when you're done with it.

    // always ping the database to ensure a connection is made.
    // any time you talk to a DB, use a context with a timeout, since DB connections could be lost or delayed indefinitely.
    ctx, cancel := context.WithTimeout(context.Background(), *timeout)
    defer cancel()
    if err := db.PingContext(ctx); err != nil {
        panic(err)
    }
    log.Println("ping successful")

}

// pgconfig is a struct that holds the configuration for connecting to a postgres database.
// each field corresponds to a component of the connection string.
// the following required environment variables are used to populate the struct:
//
//    PG_USER
//     PG_PASSWORD
//     PG_HOST
//     PG_PORT
//     PG_DATABASE
//
// additionally, the following optional environment variable is used to populate the sslmode:
//
//    PG_SSLMODE: must be one of  "", "disable", "allow", "require", "verify-ca", or "verify-full"
type pgconfig struct {
    user, database, host, password, port string // required
    sslMode                              string // optional
}

func pgConfigFromEnv() (pgconfig, error) {
    var missing []string
    // small closures like this can help reduce code duplication and make intent clearer.
    // you generally pay a small performance penalty for this, but for configuration, it's not a big deal;
    // you can spare the nanoseconds.
    // i prefer little helper functions like this to a complicated configuration framework like viper, cobra, envconfig, etc.
    get := func(key string) string {
        val := os.Getenv(key)
        if val == "" {
            missing = append(missing, key)
        }
        return val
    }
    cfg := pgconfig{
        user:     get("POSTGRES_USER"),
        database: get("POSTGRES_DATABASE"),
        host:     get("POSTGRES_HOST"),
        password: get("POSTGRES_PASSWORD"),
        port:     get("POSTGRES_PORT"),
        sslMode:  os.Getenv("POSTGRES_SSLMODE"), // optional, so we don't add it to missing
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

// String returns the connection string for the given pgconfig.
func (pg pgconfig) String() string {
    s := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", pg.user, pg.password, pg.host, pg.port, pg.database)
    if pg.sslMode != "" {
        s += "?sslmode=" + pg.sslMode
    }
    return s
}
