package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func World(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", Hello)
	mux.HandleFunc("/world", World)

	server := http.Server{
		Addr:    "localhost:5000",
		Handler: mux,
	}

	fmt.Printf("Server is running on %s", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Print(err.Error())
	}
}
