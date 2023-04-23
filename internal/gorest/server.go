package gorest

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const DefaultPayloadLimit = 1000 // #default_payload_limit

// DOCME
func NewServer(repo Repository, port string, payloadLimit int64) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Write([]byte("pong"))
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		handleIndex(w, r)
	})

	mux.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		handlePost(w, r)
	})

	handler := loggingMiddleware(mux)

	if payloadLimit > 0 {
		log.Printf(
			"Payload limit is %d bytes",
			payloadLimit,
		)
		handler = payloadLimitMiddleware(handler, payloadLimit)
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	return server
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "The posted data was: %s", body)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received %s request for %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func payloadLimitMiddleware(next http.Handler, limit int64) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ContentLength > limit {
			log.Printf(
				"%d error with payload size of %d bytes by IP %s",
				http.StatusRequestEntityTooLarge,
				r.ContentLength,
				r.RemoteAddr,
			)
			http.Error(w, "Payload too large", http.StatusRequestEntityTooLarge)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// // ServeHTTP just wraps Gin's ServeHTTP
// func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	s.router.ServeHTTP(w, r)
// }
