package gorest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	router *chi.Mux
}

// DOCME
func NewServer(repo Repository, port string, byteLimit int64) *Server {
	r := chi.NewRouter()     // no auth middleware
	rAuth := chi.NewRouter() // with auth middleware

	// NOTE: at the time of writing this, the RequestSize middleware results in a 500 error instead
	// of 413
	r.Use(middleware.RequestSize(byteLimit))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	rAuth.Use(dummyAuthMiddleware())
	r.Mount("/", rAuth)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong\n"))
	})

	rAuth.Get("/user-id", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Your user ID: " + r.Context().Value(userID).(string)))
	})

	return &Server{
		router: r,
	}
}

// ServeHTTP satisfies the http.Handler interface
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
