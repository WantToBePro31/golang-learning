package main

import (
	"fmt"
	"net/http"
	"time"
)

func GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(fmt.Sprintf("%v, %v %v %v", time.Now().Weekday(), time.Now().Day(), time.Now().Month(), time.Now().Year())))
	} // TODO: replace this
}

func main() {
	http.ListenAndServe("localhost:8080", GetHandler())
}
