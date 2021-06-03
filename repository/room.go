package repository

import (
	"github.com/kabi175/chat-app-go/domain"
)

func (s *service) SendMessage(id domain.UserId, msg domain.Message) {
	user := s.datasource.GetUser(id)
	if ok := user.Send(msg); ok != true {
		s.datasource.Store(msg)
	}
	return
}

func (s *service) SpinRoom(roomId domain.RoomId, admin domain.UserId) {
	service := NewService(s.datasource)
	room := domain.CreateRoom(roomId, admin, &service)
	s.datasource.CreateRoom(room)
	go room.Echo()
}

func (s *service) ConsumeMessage(message domain.Message) {
	roomId := message.ReceiverId
	room := s.datasource.GetRoom(roomId)
	room.Consume(message)
}
