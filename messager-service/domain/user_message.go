package domain

import "log"

type UserMessage struct {
	Type    int8       `json:"type"`
	Message Message    `json:"message"`
	Status  UserStatus `json:"status"`
	Ack     AckMessage `json:"ackMessage"`
}

func NewUserMessage(obj interface{}) *UserMessage {
	switch o := obj.(type) {
	case *Message:
		return &UserMessage{
			Type:    TypeMessage,
			Message: *o,
		}
	case *UserStatus:
		return &UserMessage{
			Type:   TypeStatus,
			Status: *o,
		}
	case *AckMessage:
		return &UserMessage{
			Type: TypeAck,
			Ack:  *o,
		}
	default:
		log.Fatalf("unknown type %T in NewUserMessage\n", obj)
	}
	return nil
}
