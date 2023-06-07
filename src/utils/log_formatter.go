package utils

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/exp/slog"
)

func NewLogFormatter(l *slog.Logger) middleware.LogFormatter {
	return &LogFormatter{l}
}

type LogFormatter struct {
	Logger *slog.Logger
}

func (l *LogFormatter) NewLogEntry(r *http.Request) middleware.LogEntry {
	attrs := []slog.Attr{
		slog.String("method", r.Method),
		slog.String("url", r.URL.String()),
	}

	return &LogEntry{l.Logger, attrs}
}

type LogEntry struct {
	Logger *slog.Logger
	Attrs  []slog.Attr
}

func (l *LogEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra any) {
	l.Attrs = append(l.Attrs,
		slog.Int("status", status),
		slog.Int("size", bytes),
		slog.Duration("duration", elapsed),
	)

	l.Logger.LogAttrs(context.Background(), slog.LevelInfo, "request handled", l.Attrs...)
}

func (l *LogEntry) Panic(v any, stack []byte) {
	l.Attrs = append(l.Attrs,
		slog.String("stack", string(stack)),
		slog.String("panic", fmt.Sprintf("%+v", v)),
	)
}
