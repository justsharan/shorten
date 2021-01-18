package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strings"
)

var routes = make(map[string]string)

func init() {
	file, err := os.Open("routes.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res := strings.Split(scanner.Text(), " ")
		routes[res[0]] = res[1]
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":4646", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
