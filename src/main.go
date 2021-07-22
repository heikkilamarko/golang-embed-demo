package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"
)

//go:embed ui
var uiFS embed.FS

func main() {
	http.Handle("/", handleUI())
	http.HandleFunc("/api/message", handleMessage)

	addr := getAddr()
	log.Printf("Application running at %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func handleUI() http.Handler {
	fsys, err := fs.Sub(uiFS, "ui")
	if err != nil {
		panic(err)
	}
	return http.FileServer(http.FS(fsys))
}

func handleMessage(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second) // Simulate some work for demo purposes.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from API"))
}

func getAddr() string {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	return fmt.Sprintf(":%s", port)
}
