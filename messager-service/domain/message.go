package domain

import (
	"time"
)

type Message struct {
	MessageID string `json:"messageID"`
	From      string `json:"from"`
	To        string `json:"to"`
	Data      string `json:"data"`
	Time      int64  `json:"time"`
}

func (m *Message) RecordTime() {
	m.Time = time.Now().Unix()
}
