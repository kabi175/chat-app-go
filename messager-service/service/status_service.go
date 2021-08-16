package service

import (
	"github.com/kabi175/chat-app-go/messager/domain"
)

type UserStatusService struct {
	sr domain.UserStatusRepository
}

type UserStatusServiceConfig struct {
	UserStatusRepository domain.UserStatusRepository
}

func NewUserStatusRepository(c *UserStatusServiceConfig) domain.UserStatusService {
	return &UserStatusService{
		sr: c.UserStatusRepository,
	}
}

func (ms *UserStatusService) Publish(status *domain.UserStatus) error {
	return ms.sr.Publish(status)
}

func (ms *UserStatusService) Listern(user *domain.User, statusChan chan domain.UserStatus) {
	go ms.sr.Listern(user, statusChan)
}
