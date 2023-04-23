package connstr

import (
	"os"
	"reflect"
	"testing"
)

func TestConnectionStringFromEnv(t *testing.T) {
	os.Setenv("POSTGRES_USER", "bob")
	os.Setenv("POSTGRES_PASSWORD", "1234")
	os.Setenv("POSTGRES_DB", "pokemon")
	os.Unsetenv("POSTGRES_PORT")
	os.Unsetenv("POSTGRES_HOST")
	os.Unsetenv("POSTGRES_SSLMODE")

	type connectionData struct {
		username string
		password string
		database string
		host     string
		port     string
		ssl      string
	}

	exp := connectionData{"bob", "1234", "pokemon", "localhost", "5432", "disable"}
	dbU, dbPW, db, dbH, dbP, dbSSL := FromEnv()
	res := connectionData{dbU, dbPW, db, dbH, dbP, dbSSL}

	if !reflect.DeepEqual(exp, res) {
		t.Fatalf("expected %s, received %s", exp, res)
	}

	os.Setenv("POSTGRES_HOST", "1.2.3.4")
	os.Setenv("POSTGRES_PORT", "69")
	os.Setenv("POSTGRES_SSLMODE", "require")

	exp = connectionData{"bob", "1234", "pokemon", "1.2.3.4", "69", "require"}
	dbU, dbPW, db, dbH, dbP, dbSSL = FromEnv()
	res = connectionData{dbU, dbPW, db, dbH, dbP, dbSSL}

	if !reflect.DeepEqual(exp, res) {
		t.Fatalf("expected %s, received %s", exp, res)
	}
}
