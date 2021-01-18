package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync/atomic"
	"time"
)

var (
	healthy int32
	routes  map[string]string
)

func init() {
	bytes, err := ioutil.ReadFile("routes.txt")
	if err != nil {
		log.Fatal(err)
	}

	routes = make(map[string]string)
	for _, line := range strings.Split(string(bytes), "\n") {
		res := strings.Split(line, " ")
		if len(res) > 1 {
			routes[res[0]] = res[1]
		}
	}
}

func main() {
	var port string
	flag.StringVar(&port, "port", "4646", "server port")
	flag.Parse()

	router := http.NewServeMux()
	router.HandleFunc("/", handleRoutes)

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      logger(router),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  14 * time.Second,
	}

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		log.Println("Server is shutting down")
		atomic.StoreInt32(&healthy, 0)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}

		close(done)
	}()

	log.Printf("Server is starting on port %s\n", port)
	atomic.StoreInt32(&healthy, 1)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	<-done
	save()
}

func save() {
	var content string
	for k, v := range routes {
		content += k + " " + v + "\n"
	}
	ioutil.WriteFile("routes.txt", []byte(content), 0466)
}
