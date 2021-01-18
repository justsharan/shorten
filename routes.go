package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

var handleRoutes = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if r.URL.Path == "/" {
			w.Write([]byte("Hello World\n"))
		} else {
			getRoute(w, r)
		}
	case "POST":
		auth(postRoute).ServeHTTP(w, r)
	case "DELETE":
		auth(deleteRoute).ServeHTTP(w, r)
	}
})

func getRoute(w http.ResponseWriter, r *http.Request) {
	split := strings.Split(r.URL.Path, "/")
	route := split[len(split)-1]
	if _, ok := urls[route]; ok {
		http.Redirect(w, r, urls[route], http.StatusTemporaryRedirect)
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

	urls[split[len(split)-1]] = string(bytes)
	w.Write([]byte("https://" + r.Host + "/" + split[len(split)-1] + "\n"))
}

func deleteRoute(w http.ResponseWriter, r *http.Request) {
	split := strings.Split(r.URL.Path, "/")
	delete(urls, split[len(split)-1])
}
