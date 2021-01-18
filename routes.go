package main

import (
	"net/http"
	"strings"
)

func handleRoutes(w http.ResponseWriter, r *http.Request) {
	http.HandlerFunc(getRoute).ServeHTTP(w, r)
}

func getRoute(w http.ResponseWriter, r *http.Request) {
	split := strings.Split(r.URL.Path, "/")
	route := split[len(split)-1]
	if _, ok := routes[route]; ok {
		http.Redirect(w, r, routes[route], http.StatusTemporaryRedirect)
	} else {
		http.Error(w, http.StatusText(404), 404)
	}
}
