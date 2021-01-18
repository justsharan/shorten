package main

import (
	"net/http"
)

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		httpLogger.Printf("%s %s\n", r.Method, r.URL.String())
	})
}
