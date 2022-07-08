package postgres

import (
	"os"
	"testing"
)

func TestConnectionStringFromEnv(t *testing.T) {
	os.Setenv("POSTGRES_USER", "bob")
	os.Setenv("POSTGRES_PASSWORD", "1234")
	os.Setenv("POSTGRES_DB", "pokemon")
	os.Unsetenv("POSTGRES_PORT")
	os.Unsetenv("POSTGRES_HOST")

	expected := "postgres://bob:1234@localhost:5432/pokemon?sslmode=disable"
	connStr := ConnectionStringFromEnv()

	if connStr != expected {
		t.Fatalf("expected %s, received %s", expected, connStr)
	}

	os.Setenv("POSTGRES_HOST", "1.2.3.4")
	os.Setenv("POSTGRES_PORT", "69")

	expected = "postgres://bob:1234@1.2.3.4:69/pokemon?sslmode=disable"
	connStr = ConnectionStringFromEnv()

	if connStr != expected {
		t.Fatalf("expected %s, received %s", expected, connStr)
	}
}

func TestConnectionString(t *testing.T) {
	connStr := ConnectionString("alice", "god", "loldb", "notlocalhost", "420")
	expected := "postgres://alice:god@notlocalhost:420/loldb?sslmode=disable"

	if connStr != expected {
		t.Fatalf("expected %s, received %s", expected, connStr)
	}
}
