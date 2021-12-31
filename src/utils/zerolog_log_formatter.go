package utils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
)

func NewZerologLogFormatter(l *zerolog.Logger) middleware.LogFormatter {
	return &ZerologLogFormatter{l}
}

type ZerologLogFormatter struct {
	Logger *zerolog.Logger
}

func (l *ZerologLogFormatter) NewLogEntry(r *http.Request) middleware.LogEntry {
	e := l.Logger.
		Info().
		Str("method", r.Method).
		Stringer("url", r.URL)

	return &ZerologLogEntry{e}
}

type ZerologLogEntry struct {
	Event *zerolog.Event
}

func (l *ZerologLogEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	l.Event.
		Int("status", status).
		Int("size", bytes).
		Dur("duration", elapsed).
		Send()
}

func (l *ZerologLogEntry) Panic(v interface{}, stack []byte) {
	l.Event.
		Str("stack", string(stack)).
		Str("panic", fmt.Sprintf("%+v", v))
}
