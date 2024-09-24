package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
	return
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", HealthCheckHandler)

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
