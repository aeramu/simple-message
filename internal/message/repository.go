package message

// MessageRepo interface for message service
type MessageRepo interface {
	Save(message Message) error
	FindAll() ([]Message, error)
}
