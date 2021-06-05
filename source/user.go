package source

import (
	"errors"
	"log"

	"github.com/kabi175/chat-app-go/domain"
)

type source struct{}

func NewSource() source {
	return source{}
}

var users map[domain.UserId]*domain.User

func init() {
	users = make(map[domain.UserId]*domain.User)
}

func (source) CreateUser(user *domain.User) error {
	if _, ok := users[user.Id()]; ok {
		return errors.New("User slready exist")
	}
	users[user.Id()] = user
	return nil
}

func (source) Store(message domain.Message) {
	log.Printf(message.Text)
}

func (source) GetUser(id domain.UserId) (*domain.User, error) {
	if _, ok := users[id]; !ok {
		return nil, errors.New("User not found")
	}
	return users[id], nil
}
