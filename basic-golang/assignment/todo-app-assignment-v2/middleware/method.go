package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"net/http"
)

func Get(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			var error model.ErrorResponse
			error.Error = "Method is not allowed!"
			jsonData, err := json.Marshal(error)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(jsonData)
		} else {
			next.ServeHTTP(w, r)
		}
	}) // TODO: replace this
}

func Post(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			var error model.ErrorResponse
			error.Error = "Method is not allowed!"
			jsonData, err := json.Marshal(error)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(jsonData)
		} else {
			next.ServeHTTP(w, r)
		}
	}) // TODO: replace this
}

func Delete(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			var error model.ErrorResponse
			error.Error = "Method is not allowed!"
			jsonData, err := json.Marshal(error)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(jsonData)
		} else {
			next.ServeHTTP(w, r)
		}
	}) // TODO: replace this
}
