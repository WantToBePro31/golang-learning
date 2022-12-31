package main

import (
	"net/http"
)

func AdminHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to Admin page"))
	}
}

func RequestMethodGetMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method is not allowed"))
		} else {
			next.ServeHTTP(w, r)
		}
	}) // TODO: replace this
}

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("role") != "ADMIN" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Role not authorized"))
		} else {
			next.ServeHTTP(w, r)
		}
	}) // TODO: replace this
}

func main() {
	// TODO: answer here
	mux := http.DefaultServeMux
	var handler http.Handler = mux
	handler = RequestMethodGetMiddleware(handler)
	
	server := new(http.Server)
    server.Addr = ":8080"
    server.Handler = handler

    server.ListenAndServe()
}
