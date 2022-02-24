package main

import (
	"flag"
	"net/http"

	"github.com/justsharan/shorten/internal/middleware"
)

var (
	port = flag.String("port", "3000", "The port to listen for incoming requests")
	auth = flag.String("auth", "password", "The Authorization header to look for")
)

func main() {
	handler := middleware.Logger(handleIncoming)
	http.ListenAndServe(":" + *port, handler)
}

func handleIncoming(w http.ResponseWriter, r *http.Request) {
	// Handle incoming requests here
}
