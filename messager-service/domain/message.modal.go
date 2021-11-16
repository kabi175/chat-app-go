package domain

import (
	"encoding/json"
	"time"

	"github.com/kabi175/chat-app-go/messager/domain/apperrors"
)

type Message struct {
	MessageID string `json:"messageID"`
	From      uint   `json:"from"`
	To        uint   `json:"to"`
	Text      string `json:"text"`
	Seen      bool   `json:"seen"`
	Received  bool   `json:"received"`
	CreatedAt int64  `json:"createdAt"`
}

func (m *Message) RecordTime() *Message {
	m.CreatedAt = time.Now().Unix()
	return m
}
func (u *Message) String() (string, error) {
	str, err := json.Marshal(u)
	if err != nil {
		return "", apperrors.NewInternalServerError(err.Error())
	}
	return string(str), nil
}

type MessageChan struct {
	Message Message
	Error   error
}
