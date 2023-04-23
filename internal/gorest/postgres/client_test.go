package postgres

import (
	"testing"
)

func TestConnectionString(t *testing.T) {
	connStr := ConnectionString("alice", "god", "loldb", "notlocalhost", "420", "verify-ca")
	expected := "postgres://alice:god@notlocalhost:420/loldb?sslmode=verify-ca"

	if connStr != expected {
		t.Fatalf("expected %s, received %s", expected, connStr)
	}
}
