package chat

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WSHandler struct {
	hub      *Hub
	upgrader *websocket.Upgrader
}

func NewWSHandler(hub *Hub) *WSHandler {
	return &WSHandler{
		hub,
		&websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
}

func (h *WSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	(&client{
		hub:  h.hub,
		conn: conn,
		send: make(chan []byte, 256),
	}).run()
}
