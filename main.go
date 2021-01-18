package main

import (
	"flag"
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
	defer save()

	var port string
	flag.StringVar(&port, "port", "4646", "server port")
	flag.Parse()

	http.HandleFunc("/", handleRoutes)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func save() {
	var content string
	for k, v := range routes {
		content += k + " " + v
	}
	ioutil.WriteFile("routes.txt", []byte(content), 0466)
}
