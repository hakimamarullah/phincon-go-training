package main

import "net/http"

func main() {
	fs := http.FileServer(http.Dir("./static"))

	mux := http.NewServeMux()
	mux.Handle("/fileserver/", http.StripPrefix("/fileserver/", fs))

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
