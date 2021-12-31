package main

import (
	"embed"
	"golang-embed-demo/chat"
	"golang-embed-demo/utils"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/heikkilamarko/goutils"
	"github.com/rs/zerolog"
)

//go:embed ui
var uiFS embed.FS

func main() {
	logger := createLogger()

	logFormatter := utils.NewZerologLogFormatter(logger)

	spaHandler, err := goutils.NewSPAHandler(uiFS, "ui", "index.html", prepareSPAResponse)
	if err != nil {
		logger.Fatal().Err(err).Send()
	}

	chatHub := chat.NewHub()
	go chatHub.Run()

	chatHandler := chat.NewWSHandler(chatHub, logger)

	router := chi.NewRouter()

	router.Use(middleware.RequestLogger(logFormatter))
	router.Use(middleware.Recoverer)

	router.Method(http.MethodGet, "/ws", chatHandler)
	router.Get("/api/message", handleMessage)
	router.Method(http.MethodGet, "/*", http.StripPrefix("/", spaHandler))

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Addr:         ":" + env("PORT", "8080"),
		Handler:      router,
	}

	logger.Info().Msgf("application is running at %s", server.Addr)

	logger.Fatal().Err(server.ListenAndServe()).Send()
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

func createLogger() *zerolog.Logger {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	l := zerolog.
		New(os.Stderr).
		With().
		Timestamp().
		Logger()

	return &l
}

func env(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return value
}
