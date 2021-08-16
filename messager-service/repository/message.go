package repository

import (
	"context"
	"encoding/json"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/kabi175/chat-app-go/messager/domain"
)

type MessageRepository struct {
	db *redis.Client
}

type MessageRepositoryConfig struct {
	Redis *redis.Client
}

func NewMessageRepository(c *MessageRepositoryConfig) domain.MessageRepository {
	return &MessageRepository{
		db: c.Redis,
	}
}

func (m *MessageRepository) Publish(message *domain.UserMessage) error {
	var userQueue string

	switch message.Type {
	case domain.TypeAck:
		userQueue = message.Ack.To
	case domain.TypeMessage:
		userQueue = message.Message.To
	}

	messageByte, err := json.Marshal(message)
	if err != nil {
		return err
	}
	result := m.db.LPush(context.TODO(), userQueue, string(messageByte))
	return result.Err()
}

func (m *MessageRepository) Listern(user *domain.User, messageChan chan domain.UserMessage) {
	for {
		resultString, err := m.db.BRPop(context.TODO(), 0, user.UserName).Result()
		if err != nil {
			log.Println(err)
			continue
		}

		byteMessage := []byte(resultString[1])

		var message domain.UserMessage
		err = json.Unmarshal(byteMessage, &message)
		if err != nil {
			log.Println(err)
			continue
		}
		messageChan <- *domain.NewUserMessage(message)
	}
}
