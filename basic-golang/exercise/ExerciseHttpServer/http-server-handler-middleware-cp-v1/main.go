package main

import (
	"net/http"
)

func StudentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to Student page"))
	}
}

func RequestMethodGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		if method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
        	w.Write([]byte("Method is not allowed"))
		} else {
			next.ServeHTTP(w, r)
		}
	}) // TODO: replace this
}

func main() {
	// TODO: answer here
	mux := http.DefaultServeMux
	var handler http.Handler = mux
	handler = RequestMethodGet(handler)

	server := new(http.Server)
    server.Addr = ":8080"
    server.Handler = handler

    server.ListenAndServe()
}
