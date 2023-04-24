package gorest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	router *chi.Mux
}

// DOCME
func NewServer(repo Repository, port string, byteLimit int64) *Server {
	r := chi.NewRouter()

	// NOTE: at the time of writing this, the RequestSize middleware results in a 500 error instead
	// of 413
	r.Use(middleware.RequestSize(byteLimit))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Post("/post", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		type response struct {
			Data string `json:"posted_data"`
		}

		w.Header().Set("Content-type", "application/json")
		resp := response{string(data)}
		json.NewEncoder(w).Encode(resp)
	})

	return &Server{
		router: r,
	}
}

// ServeHTTP satisfies the http.Handler interface
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
