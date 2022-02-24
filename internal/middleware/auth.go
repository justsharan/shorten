package middleware

import "net/http"

// Auth makes sure that the incoming request has the proper Authorization header
func Auth(next http.HandlerFunc, token string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != token {
			http.Error(w, http.StatusText(401), 401)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
