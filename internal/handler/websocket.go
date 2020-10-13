package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients []*websocket.Conn

func broadcaster() {
	// listen to the channel
	for message := range messageCh {
		// write message to every clients
		json, err := json.Marshal(message)
		if err != nil {
			log.Println(err)
		}
		for _, client := range clients {
			client.WriteMessage(1, []byte(json))
		}
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func live(w http.ResponseWriter, r *http.Request) {
	// make connection to be websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	// add client to the clients list to be brodcasted
	clients = append(clients, ws)
	println(len(clients))

}
