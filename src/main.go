package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
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

	addr := getAddr()
	log.Printf("App running at %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func handleMessage(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second) // Simulate some work for demo purposes.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from API"))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(indexHTML)
}

func getAddr() string {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	return fmt.Sprintf(":%s", port)
}
