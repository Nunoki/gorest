package postgres

import (
	"context"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v4/pgxpool"
)

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
