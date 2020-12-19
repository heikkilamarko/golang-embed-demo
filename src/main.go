package main

import (
	"embed"
	"log"
	"net/http"
	"time"
)

//go:embed static
var staticFS embed.FS

//go:embed index.html
var indexHTML []byte

func main() {
	fs := http.FileServer(http.FS(staticFS))

	http.Handle("/static/", fs)
	http.HandleFunc("/api/message", handleMessage)
	http.HandleFunc("/", handleIndex)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handleMessage(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from API"))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(indexHTML)
}
