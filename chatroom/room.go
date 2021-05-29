// Package chatroom provides room to communicate
package chatroom

import "math/rand"

const buff = 4

type Member struct {
	Name string
	Id   uint32
}

type Message struct {
	text     string
	senderId uint32
}

type ChatRoom struct {
	Name     string
	Id       uint32
	Members  map[uint32]Member
	echoChnl chan Message
}

func Create(name string, member Member) *ChatRoom {
	var members map[uint32]Member
	members[member.Id] = member
	echoChnl := make(chan Message, buff)
	return &ChatRoom{
		Name:     name,
		Id:       rand.Uint32(),
		Members:  members,
		echoChnl: echoChnl,
	}
}

func Delete() error {
	return nil
}

func (r *ChatRoom) Join() {}

func (r *ChatRoom) Leave() {}

func (r *ChatRoom) Echo() {}
