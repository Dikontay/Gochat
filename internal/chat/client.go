package chat

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn

	Receive chan []byte

	Room *Room
}

func (c *Client) Read() {
	defer c.Conn.Close()
	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			return
		}
		c.Room.Forward <- msg
	}

}

func (c *Client) Write() {
	defer c.Conn.Close()
	for msg := range c.Receive {
		err := c.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
