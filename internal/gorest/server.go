package gorest

import (
	"net/http"
)

const DefaultPayloadLimit = 1000 // #default_payload_limit

// DOCME
func NewServer(repo Repository, port string) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	return server
}

// // ServeHTTP just wraps Gin's ServeHTTP
// func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	s.router.ServeHTTP(w, r)
// }

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
