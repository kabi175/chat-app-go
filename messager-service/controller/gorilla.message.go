package controller

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/kabi175/chat-app-go/messager/domain"
)

type GorillaMessageController struct {
	messageService domain.MessageService
}

func NewGorillaMessageController(messageService domain.MessageService) domain.MessageController {
	return &GorillaMessageController{messageService: messageService}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (c *GorillaMessageController) Upgrade(w http.ResponseWriter, r *http.Request) {
	userIDstr := r.Header.Get("userID")
	userID, err := strconv.ParseUint(userIDstr, 10, 64)
	if err != nil {
		http.Error(w, "invalid userID", 500)
		return
	}
	user := domain.User{
		ID: uint(userID),
	}

	// Upgrade to socket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	handleError := func(err error) {
		log.Println(err)
		cancel()
	}

	// Publish message new messages
	publisher := c.messageService.PublishMessage()
	go func() {
		var message domain.Message
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}
			err := conn.ReadJSON(&message)
			if err != nil {
				handleError(err)
				return
			}
			err = publisher(&message)
			if err != nil {
				handleError(err)
				return
			}
		}
	}()

	// Receive Incomming Message
	consumer, err := c.messageService.ConsumeMessage(&user, ctx)
	handleError(err)
	for {
		select {
		case <-ctx.Done():
			return

		case msg := <-consumer():
			if msg.Error != nil {
				handleError(err)
				return
			}
			err := conn.WriteJSON(msg.Message)
			if err != nil {
				handleError(err)
				return
			}
		}
	}
}
