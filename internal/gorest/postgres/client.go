package postgres

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

const defaultPostgresPort = "5432"
const defaultPostgresHost = "localhost"

type Client struct {
	conn *pgxpool.Pool
}

// NewClient is a constructor for a new Repository instance with conn as the database connection
func NewClient(ctx context.Context, connStr string) (*Client, error) {
	conn, err := pgxpool.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}
	return &Client{
		conn: conn,
	}, nil
}

// Close calls .Close() on the underlying pgxpool.Pool instance
func (c *Client) Close() {
	if c == nil || c.conn == nil {
		return
	}
	c.conn.Close()
}

// ConnectionStringFromEnv parses connection string from environment
// and panics if required paramers are missing
func ConnectionStringFromEnv() string {
	u := os.Getenv("POSTGRES_USER")
	pw := os.Getenv("POSTGRES_PASSWORD")
	db := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")
	host := os.Getenv("POSTGRES_HOST")
	sslmode := os.Getenv("POSTGRES_SSLMODE")

	if len(port) == 0 {
		port = defaultPostgresPort
	}

	if len(host) == 0 {
		host = defaultPostgresHost
	}

	if len(u) == 0 || len(pw) == 0 || len(db) == 0 {
		panic("missing required parameter for database connection; required: POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB, optional: POSTGRES_HOST, POSTGRES_PORT")
	}

	return ConnectionString(u, pw, db, host, port, sslmode)
}

// ConnectionString returns the postgres connection string with parameters defined in the
// call arguments
func ConnectionString(username, password, db, host, port, sslmode string) string {
	username = url.QueryEscape(username)
	password = url.QueryEscape(password)
	db = url.QueryEscape(db)

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		username,
		password,
		host,
		port,
		db,
		sslmode,
	)
}
