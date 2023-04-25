package gorest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
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
