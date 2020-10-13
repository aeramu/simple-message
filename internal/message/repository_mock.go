package message

import "github.com/stretchr/testify/mock"

type (
	messageRepoMock struct {
		mock.Mock
	}
)

func (r *messageRepoMock) Save(message Message) error {
	args := r.Called(message)
	return args.Error(0)
}

func (r *messageRepoMock) FindAll() ([]Message, error) {
	args := r.Called()
	return args.Get(0).([]Message), args.Error(1)
}
