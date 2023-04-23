package gorest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const DefaultPayloadLimit = 1000 // #default_payload_limit

type Server struct {
	router *chi.Mux
}

// DOCME
func NewServer(repo Repository, port string) *Server {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	return &Server{
		router: r,
	}
}

// DOCME
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// // Returns the byte limit for the payload, which should be passed as an environment variable
// // PAYLOAD_BYTE_LIMIT; if it isn't, then a default limit of DefaultPayloadLimit will be returned
// func getPayloadSizeLimit() int64 {
// 	ls := os.Getenv("PAYLOAD_BYTE_LIMIT")
// 	limit, err := strconv.Atoi(ls)
// 	if err != nil || limit <= 0 {
// 		limit = DefaultPayloadLimit
// 		log.Printf(
// 			"Payload limit set to default of %d bytes (use environment variable PAYLOAD_BYTE_LIMIT to override)",
// 			limit,
// 		)
// 	} else {
// 		log.Printf("Payload limit set to %d bytes", limit)
// 	}

// 	return int64(limit)
// }
