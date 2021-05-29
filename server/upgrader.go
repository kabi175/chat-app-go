package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/kabi175/chat-app-go/chatroom"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var room = chatroom.New("hello room")
var roomCh = make(chan string, 1)

func init() {
	go room.Echo(roomCh)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	defer conn.Close()

	if err != nil {
		log.Println(err)
	}
	room.Add(conn)
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		roomCh <- string(msg)
	}
}
