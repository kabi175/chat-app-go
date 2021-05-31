package chatroom

import "github.com/kabi175/chat-app-go/user"

type Message struct {
	Text     string
	SenderId user.UserId
}
