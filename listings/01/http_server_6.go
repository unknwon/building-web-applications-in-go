package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &helloHandler{})

	server := &http.Server{
		Addr:    ":4000",
		Handler: mux,
	}
	log.Println("Starting HTTP server...")
	log.Fatal(server.ListenAndServe())
}

type helloHandler struct{}

func (*helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}
