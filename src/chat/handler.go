package chat

import (
	"net/http"

	"github.com/gorilla/websocket"
	"golang.org/x/exp/slog"
)

type WSHandler struct {
	hub      *Hub
	upgrader *websocket.Upgrader
	logger   *slog.Logger
}

func NewWSHandler(hub *Hub, logger *slog.Logger) *WSHandler {
	return &WSHandler{hub, &websocket.Upgrader{}, logger}
}

func (h *WSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.logger.Error(err.Error())
		return
	}

	(&client{
		hub:    h.hub,
		conn:   conn,
		send:   make(chan []byte, 256),
		logger: h.logger,
	}).run()
}
