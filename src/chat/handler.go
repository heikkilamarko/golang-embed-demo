package chat

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog"
)

type WSHandler struct {
	hub      *Hub
	upgrader *websocket.Upgrader
	logger   *zerolog.Logger
}

func NewWSHandler(hub *Hub, logger *zerolog.Logger) *WSHandler {
	return &WSHandler{hub, &websocket.Upgrader{}, logger}
}

func (h *WSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.logger.Error().Err(err).Send()
		return
	}

	(&client{
		hub:    h.hub,
		conn:   conn,
		send:   make(chan []byte, 256),
		logger: h.logger,
	}).run()
}
