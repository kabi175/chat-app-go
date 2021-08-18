package domain

import (
	"encoding/json"
	"fmt"
)

type AckMessage struct {
	To        string `json:"to"`
	From      string `json:"from"`
	Type      int8   `json:"type"`
	MessageID string `json:"messageID"`
	Status    string `json:"status"`
	Error     string `json:"error"`
}

func (AckMessage) NewAckMessage(obj interface{}) (*AckMessage, error) {
	var message AckMessage
	switch o := obj.(type) {
	case string:
		err := json.Unmarshal([]byte(o), &message)
		return &message, err
	case []byte:
		err := json.Unmarshal(o, &message)
		return &message, err
	case *AckMessage:
		return o, nil
	case AckMessage:
		return &o, nil
	default:
		panic(fmt.Errorf("Unknown obj of type %T in NewAckMessage", obj))
	}
}

func NewError(message Message, err error) *AckMessage {
	return &AckMessage{
		To:        message.To,
		From:      message.From,
		Type:      TypeAck,
		MessageID: message.MessageID,
		Status:    "failed",
		Error:     err.Error(),
	}
}

func NewAckRead(message Message) *AckMessage {
	return &AckMessage{
		To:        message.To,
		From:      message.From,
		Type:      TypeAck,
		MessageID: message.MessageID,
		Status:    "read",
	}
}

func NewAckDelivered(message Message) *AckMessage {
	return &AckMessage{
		To:        message.To,
		From:      message.From,
		Type:      TypeAck,
		MessageID: message.MessageID,
		Status:    "delivered",
	}
}

func NewAckReceived(message Message) *AckMessage {
	return &AckMessage{
		To:        message.To,
		From:      message.From,
		Type:      TypeAck,
		MessageID: message.MessageID,
		Status:    "received",
	}
}
