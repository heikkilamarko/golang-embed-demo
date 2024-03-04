package chat

import (
	"log/slog"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeTimeout = 10 * time.Second
	readTimeout  = 60 * time.Second
	pingInterval = (readTimeout * 9) / 10
	readLimit    = 1024
)

type client struct {
	hub    *Hub
	conn   *websocket.Conn
	send   chan []byte
	logger *slog.Logger
}

func (c *client) run() {
	c.hub.register <- c
	go c.write()
	go c.read()
}

func (c *client) write() {
	ticker := time.NewTicker(pingInterval)

	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.setWriteDeadline()

			if !ok {
				c.writeCloseMessage()
				return
			}

			if err := c.writeTextMessage(message); err != nil {
				return
			}
		case <-ticker.C:
			c.setWriteDeadline()

			if err := c.writePingMessage(); err != nil {
				return
			}
		}
	}
}

func (c *client) read() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	c.conn.SetReadLimit(readLimit)
	c.conn.SetPongHandler(func(string) error {
		c.setReadDeadline()
		return nil
	})

	c.setReadDeadline()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			c.logger.Error(err.Error())
			break
		}

		c.hub.broadcast <- message
	}
}

func (c *client) setWriteDeadline() error {
	return c.conn.SetWriteDeadline(time.Now().Add(writeTimeout))
}

func (c *client) setReadDeadline() error {
	return c.conn.SetReadDeadline(time.Now().Add(readTimeout))
}

func (c *client) writeTextMessage(message []byte) error {
	return c.conn.WriteMessage(websocket.TextMessage, message)
}

func (c *client) writeCloseMessage() error {
	return c.conn.WriteMessage(websocket.CloseMessage, []byte{})
}

func (c *client) writePingMessage() error {
	return c.conn.WriteMessage(websocket.PingMessage, nil)
}
