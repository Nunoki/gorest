package gorest

import (
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

type MyWriter struct {
	Content []byte
}

func (w *MyWriter) Write(p []byte) (n int, err error) {
	w.Content = p
	return len(p), nil
}

func TestGinEngineAttached(t *testing.T) {
	repo := RepositoryMock{}
	s := NewServer(&repo)

	if fmt.Sprintf("%T", s.router) != "*gin.Engine" {
		t.Fatal("Gin engine not attached to server")
	}
}

/*
func TestJWTAuthMiddlewareAttached(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	repo := RepositoryMock{}
	gin.SetMode(gin.TestMode)

	router := NewServer(&repo)

	token := "invalid.token.string"

	writer := httptest.NewRecorder()
	resp, _ := http.NewRequest("GET", "/", strings.NewReader("123"))
	resp.Header.Set("Authorization", token)
	router.ServeHTTP(writer, resp)

	if writer.Code != http.StatusUnauthorized {
		t.Fatalf("Expected code %d, received %d", http.StatusUnauthorized, writer.Code)
	}
}

func TestGetJWTPublicKey(t *testing.T) {
	testKey := "hello i am key"
	os.Setenv("JWT_PUBLIC_KEY", testKey)
	key := getJWTPublicKey()
	if key != testKey {
		t.Fatalf("Incorrect key returned; received %q, expected %q", key, testKey)
	}

	w := &MyWriter{}
	log.SetOutput(w)
	os.Unsetenv("JWT_PUBLIC_KEY")
	key = getJWTPublicKey()
	if key != "" {
		t.Fatal("Expected key to be empty, received:", key)
	}
	if !strings.Contains(string(w.Content), "JWT_PUBLIC_KEY") {
		t.Fatal("Warning should have been logged about missing public key declaration")
	}
}
*/

func TestGetPayloadSizeLimit(t *testing.T) {
	log.SetOutput(io.Discard)
	type test struct {
		inp string
		exp int64
	}
	var def int64 = 1000 // #default_payload_limit
	tests := []test{
		{"600", 600},
		{"10000", 10000},
		{"-5", def},
		{"600", 600},
		{"600", 600},
	}

	for _, test := range tests {
		os.Setenv("PAYLOAD_BYTE_LIMIT", test.inp)
		limit := getPayloadSizeLimit()
		if limit != test.exp {
			t.Fatalf("Got %d, expected %d", limit, test.exp)
		}
	}

	os.Unsetenv("PAYLOAD_BYTE_LIMIT")
	limit := getPayloadSizeLimit()
	if limit != def {
		t.Fatalf("Got %d, expected %d", limit, def)
	}
}

// TODO: How can i circumvent the auth middleware in a test to only test if the size limiter gets
// its default size set correctly
// func TestSizeLimiterGetsDefaultSize(t *testing.T) {
// 	server := NewServer()
// }
