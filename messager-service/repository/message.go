package repository

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/kabi175/chat-app-go/messager/model"
)

type MessageRepository struct {
	db *redis.Client
}

type MessageRepositoryConfig struct {
	Redis *redis.Client
}

func NewMessageRepository(c *MessageRepositoryConfig) model.MessageRepository {
	return &MessageRepository{
		db: c.Redis,
	}
}

func (m *MessageRepository) Publish(message *model.Message) error {
	userQueue := message.To
	messageByte, err := json.Marshal(message)
	if err != nil {
		return err
	}
	result := m.db.LPush(context.TODO(), userQueue, string(messageByte))
	return result.Err()
}

func (m *MessageRepository) Listern(user *model.User) (*model.Message, error) {
	resultString, err := m.db.BRPop(context.TODO(), 0, user.UserName).Result()
	if err != nil {
		return nil, err
	}
	var message model.Message
	byteMessage := []byte(resultString[1])
	err = json.Unmarshal(byteMessage, &message)
	if err != nil {
		return nil, err
	}
	return &message, nil
}
