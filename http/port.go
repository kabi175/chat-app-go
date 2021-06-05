package http

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/kabi175/chat-app-go/domain"
)

type Service interface {
	SpinUser(domain.UserId, *websocket.Conn) error
	SpinRoom(domain.UserId, domain.RoomId) error
	ConsumeMessage(domain.Message) error
	SubscribeRoom(domain.UserId, domain.RoomId) error
	SignUp(string, string) (string, error)
	LogIn(string, string) (string, error)
	Validate(string) (string, error)
}

type Handler interface {
	Login(http.ResponseWriter, *http.Request)
	SignUp(http.ResponseWriter, *http.Request)
	CreateRoom(http.ResponseWriter, *http.Request)
	JoinRoom(http.ResponseWriter, *http.Request)
	Upgrader(http.ResponseWriter, *http.Request)
}
