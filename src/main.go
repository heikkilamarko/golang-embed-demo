package main

import (
	"context"
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/pkg/profile"
)

//go:embed static
var staticFS embed.FS

//go:embed index.html
var indexHTML []byte

func main() {
	defer profile.Start(profile.TraceProfile, profile.ProfilePath("."), profile.NoShutdownHook).Stop()

	addr := getAddr()

	var s http.Server

	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)

		<-sigint

		if err := s.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}

		close(idleConnsClosed)
	}()

	m := http.NewServeMux()

	fs := http.FileServer(http.FS(staticFS))

	m.Handle("/static/", fs)
	m.HandleFunc("/api/message", handleMessage)
	m.HandleFunc("/", handleIndex)

	s = http.Server{Addr: addr, Handler: m}

	log.Printf("HTTP server running at %s", addr)

	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
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
