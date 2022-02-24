package routes

import "net/http"

func DeleteRoute(routes *map[string]string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		delete(*routes, r.URL.Path)
		rw.WriteHeader(http.StatusNoContent)
	}
}
