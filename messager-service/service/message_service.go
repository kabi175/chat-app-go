package service

import "github.com/kabi175/chat-app-go/messager/domain"

type MessageService struct {
	mr domain.MessageRepository
}

type MessageServiceConfig struct {
	MessageRepository domain.MessageRepository
}

func NewMessageRepository(c *MessageServiceConfig) domain.MessageService {
	return &MessageService{
		mr: c.MessageRepository,
	}
}

func (ms *MessageService) Publish(message *domain.UserMessage) error {
	return ms.mr.Publish(message)
}
func (ms *MessageService) Listern(user *domain.User, messageChan chan domain.UserMessage) {
	ms.mr.Listern(user, messageChan)
}
