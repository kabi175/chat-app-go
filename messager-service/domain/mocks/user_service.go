package mocks

import (
	"log"

	"github.com/kabi175/chat-app-go/messager/model"
	"github.com/stretchr/testify/mock"
)

type UserService struct {
	mock.Mock
}

func (us *UserService) Listner(user *model.User) {
	_, message, err := user.Conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(message))
}

func (us *UserService) Writer(user *model.User) {}

func (us *UserService) Validate(token string) (*model.User, error) {
	return &model.User{UserName: "mock"}, nil
}
