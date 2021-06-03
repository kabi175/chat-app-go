package source

import (
	"github.com/kabi175/chat-app-go/domain"
)

var rooms map[domain.RoomId]*domain.Room

func (source) CreateRoom(room *domain.Room) {
	rooms[room.Id()] = room
}

func (s *source) GetRoom(roomId domain.RoomId) *domain.Room {
	return rooms[roomId]
}
