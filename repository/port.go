package repository

import "github.com/kabi175/chat-app-go/domain"

type Datasource interface {
	GetUser(domain.UserId) *domain.User
	GetRoom(domain.RoomId) *domain.Room
	Store(domain.Message)
	CreateUser(*domain.User)
	CreateRoom(*domain.Room)
}
