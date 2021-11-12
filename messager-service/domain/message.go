package domain

import (
	"time"
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
