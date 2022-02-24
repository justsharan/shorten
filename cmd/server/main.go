package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"

	"github.com/justsharan/shorten/internal/middleware"
	"github.com/justsharan/shorten/internal/routes"
)

var (
	port = flag.Int("port", 3000, "The port to listen for incoming requests")
	token = flag.String("token", "password", "The Authorization header to look for")
	store = flag.String("store", "routes.txt", "File to store routes in")
	routeStore map[string]string
)

func init() {
	flag.Parse()
	var err error
	routeStore, err = readStore(*store)
	if err != nil {
		fmt.Printf("Error reading store: %v", err)
		return
	}
}

func main() {
	handler := middleware.Logger(handleIncoming)
	http.ListenAndServe(":" + strconv.Itoa(*port), handler)
}

func handleIncoming(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		routes.GetRoute(rw, r, &routeStore)
	case "POST":
		middleware.Auth(routes.PostRoute(&routeStore), *token).ServeHTTP(rw, r)
	case "DELETE":
		middleware.Auth(routes.DeleteRoute(&routeStore), *token).ServeHTTP(rw, r)
	}
}
