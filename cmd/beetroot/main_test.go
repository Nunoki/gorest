package main

import (
	"os"
	"testing"
)

func TestGetPort(t *testing.T) {
	os.Unsetenv("PORT")

	port := getPort()
	exp := "80"

	if port != exp {
		t.Fatalf("expected %s, received %s", exp, port)
	}

	os.Setenv("PORT", "69")

	exp = "69"
	port = getPort()

	if port != exp {
		t.Fatalf("expected %s, received %s", exp, port)
	}
}
