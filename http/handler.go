package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/kabi175/chat-app-go/domain"
)

type handler struct {
	service Service
}

func NewHandler(service Service) handler {
	return handler{
		service: service,
	}
}

func (h handler) Login(w http.ResponseWriter, r *http.Request) {
	var user struct {
		id       string
		password string
	}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return
	}
	h.service.LogIn(user.id, user.password)
}

func (h handler) Upgrader(w http.ResponseWriter, r *http.Request) {

	var (
		userId domain.UserId
	)

	upgrader := websocket.Upgrader{

		CheckOrigin: func(r *http.Request) bool {
			auth, err := r.Cookie("auth")
			userId, err = h.service.Validate(auth.String())
			if err != nil {
				return false
			}
			return true
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	h.service.SpinUser(userId, conn)

	var msg domain.Message

	for {

		err := conn.ReadJSON(msg)

		if err != nil {
			log.Println(err)
			return
		}
		h.service.ConsumeMessage(msg)

	}
}
