package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"path"
	"time"

	"github.com/gorilla/mux"
)

//go:embed ui/index.html
var indexHTML []byte

//go:embed ui
var uiFS embed.FS

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/message", handleAPI).Methods(http.MethodGet)
	router.PathPrefix("/").HandlerFunc(handleSPA).Methods(http.MethodGet)

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

func handleAPI(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from API"))
}

func handleSPA(w http.ResponseWriter, r *http.Request) {
	path := path.Join("ui", r.URL.Path)

	if _, err := uiFS.ReadFile(path); err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write(indexHTML)
		return
	}

	fsys, _ := fs.Sub(uiFS, "ui")
	http.FileServer(http.FS(fsys)).ServeHTTP(w, r)
}
