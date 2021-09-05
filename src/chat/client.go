package chat

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeDuration = 10 * time.Second
	readDuration  = 60 * time.Second
	readLimit     = 512
	pingPeriod    = (readDuration * 9) / 10
)

type client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

func (c *client) run() {
	c.hub.register <- c
	go c.write()
	go c.read()
}

func (c *client) write() {
	ticker := time.NewTicker(pingPeriod)

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

	c.setReadDeadline()

	c.conn.SetPongHandler(func(string) error {
		c.setReadDeadline()
		return nil
	})

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		c.hub.broadcast <- message
	}
}

func (c *client) setWriteDeadline() error {
	return c.conn.SetWriteDeadline(time.Now().Add(writeDuration))
}

func (c *client) setReadDeadline() error {
	return c.conn.SetReadDeadline(time.Now().Add(readDuration))
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
