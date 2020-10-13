package handler

import (
	"net/http"

	"github.com/aeramu/simple-message/internal/message"
	"github.com/gorilla/mux"
)

// NewHandler return handler implementation
func NewHandler(messageService message.Service) http.Handler {
	r := mux.NewRouter()

	messageHandler := &messageHandler{
		messageService: messageService,
	}

	r.HandleFunc("/messages", messageHandler.getMessages).Methods("GET")
	r.HandleFunc("/messages", messageHandler.postMessages).Methods("POST")

	return r
}
