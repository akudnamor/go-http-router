package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello world!\n")
	return
}

func main() {
	router := httprouter.New()
	router.GET("/hello", Index)

	log.Fatal(http.ListenAndServe(":8080", router))
}
