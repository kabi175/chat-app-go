package service

import (
	"context"

	"github.com/kabi175/chat-app-go/messager/domain"
)

type DefaultMessageService struct {
	messageRepo domain.MessageRepo
}

func NewDefaultMessageService(messageRepo domain.MessageRepo) domain.MessageService {
	return &DefaultMessageService{messageRepo: messageRepo}
}

func (s *DefaultMessageService) PublishMessage() func(*domain.Message) error {
	return func(message *domain.Message) error {
		err := s.messageRepo.Producer(message, context.Background())
		return err
	}
}

func (s *DefaultMessageService) ConsumeMessage(user *domain.User, ctx context.Context) (func() <-chan domain.MessageChan, error) {
	msgChan, err := s.messageRepo.Consumer(user, ctx)
	return func() <-chan domain.MessageChan { return msgChan }, err
}
