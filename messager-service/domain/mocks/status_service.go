package mocks

import (
	"github.com/kabi175/chat-app-go/messager/model"
	"github.com/stretchr/testify/mock"
)

type UserStatusService struct {
	mock.Mock
}

func (us *UserStatusService) Publish(status *model.UserStatus) error {
	args := us.Called(status)
	return args.Error(0)
}
func (us *UserStatusService) Listern(user *model.User) (*model.UserStatus, error) {
	args := us.Called(user)
	r0 := args.Get(0)
	r1 := args.Error(1)
	if r0 == nil {
		return nil, r1
	}
	return r0.(*model.UserStatus), r1
}
