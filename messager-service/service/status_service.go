package service

import (
	"github.com/go-redis/redis/v8"
	"github.com/kabi175/chat-app-go/messager/model"
)

type UserStatusService struct {
	sr model.UserStatusRepository
}

type UserStatusServiceConfig struct {
	UserStatusRepository model.UserStatusRepository
}

func NewUserStatusRepository(c *UserStatusServiceConfig) model.UserStatusService {
	return &UserStatusService{
		sr: c.UserStatusRepository,
	}
}

func (ms *UserStatusService) Publish(status *model.UserStatus) error {
	return ms.sr.Publish(status)
}

func (ms *UserStatusService) Listern(user *model.User) *redis.PubSub {
	return ms.sr.Listern(user)
}
