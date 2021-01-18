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
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":4646", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
