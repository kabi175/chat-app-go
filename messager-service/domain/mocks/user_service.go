package mocks

import (
	"log"

	"github.com/kabi175/chat-app-go/messager/domain"
	"github.com/stretchr/testify/mock"
)

type UserService struct {
	mock.Mock
}

func (us *UserService) Listner(user *domain.User) {
	_, message, err := user.Conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(message))
}

func (us *UserService) Writer(user *domain.User) {}

func (us *UserService) Validate(token string) (*domain.User, error) {
	return &domain.User{UserName: "mock"}, nil
}
