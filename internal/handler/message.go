package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aeramu/simple-message/internal/message"
)

type messageHandler struct {
	messageService message.Service
}

func (h messageHandler) getMessages(w http.ResponseWriter, r *http.Request) {
	// get messages from service
	messages, err := h.messageService.GetMessages()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// marshal to json
	json, err := json.Marshal(messages)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(json)
}

func (h messageHandler) postMessages(w http.ResponseWriter, r *http.Request) {
	// read from body request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	// make send message request model
	var cmd message.SendMessage
	if err := json.Unmarshal(body, &cmd); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// make sure the request valid
	if (cmd == message.SendMessage{}) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// send message to service
	message, err := h.messageService.SendMessage(cmd)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// send message to bradcast to websocket client
	messageCh <- message

	w.WriteHeader(http.StatusOK)
}
