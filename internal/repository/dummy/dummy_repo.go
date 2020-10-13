package dummy

import (
	"github.com/aeramu/simple-message/internal/message"
)

// NewMessageRepo return message repo with dummy implementation
func NewMessageRepo() message.MessageRepo {
	return &messageRepo{
		data: []message.Message{},
	}
}

type messageRepo struct {
	data []message.Message
}

func (r *messageRepo) Save(message message.Message) error {
	r.data = append(r.data, message)
	return nil
}

func (r messageRepo) FindAll() ([]message.Message, error) {
	return r.data, nil
}
