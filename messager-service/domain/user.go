package domain

import (
	"github.com/gorilla/websocket"
)

type User struct {
	UserName string
	Conn     *websocket.Conn
	Wait     chan struct{}
}

func NewUser(userName string) *User {
	return &User{
		UserName: userName,
		Conn:     nil,
		Wait:     make(chan struct{}),
	}
}
