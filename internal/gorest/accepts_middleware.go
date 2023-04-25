package gorest

import (
	"net/http"
	"strings"
)

// acceptsJSON validates that the client accepts a JSON content-type response
func acceptsJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accept := r.Header.Get("Accept")
		if !strings.Contains(accept, "application/json") && !strings.Contains(accept, "*/*") {
			http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
			return
		}

		next.ServeHTTP(w, r)
	})
}
