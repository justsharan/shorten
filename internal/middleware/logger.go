package middleware

import (
	"log"
	"net/http"
)

// Logger logs incoming requests to the console
func Logger(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s\n", r.Method, r.URL.String())
		next.ServeHTTP(w, r)
	})
}
