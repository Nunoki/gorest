package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/nunoki/gorest/internal/gorest"
	"github.com/nunoki/gorest/internal/gorest/postgres"
)

var (
	ctx = context.Background()
)

func main() {
	// connect to database
	connStr := postgres.ConnectionStringFromEnv()
	pg, err := postgres.NewClient(ctx, connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		pg.Close()
	}()

	s := gorest.NewServer(pg, true)
	port := getPort()
	println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, s))
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
