package message

import (
	"log"

	"github.com/google/uuid"
)

// Service interface for message service
type Service interface {
	SendMessage(cmd SendMessage) (Message, error)
	GetMessages() ([]Message, error)
}

//NewService return message service implementation as service
func NewService(messageRepo MessageRepo) Service {
	return &ServiceImpl{
		messageRepo: messageRepo,
	}
}

// ServiceImpl implementation for Service
type ServiceImpl struct {
	messageRepo MessageRepo
}

// SendMessage send message service
func (s *ServiceImpl) SendMessage(cmd SendMessage) (Message, error) {
	messageID := uuid.New().String()
	message := Message{
		ID:   messageID,
		Body: cmd.Body,
	}

	if err := s.messageRepo.Save(message); err != nil {
		log.Println(err)
		return Message{}, err
	}

	return message, nil
}

// GetMessages get all message service
func (s *ServiceImpl) GetMessages() ([]Message, error) {
	messages, err := s.messageRepo.FindAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return messages, nil
}
