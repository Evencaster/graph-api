package server

import (
	"net/http"
)

const allowedHeaders = "Accept, Authorization, Content-Type, Origin, X-Requested-With"

// CORS wraps an HTTP request and serves it with the correct CORS headers.
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		w.Header().Set(
			"Access-Control-Allow-Credentials",
			"true",
		)
		w.Header().Set(
			"Access-Control-Allow-Headers",
			allowedHeaders,
		)
		w.Header().Set(
			"Access-Control-Allow-Origin",
			origin,
		)
		w.Header().Set(
			"Access-Control-Request-Method",
			"GET, POST, OPTIONS, DELETE, PUT",
		)
		w.Header().Set(
			"Allow",
			"GET, POST, OPTIONS, DELETE, PUT",
		)

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
