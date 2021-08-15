package mocks

import (
	"github.com/kabi175/chat-app-go/messager/model"
	"github.com/stretchr/testify/mock"
)

type MessageRepository struct {
	mock.Mock
}

func (mr *MessageRepository) Publish(message *model.Message) error {
	args := mr.Called(message)
	return args.Error(0)
}

func (mr *MessageRepository) Listern(user *model.User) (*model.Message, error) {
	args := mr.Called(user)
	r0 := args.Get(0)
	r1 := args.Error(1)
	if r0 == nil {
		return nil, r1
	}
	return r0.(*model.Message), r1
}
