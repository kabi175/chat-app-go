package domain

import (
	"encoding/json"
	"errors"
)

type AckMessage struct {
	To        string `json:"to"`
	From      string `json:"from"`
	Type      int8   `json:"type"`
	MessageID string `json:"messageID"`
	Status    string `json:"status"`
	Error     string `json:"error"`
}

func (message *AckMessage) Bytes() ([]byte, error) {
	return json.Marshal(message)
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
		return nil, errors.New("Unknown type obj")
	}
}

func NewError(messageID string, err error) *AckMessage {
	return &AckMessage{
		Type:      TypeAck,
		MessageID: messageID,
		Status:    "failed",
		Error:     err.Error(),
	}
}

func NewAckRead(messageid string) *AckMessage {
	return &AckMessage{
		Type:      TypeAck,
		MessageID: messageid,
		Status:    "read",
	}
}

func NewAckDelivered(messageid string) *AckMessage {
	return &AckMessage{
		Type:      TypeAck,
		MessageID: messageid,
		Status:    "delivered",
	}
}

func NewAckReceived(messageid string) *AckMessage {
	return &AckMessage{
		Type:      TypeAck,
		MessageID: messageid,
		Status:    "received",
	}
}
