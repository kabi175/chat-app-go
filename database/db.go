// Package databse provides virtual databse for ChatApp
package databse

import (
	"errors"

	"github.com/kabi175/chat-app-go/chatroom"
)

type mango struct{}

func NewDb() *mango {
	return &mango{}
}

// Refercing types
type ChatRoom = chatroom.ChatRoom

type Member = chatroom.Member

type ChatRooms map[uint32]ChatRoom

var chatRooms ChatRooms

func (mango) Rooms() (*ChatRooms, error) {
	return &chatRooms, nil
}

func (mango) AddRoom(room ChatRoom) error {
	if _, ok := chatRooms[room.Id]; ok {
		return errors.New("ChatRoom alredy exist")
	}
	chatRooms[room.Id] = room
	return nil
}

func (mango) DeleteRoom(roomId uint32) error {
	if _, ok := chatRooms[roomId]; ok == false {
		return errors.New("ChatRoom not found")
	}
	delete(chatRooms, roomId)
	return errors.New("ChatRoom not found")
}

func (mango) AddUser(roomId uint32, member Member) error {

	if _, ok := chatRooms[roomId]; ok == false {
		return errors.New("ChatRoom not found")
	}
	room := chatRooms[roomId]
	if _, ok := room.Members[member.Id]; ok {
		return errors.New("Member alreday exist")
	}
	room.Members[member.Id] = member
	chatRooms[roomId] = room
	return nil
}

func (mango) RemoveUser(roomId uint32, member Member) error {

	if _, ok := chatRooms[roomId]; ok == false {
		return errors.New("ChatRoom not found")
	}
	room := chatRooms[roomId]
	if _, ok := room.Members[member.Id]; ok == false {
		return errors.New("Member not exist")
	}
	delete(room.Members, member.Id)
	chatRooms[roomId] = room
	return nil
}
