package gorest

import (
	"context"
	"fmt"
	"net/http"
)

type ctxKey string

const userID ctxKey = "userID"

// dummyAuthMiddleware is a blueprint for implementing auth middlewares.
// It requires a hardcoded value of "Bearer debug" in the Authorization header.
// It then sets a valid UUID v4 in the context to serve as the user ID.
// If the dummy bearer token is missing, it will write a 401 response.
func dummyAuthMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")

			if authHeader != "Bearer debug" {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, "Authorization needs header value of \"Bearer debug\"")
				return
			}

			ctx := context.WithValue(r.Context(), userID, "00000000-0000-0000-0000-000000000000")

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
