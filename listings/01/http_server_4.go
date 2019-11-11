package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting HTTP server...")
	log.Fatal(http.ListenAndServe("localhost:4000", &helloHandler{}))
}

type helloHandler struct{}

func (*helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}
