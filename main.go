package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var routes = make(map[string]string)

func init() {
	bytes, err := ioutil.ReadFile("routes.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range strings.Split(string(bytes), "\n") {
		res := strings.Split(line, " ")
		routes[res[0]] = res[1]
	}
}

func main() {
	http.HandleFunc("/", route)
	log.Fatal(http.ListenAndServe(":4646", nil))
}

func route(w http.ResponseWriter, r *http.Request) {
	split := strings.Split(r.URL.Path, "/")
	route := split[len(split)-1]
	if _, ok := routes[route]; ok {
		http.Redirect(w, r, routes[route], http.StatusTemporaryRedirect)
	} else {
		http.Error(w, http.StatusText(404), 404)
	}
}
