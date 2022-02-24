package routes

import (
	"bytes"
	"fmt"
	"net/http"
	"path"
)

func PostRoute(routes *map[string]string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		// Read request body
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
	
		// Set new value
		(*routes)[path.Base(r.URL.Path)] = buf.String()
		rw.WriteHeader(http.StatusCreated)
		fmt.Fprintf(rw, r.URL.String())
	}
}
