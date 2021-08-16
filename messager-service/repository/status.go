package repository

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/kabi175/chat-app-go/messager/domain"
)

type UserStatusRepository struct {
	db *redis.Client
}

type UserStatusRepositoryConfig struct {
	Redis *redis.Client
}

func NewUserStatusRepository(c *UserStatusRepositoryConfig) domain.UserStatusRepository {
	return &UserStatusRepository{
		db: c.Redis,
	}
}

func (s *UserStatusRepository) Publish(status *domain.UserStatus) error {
	statusByte, err := json.Marshal(status)
	if err != nil {
		return err
	}
	s.db.Publish(context.TODO(), status.UserName, statusByte)
	return nil
}

func (s *UserStatusRepository) Listern(user *domain.User, statusChan chan domain.UserStatus) {
	sub := s.db.Subscribe(context.TODO(), user.UserName)
	redisChan := sub.Channel()
	for msg := range redisChan {
		statusChan <- *domain.NewUserStatus(msg.Payload)
	}
}
