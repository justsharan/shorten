package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func handleRoutes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.HandlerFunc(getRoute).ServeHTTP(w, r)
	case "POST":
		http.HandlerFunc(postRoute).ServeHTTP(w, r)
	case "DELETE":
		http.HandlerFunc(deleteRoute).ServeHTTP(w, r)
	}
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

func postRoute(w http.ResponseWriter, r *http.Request) {
	split := strings.Split(r.URL.Path, "/")

	defer r.Body.Close()
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	routes[split[len(split)-1]] = string(bytes)
	w.Write([]byte("https://" + r.Host + "/" + split[len(split)-1] + "\n"))
}

func deleteRoute(w http.ResponseWriter, r *http.Request) {
	split := strings.Split(r.URL.Path, "/")
	delete(routes, split[len(split)-1])
}
