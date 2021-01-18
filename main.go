package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var urls map[string]string

func init() {
	urls = make(map[string]string)
	if err := readFile("routes.txt", &urls); err != nil {
		log.Fatal(err)
	}
}

func main() {
	var port string
	flag.StringVar(&port, "port", "4646", "server port")
	flag.Parse()

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      logger(handleRoutes),
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

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}

		close(done)
	}()

	log.Printf("Server is starting on port %s\n", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	<-done
	saveFile("routes.txt", &urls)
}
