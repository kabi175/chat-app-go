// Package user provides methods to handle
// sending messages and closing connection
package user

import "github.com/gorilla/websocket"

func NewUser(id UserId, conn *websocket.Conn) *User {
	return &User{
		id:   id,
		conn: conn,
	}
}

type UserId string

type User struct {
	id   UserId
	conn *websocket.Conn
}

func (u *User) Close() {
	u.conn.Close()
	u.conn = nil
}

func (u *User) Id() UserId {
	return u.id
}

func (u *User) Send(message interface{}) bool {
	if ok := u.IsOnline(); ok {
		err := u.conn.WriteJSON(message)

		if err == nil {
			return true
		}
	}
	return false
}

func (u *User) IsOnline() bool {
	if u.conn != nil {
		return true
	}
	return false
}
