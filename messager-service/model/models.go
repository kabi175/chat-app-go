package model

import (
	"time"

	"github.com/gorilla/websocket"
)

type User struct {
	UserName string
	Conn     *websocket.Conn
}

type Message struct {
	From string        `json:"from"`
	To   string        `json:"to"`
	Data string        `json:"data"`
	Time time.Duration `json:"time"`
}

type UserStatus struct {
	UserName string `json:"username"`
	Status   string `json:"status"`
}

type AckMessage struct {
	Type      int8   `json:"type"`
	MessageID string `json:"messageID"`
	Status    string `json:"status"`
}

type UserMessage struct {
	Type    int8       `json:"type"`
	Message Message    `json:"message"`
	Status  UserStatus `json:"status"`
	Ack     AckMessage `json:"ackMessage"`
}

const (
	TypeMessage = iota
	TypeStatus
)
