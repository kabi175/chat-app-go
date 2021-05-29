// Package chatroom provides room to communicate
package chatroom

type Member struct {
	Name string
	Id   uint8
}

type ChatRoom struct {
	Name    string
	Id      uint8
	Members map[uint8]Member
}

func (r *ChatRoom) Create() {}

func (r *ChatRoom) Join() {}

func (r *ChatRoom) Leave() {}

func (r *ChatRoom) Echo() {}
