package main

import (
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `Hello world!`)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", HealthCheckHandler)

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
