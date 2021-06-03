package source

import (
	"log"

	"github.com/kabi175/chat-app-go/domain"
)

type source struct{}

func NewSource() source {
	return source{}
}

var users map[domain.UserId]*domain.User

func (source) CreateUser(user *domain.User) {
	users[user.Id()] = user
}

func (source) Store(message domain.Message) {
	log.Printf(message.Text)
}

func (source) GetUser(id domain.UserId) *domain.User {
	return users[id]
}
