package handler

import (
	"net/http"

	"github.com/aeramu/simple-message/internal/message"
	"github.com/gorilla/mux"
)

var messageCh chan message.Message

// NewHandler return handler implementation
func NewHandler(messageService message.Service) http.Handler {
	r := mux.NewRouter()

	messageHandler := &messageHandler{
		messageService: messageService,
	}

	messageCh = make(chan message.Message)

	// start messageCh receiver
	go broadcaster()

	r.HandleFunc("/messages", messageHandler.getMessages).Methods("GET")
	r.HandleFunc("/messages", messageHandler.postMessages).Methods("POST")
	r.HandleFunc("/live", live)

	return r
}
