package repository

import (
	"github.com/kabi175/chat-app-go/domain"
)

func (s *service) SendMessage(id domain.UserId, msg domain.Message) error {
	user, err := s.dataSource.GetUser(id)
	if err != nil {
		return err
	}
	if ok := user.Send(msg); !ok {
		s.dataSource.Store(msg)
	}
	return nil
}

func (s *service) SpinRoom(admin domain.UserId, roomId domain.RoomId) error {
	service := NewService(s.dataSource)
	room := domain.CreateRoom(roomId, admin, &service)
	err := s.dataSource.CreateRoom(room)
	if err != nil {
		return err
	}
	go room.Echo()
	return nil
}

func (s *service) ConsumeMessage(message domain.Message) error {
	roomId := message.ReceiverId
	room, err := s.dataSource.GetRoom(roomId)
	if err != nil {
		return err
	}
	room.Consume(message)
	return nil
}

func (s *service) SubscribeRoom(userId domain.UserId, roomId domain.RoomId) error {
	room, err := s.dataSource.GetRoom(roomId)
	if err != nil {
		return err
	}
	room.Join(userId)
	return nil
}
