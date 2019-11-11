package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &helloHandler{})
	mux.HandleFunc("/timeout", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		w.Write([]byte("Timeout"))
	})

	server := &http.Server{
		Addr:         ":4000",
		Handler:      mux,
		WriteTimeout: 2 * time.Second,
	}
	log.Println("Starting HTTP server...")
	log.Fatal(server.ListenAndServe())
}

type helloHandler struct{}

func (*helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}
