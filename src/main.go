package main

import (
	"embed"
	"golang-embed-demo/chat"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/heikkilamarko/goutils"
)

//go:embed ui
var uiFS embed.FS

func main() {
	spaHandler, err := goutils.NewSPAHandler(uiFS, "ui", "index.html", prepareSPAResponse)
	if err != nil {
		log.Fatal(err)
	}

	chatHub := chat.NewHub()
	go chatHub.Run()

	router := mux.NewRouter()

	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		chat.ServeWs(chatHub, w, r)
	})

	router.HandleFunc("/api/message", handleMessage).Methods(http.MethodGet)

	router.PathPrefix("/").Handler(http.StripPrefix("/", spaHandler)).Methods(http.MethodGet)

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Addr:         ":" + env("PORT", "8080"),
		Handler:      router,
	}

	log.Printf("application is running at %s", server.Addr)

	log.Fatal(server.ListenAndServe())
}

func handleMessage(w http.ResponseWriter, _ *http.Request) {
	time.Sleep(2 * time.Second) // simulate a slow response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from API"))
}

func prepareSPAResponse(w http.ResponseWriter, r *http.Request, isIndex bool) {
	if isIndex {
		w.Header().Set("Cache-Control", "no-store, max-age=0")
	} else if strings.HasPrefix(r.URL.Path, "assets/") {
		w.Header().Set("Cache-Control", "max-age=31536000")
	}
}

func env(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return value
}
