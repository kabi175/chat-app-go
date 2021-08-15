package service

import "github.com/kabi175/chat-app-go/messager/model"

type MessageService struct {
	mr model.MessageRepository
}

type MessageServiceConfig struct {
	MessageRepository model.MessageRepository
}

func NewMessageRepository(c *MessageServiceConfig) model.MessageService {
	return &MessageService{
		mr: c.MessageRepository,
	}
}

func (ms *MessageService) Publish(message *model.Message) error {
	return ms.mr.Publish(message)
}
func (ms *MessageService) Listern(user *model.User) (*model.Message, error) {
	return ms.mr.Listern(user)
}
