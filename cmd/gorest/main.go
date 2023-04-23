package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nunoki/gorest/internal/gorest"
	"github.com/nunoki/gorest/internal/gorest/postgres"
)

var (
	ctx = context.Background()
)

const defaultPostgresPort = "5432"
const defaultPostgresHost = "localhost"
const defaultPostgresSSLMode = "disable"

func main() {
	// connect to database
	dbU, dbPW, db, dbH, dbP, dbSSL := ConnectionStringFromEnv()
	connStr := postgres.ConnectionString(dbU, dbPW, db, dbH, dbP, dbSSL)
	pg, err := postgres.NewClient(ctx, connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		pg.Close()
	}()

	s := gorest.NewServer(pg, true)
	port := getPort()
	fmt.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, s))
}

// ConnectionStringFromEnv collects connection data from environment vars.
// It has some defaults for parameters that are not defined.
// It panics if required parameters are missing.
func ConnectionStringFromEnv() (string, string, string, string, string, string) {
	u := os.Getenv("POSTGRES_USER")
	pw := os.Getenv("POSTGRES_PASSWORD")
	db := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	sslmode := os.Getenv("POSTGRES_SSLMODE")

	if len(sslmode) == 0 {
		sslmode = defaultPostgresSSLMode
	}

	if len(port) == 0 {
		port = defaultPostgresPort
	}

	if len(host) == 0 {
		host = defaultPostgresHost
	}

	if len(u) == 0 || len(pw) == 0 || len(db) == 0 {
		panic("missing required parameter for database connection; required: POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB, optional: POSTGRES_HOST, POSTGRES_PORT")
	}

	return u, pw, db, host, port, sslmode
}

// getPort returns the port number that the app is specified to run on. It will try to read from
// the environment, or return "80" by default.
func getPort() string {
	port := os.Getenv("APP_PORT")
	if len(port) == 0 {
		port = "80"
	}
	return port
}
