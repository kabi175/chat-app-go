package http

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/kabi175/chat-app-go/domain"
)

type Service interface {
	SpinUser(domain.UserId, *websocket.Conn)
	SpinRoom(domain.UserId, domain.RoomId)
	ConsumeMessage(domain.Message)
	SignUp()
	LogIn(string, string)
	Validate(string) (domain.UserId, error)
}

type Handler interface {
	Login(http.ResponseWriter, *http.Request)
	SignUp(http.ResponseWriter, *http.Request)
	CreateUser(http.ResponseWriter, *http.Request)
	CreateRoom(http.ResponseWriter, *http.Request)
	JoinRoom(http.ResponseWriter, *http.Request)
	Upgrader(http.ResponseWriter, *http.Request)
}
