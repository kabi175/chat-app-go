package mocks

import (
	"github.com/kabi175/chat-app-go/messager/domain"
	"github.com/stretchr/testify/mock"
)

type MessageService struct {
	mock.Mock
}

func (ms *MessageService) Publish(message *domain.Message) error {
	args := ms.Called(message)
	return args.Error(0)
}

func (ms *MessageService) Listern(user *domain.User) (*domain.Message, error) {
	args := ms.Called(user)
	r0 := args.Get(0)
	r1 := args.Error(1)
	if r0 == nil {
		return nil, r1
	}
	return r0.(*domain.Message), r1
}
