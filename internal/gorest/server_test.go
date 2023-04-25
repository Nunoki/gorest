package gorest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type MyWriter struct {
	Content []byte
}

func (w *MyWriter) Write(p []byte) (n int, err error) {
	w.Content = p
	return len(p), nil
}

func TestChiMuxAttached(t *testing.T) {
	repo := RepositoryMock{}
	s := NewServer(&repo, "69", 0, false)

	if fmt.Sprintf("%T", s.router) != "*chi.Mux" {
		t.Fatal("chi router not attached to server")
	}
}

func TestAuthMiddlewareAttached(t *testing.T) {
	repo := RepositoryMock{}
	repo.FindFunc = func(s string) ([]byte, time.Time, error) {
		return []byte(""), time.Now(), nil
	}
	s := NewServer(&repo, "69", 0, false)

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("Accept", "application/json")
	rec := httptest.NewRecorder()
	s.ServeHTTP(rec, req)

	if rec.Result().StatusCode != 401 {
		t.Fatalf("expected response code 401, received %d", rec.Result().StatusCode)
	}

	req.Header.Set("Authorization", "Bearer debug")
	rec = httptest.NewRecorder()
	s.ServeHTTP(rec, req)

	if rec.Result().StatusCode != 200 {
		t.Fatalf("expected response code 200, received %d", rec.Result().StatusCode)
	}
}

func TestPayloadLimiter(t *testing.T) {
	repo := RepositoryMock{}
	repo.UpdateFunc = func(uid string, content []byte) error {
		return nil
	}

	s := NewServer(&repo, "69", 5, false)

	req, _ := http.NewRequest("PUT", "/", strings.NewReader(`123456`))
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer debug")

	rec := httptest.NewRecorder()
	s.ServeHTTP(rec, req)

	// NOTE: the built-in RequestSize middleware for some reason results in a 500 error instead of a
	// proper 413
	if rec.Result().StatusCode != 500 {
		t.Fatalf("expected response code 500, received %d", rec.Result().StatusCode)
	}

	req.Body = ioutil.NopCloser(strings.NewReader(`12345`))

	rec = httptest.NewRecorder()
	s.ServeHTTP(rec, req)

	if rec.Result().StatusCode != 200 {
		t.Fatalf("expected response code 200, received %d", rec.Result().StatusCode)
	}
}
