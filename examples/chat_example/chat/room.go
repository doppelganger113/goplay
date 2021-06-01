package chat

import (
	"github.com/gorilla/websocket"
	"goplay.com/m/v2/examples/chat_example/trace"
	"log"
	"net/http"
)

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: messageBufferSize}

type room struct {
	forward chan []byte
	leave   chan *client
	join    chan *client
	clients map[*client]bool
	Tracer  trace.Tracer
}

func (r *room) Run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
			r.Tracer.Trace("New client joined")
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
			r.Tracer.Trace("Client left")
		case msg := <-r.forward:
			r.Tracer.Trace("Message received: ", string(msg))
			for client := range r.clients {
				client.send <- msg
				r.Tracer.Trace("--sent to client")
			}
		}
	}
}

func NewRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		Tracer:  trace.Off(),
	}
}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP: ", err)
		return
	}

	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
	}

	r.join <- client
	defer func() { r.leave <- client }()

	go client.write()
	client.read()
}
