package main

import (
	"embed"
	"golang-embed-demo/chat"
	"golang-embed-demo/utils"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/heikkilamarko/goutils"
)

//go:embed ui/*
var uiFS embed.FS

func main() {
	logger := createLogger()

	logFormatter := utils.NewLogFormatter(logger)

	spaHandler, err := goutils.NewSPAHandler(uiFS, "ui", "app.html", prepareSPAResponse)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
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

	logger.Info("application is running at " + server.Addr)

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

func handleMessage(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from API"))
}

func prepareSPAResponse(w http.ResponseWriter, r *http.Request, isIndex bool) {
	if isIndex {
		w.Header().Set("Cache-Control", "no-store, max-age=0")
	}
}

func createLogger() *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	handler := slog.NewJSONHandler(os.Stderr, opts)

	logger := slog.New(handler)

	slog.SetDefault(logger)

	return logger
}

func env(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return value
}
