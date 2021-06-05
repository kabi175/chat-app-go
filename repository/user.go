package repository

import (
	"github.com/gorilla/websocket"
	"github.com/kabi175/chat-app-go/domain"
)

func (s *service) SpinUser(userId domain.UserId, conn *websocket.Conn) error {
	user := domain.NewUser(userId, conn)
	err := s.dataSource.CreateUser(user)
	return err
}
