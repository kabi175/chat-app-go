package source

import (
	"errors"

	"github.com/kabi175/chat-app-go/domain"
)

var rooms map[domain.RoomId]*domain.Room

func init() {
	rooms = make(map[domain.RoomId]*domain.Room)
}
func (source) CreateRoom(room *domain.Room) error {
	if _, ok := rooms[room.Id()]; ok {
		return errors.New("Room already exist")
	}
	rooms[room.Id()] = room
	return nil
}

func (s *source) GetRoom(roomId domain.RoomId) (*domain.Room, error) {
	if _, ok := rooms[roomId]; !ok {
		return nil, errors.New("Room not found")
	}
	return rooms[roomId], nil
}
