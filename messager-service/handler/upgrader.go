package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h *Handler) Upgrader(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("AUTH")

	user, err := h.us.Validate(token)
	if err != nil {
		log.Println(err)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(user.UserName, " connected...")
	defer conn.Close()
	user.Conn = conn
	go h.us.Writer(user)
	h.us.Listner(user)
}
