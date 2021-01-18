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

var httpLogger *log.Logger

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

	httpLogger = log.New(os.Stdout, "http: ", log.LstdFlags)
	httpLogger.Printf("Server is starting on port %s\n", port)

	router := http.NewServeMux()
	router.HandleFunc("/", handleRoutes)

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      logger(router),
		ErrorLog:     httpLogger,
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
