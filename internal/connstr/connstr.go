package connstr

import "os"

const defaultPostgresPort = "5432"
const defaultPostgresHost = "localhost"
const defaultPostgresSSLMode = "disable"

// ConnectionStringFromEnv collects connection data from environment vars.
// It has some defaults for parameters that are not defined.
// It panics if required parameters are missing.
func FromEnv() (string, string, string, string, string, string) {
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
