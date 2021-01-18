package main

import (
	"log"
	"net/http"
)

func auth(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != token {
			http.Error(w, http.StatusText(401), 401)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		log.Printf("%s %s\n", r.Method, r.URL.String())
	})
}
