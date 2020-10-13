package main

import (
	"log"
	"net/http"

	"github.com/aeramu/simple-message/internal/handler"
	"github.com/aeramu/simple-message/internal/message"
	"github.com/aeramu/simple-message/internal/repository/dummy"
)

func main() {
	messageRepo := dummy.NewMessageRepo()
	service := message.NewService(messageRepo)
	handler := handler.NewHandler(service)
	log.Println(" Server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
