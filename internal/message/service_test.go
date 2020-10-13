package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestServiceImpl_SendMessage(t *testing.T) {
	messageRepo := &messageRepoMock{}

	service := ServiceImpl{
		messageRepo: messageRepo,
	}

	cmd := SendMessage{
		Body: "this is a test message",
	}

	messageRepo.On("Save", mock.MatchedBy(func(req Message) bool {
		assert.Equal(t, cmd.Body, req.Body)
		return true
	})).Return(nil)

	err := service.SendMessage(cmd)
	assert.NoError(t, err)
}

func TestServiceImpl_GetMessages(t *testing.T) {
	messageRepo := &messageRepoMock{}

	service := ServiceImpl{
		messageRepo: messageRepo,
	}

	allMessages := []Message{
		{"id1", "message 1"},
		{"id2", "message 2"},
		{"id3", "message 3"},
	}

	messageRepo.On("FindAll").Return(allMessages, nil)

	messages, err := service.GetMessages()
	assert.NoError(t, err)
	assert.NotNil(t, messages)
}
