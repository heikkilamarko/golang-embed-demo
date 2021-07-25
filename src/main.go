package main

import (
	"embed"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/heikkilamarko/goutils"
)

//go:embed ui
var uiFS embed.FS

func main() {
	spaHandler, err := goutils.NewSPAHandler(uiFS, "ui", "index.html")
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/api/message", handleMessage).Methods(http.MethodGet)
	router.PathPrefix("/").Handler(http.StripPrefix("/", spaHandler)).Methods(http.MethodGet)

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Addr:         ":8080",
		Handler:      router,
	}

	log.Printf("application is running at %s", server.Addr)

	log.Fatal(server.ListenAndServe())
}

func handleMessage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from API"))
}
