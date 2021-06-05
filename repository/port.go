package repository

import "github.com/kabi175/chat-app-go/domain"

type DataSource interface {
	GetUser(domain.UserId) (*domain.User, error)
	GetRoom(domain.RoomId) (*domain.Room, error)
	Store(domain.Message)

	CreateUser(*domain.User) error
	CreateRoom(*domain.Room) error

	AddUser(string, string) error
	RemoveUser(string) error
	UpdateUser(string, string) error
	GetPass(string) (string, error)
}
