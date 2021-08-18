package mocks

import (
	"github.com/kabi175/chat-app-go/messager/domain"
	"github.com/stretchr/testify/mock"
)

type UserStatusRepository struct {
	mock.Mock
}

func (ur *UserStatusRepository) Publish(status *domain.UserStatus) error {
	args := ur.Called(status)
	return args.Error(0)
}
func (ur *UserStatusRepository) Listern(user *domain.User) (*domain.UserStatus, error) {
	args := ur.Called(user)
	r0 := args.Get(0)
	r1 := args.Error(1)
	if r0 == nil {
		return nil, r1
	}
	return r0.(*domain.UserStatus), r1
}
