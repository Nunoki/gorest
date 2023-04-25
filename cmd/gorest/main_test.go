package main

import (
	"io"
	"log"
	"os"
	"testing"
)

func TestGetPort(t *testing.T) {
	os.Unsetenv("APP_PORT")

	port := getPort()
	exp := "80"

	if port != exp {
		t.Fatalf("expected %s, received %s", exp, port)
	}

	os.Setenv("APP_PORT", "69")

	exp = "69"
	port = getPort()

	if port != exp {
		t.Fatalf("expected %s, received %s", exp, port)
	}
}

func TestGetPayloadSizeLimit(t *testing.T) {
	log.SetOutput(io.Discard)
	type test struct {
		inp string
		exp int64
	}
	var default_ int64 = 5000 // #default_payload_limit

	tests := []test{
		{"600", 600},
		{"10000", 10000},
		{"-5", default_},
		{"1", 1},
	}

	for _, test := range tests {
		os.Setenv("PAYLOAD_BYTE_LIMIT", test.inp)
		limit := getPayloadLimit()
		if limit != test.exp {
			t.Fatalf("expected %d, received %d", test.exp, limit)
		}
	}

	os.Unsetenv("PAYLOAD_BYTE_LIMIT")
	limit := getPayloadLimit()
	if limit != default_ {
		t.Fatalf("expected %d, received %d", default_, limit)
	}
}
