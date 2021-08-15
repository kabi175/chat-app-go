package repository

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/kabi175/chat-app-go/messager/model"
)

type UserStatusRepository struct {
	db *redis.Client
}

type UserStatusRepositoryConfig struct {
	Redis *redis.Client
}

func NewUserStatusRepository(c *UserStatusRepositoryConfig) model.UserStatusRepository {
	return &UserStatusRepository{
		db: c.Redis,
	}
}

func (s *UserStatusRepository) Publish(status *model.UserStatus) error {
	statusByte, err := json.Marshal(status)
	if err != nil {
		return err
	}
	s.db.Publish(context.TODO(), status.UserName, statusByte)
	return nil
}

func (s *UserStatusRepository) Listern(user *model.User) *redis.PubSub {
	sub := s.db.Subscribe(context.TODO(), user.UserName)
	return sub
}
