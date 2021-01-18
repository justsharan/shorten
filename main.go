package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
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

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("Server is starting...")

	router := http.NewServeMux()
	router.HandleFunc("/", handleRoutes)

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  14 * time.Second,
	}

	server.ListenAndServe()
}

func save() {
	var content string
	for k, v := range routes {
		content += k + " " + v
	}
	ioutil.WriteFile("routes.txt", []byte(content), 0466)
}
