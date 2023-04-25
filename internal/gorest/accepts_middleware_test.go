package gorest

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAcceptMiddlewareSucceedsWithJSON(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	testServer := httptest.NewServer(acceptsJSON(handler))
	defer testServer.Close()

	req, err := http.NewRequest("GET", testServer.URL, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Add("Accept", "foo,application/json,bar")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to make request to test server: %v", err)
	}
	defer resp.Body.Close()

	exp := http.StatusOK
	if resp.StatusCode != exp {
		t.Errorf("expected status code %d, received %d", exp, resp.StatusCode)
	}
}

func TestAcceptMiddlewareSucceedsWithAsterisk(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	testServer := httptest.NewServer(acceptsJSON(handler))
	defer testServer.Close()

	req, err := http.NewRequest("GET", testServer.URL, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Add("Accept", "foo,*/*,bar")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to make request to test server: %v", err)
	}
	defer resp.Body.Close()

	exp := http.StatusOK
	if resp.StatusCode != exp {
		t.Errorf("expected status code %d, received %d", exp, resp.StatusCode)
	}
}

func TestAcceptMiddlewareFailsWhenNone(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	testServer := httptest.NewServer(acceptsJSON(handler))
	defer testServer.Close()

	req, err := http.NewRequest("GET", testServer.URL, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to make request to test server: %v", err)
	}
	defer resp.Body.Close()

	exp := http.StatusNotAcceptable
	if resp.StatusCode != exp {
		t.Errorf("expected status code %d, received %d", exp, resp.StatusCode)
	}
}
