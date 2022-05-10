package main

import (
	"context"
	"log"
	"net/http"

	"github.com/nunoki/demo-go-microservice/internal/beetroot"
	"github.com/nunoki/demo-go-microservice/internal/beetroot/postgres"
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

	s := beetroot.NewServer(pg)
	log.Fatal(http.ListenAndServe(":80", s))
}
