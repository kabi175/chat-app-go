package mocks

import (
	"github.com/kabi175/chat-app-go/messager/domain"
	"github.com/stretchr/testify/mock"
)

type MessageRepository struct {
	mock.Mock
}

func (mr *MessageRepository) Publish(message *domain.Message) error {
	args := mr.Called(message)
	return args.Error(0)
}

func (mr *MessageRepository) Listern(user *domain.User) (*domain.Message, error) {
	args := mr.Called(user)
	r0 := args.Get(0)
	r1 := args.Error(1)
	if r0 == nil {
		return nil, r1
	}
	return r0.(*domain.Message), r1
}
