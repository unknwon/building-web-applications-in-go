package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)

	log.Println("Starting HTTP server...")
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}
