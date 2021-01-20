package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ory/graceful"
	// "github.com/pkg/profile"
)

//go:embed static
var staticFS embed.FS

//go:embed index.html
var indexHTML []byte

func main() {
	// defer profile.Start(profile.TraceProfile, profile.ProfilePath("."), profile.NoShutdownHook).Stop()

	addr := getAddr()

	m := http.NewServeMux()

	fs := http.FileServer(http.FS(staticFS))

	m.Handle("/static/", fs)
	m.HandleFunc("/api/message", handleMessage)
	m.HandleFunc("/", handleIndex)

	server := graceful.WithDefaults(&http.Server{
		Addr:    addr,
		Handler: m})

	log.Printf("Application running at %s", addr)

	if err := graceful.Graceful(server.ListenAndServe, server.Shutdown); err != nil {
		log.Fatal(err)
	}

	log.Printf("Application shutdown gracefully")
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
