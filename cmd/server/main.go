package main

import (
	"flag"
	"net/http"
	"strconv"

	"github.com/justsharan/shorten/internal/middleware"
)

var (
	port = flag.Int("port", 3000, "The port to listen for incoming requests")
	auth = flag.String("auth", "password", "The Authorization header to look for")
	store = flag.String("store", "routes.txt", "File to store routes in")
)

func main() {
	handler := middleware.Logger(handleIncoming)
	http.ListenAndServe(":" + strconv.Itoa(*port), handler)
}

func handleIncoming(w http.ResponseWriter, r *http.Request) {
	// Handle incoming requests here
}
