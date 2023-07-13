package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))

	mux := http.NewServeMux()
	mux.Handle("/fileserver/", http.StripPrefix("/fileserver/", fs))
	mux.HandleFunc("/download", downloadHandler)
	mux.HandleFunc("/upload", uploadHandler)

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	path := "./static/example.html"
	file, err := os.Open(path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Can't download file")
		return
	}

	fileInfo, _ := file.Stat()
	w.Header().Set("Content-Disposition", "attachment; filename="+fileInfo.Name())
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))

	_, err = io.Copy(w, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Can't download file")
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	saveFile, err := os.Create("./static/" + header.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer saveFile.Close()

	_, err = io.Copy(saveFile, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "File Uploaded")

}
