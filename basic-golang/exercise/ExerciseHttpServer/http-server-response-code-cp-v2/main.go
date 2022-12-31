package main

import (
	"fmt"
	"net/http"
)

var students = []string{
	"Aditira",
	"Dito",
	"Afis",
	"Eddy",
}

func IsNameExists(name string) bool {
	for _, n := range students {
		if n == name {
			return true
		}
	}

	return false
}

func CheckStudentName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		name := r.URL.Query().Get("name")
		fmt.Println(method)
		if method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("Method is not allowed"))
		} else {
			if IsNameExists(name) {
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "text/plain")
				w.Write([]byte("Name is exists"))
			} else {
				w.WriteHeader(http.StatusNotFound)
				w.Header().Set("Content-Type", "text/plain")
				w.Write([]byte("Data not found"))
			}
		}
	} // TODO: replace this
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	// TODO: answer here
	mux.HandleFunc("/students", CheckStudentName())
	return mux
}

func main() {
	http.ListenAndServe("localhost:8080", GetMux())
}
