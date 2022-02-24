package routes

import (
	"net/http"
	"path"
)

func GetRoute(rw http.ResponseWriter, r *http.Request, routes *map[string]string) {
	if url, ok := (*routes)[path.Base(r.URL.Path)]; ok {
		http.Redirect(rw, r, url, http.StatusTemporaryRedirect)
	} else {
		http.Error(rw, http.StatusText(404), 404)
	}
}
