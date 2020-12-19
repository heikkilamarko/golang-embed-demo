package main

import (
	"embed"
	"log"
	"net/http"
	"time"
)

func main() {
	//go:embed static
	//go:embed index.html
	var staticFS embed.FS

	fs := http.FileServer(http.FS(staticFS))

	http.Handle("/", fs)
	http.HandleFunc("/api/message", handleMessage)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handleMessage(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from API"))
}
