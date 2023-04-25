package gorest

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func TestDummyAuthMiddlewareOK(t *testing.T) {
	r := chi.NewRouter()
	r.Use(dummyAuthMiddleware())

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.Status(r, http.StatusOK)
		render.PlainText(w, r, "ok")
	})

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer debug")
	r.ServeHTTP(resp, req)

	res := resp.Result()
	if res.StatusCode != 200 {
		t.Fatalf("expected code 200, received %d", res.StatusCode)
	}
	content, _ := io.ReadAll(res.Body)
	if string(content) != "ok" {
		t.Fatalf("expected body to equal \"ok\", received %s", string(content))
	}
}

func TestDummyAuthMiddlewareForbidden(t *testing.T) {
	r := chi.NewRouter()
	r.Use(dummyAuthMiddleware())

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.Status(r, http.StatusOK)
		render.PlainText(w, r, "ok")
	})

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(resp, req)

	res := resp.Result()
	if res.StatusCode != 401 {
		t.Fatalf("expected code 401, received %d", res.StatusCode)
	}
}
