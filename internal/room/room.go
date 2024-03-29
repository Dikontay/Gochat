package room

import (
	"github.com/gorilla/websocket"
	"gochat/internal/client"
	"log"
	"net/http"
)

const socketBufferSize = 1024
const messageBufferSize = 256

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: messageBufferSize}

type Room struct {
	Clients map[*client.Client]bool
	Join    chan *client.Client
	Leave   chan *client.Client
	Forward chan []byte
}

func (r *Room) Run() {
	for {
		select {
		case Client := <-r.Join:
			r.Clients[Client] = true
		case Client := <-r.Leave:
			delete(r.Clients, Client)
		case msg := <-r.Forward:
			for cli := range r.Clients {
				cli.Receive <- msg
			}
		}
	}
}

func (room *Room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeTTP:", err)
		return
	}
	cli := &client.Client{
		Receive: make(chan []byte, messageBufferSize),
		Conn:    socket,
		Room:    room,
	}
	room.Join <- cli
	defer func() {
		room.Leave <- cli
	}()

	go cli.Write()
	cli.Read()

}
