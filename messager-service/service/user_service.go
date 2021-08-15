package service

import (
	"context"
	"encoding/json"
	"log"

	"github.com/kabi175/chat-app-go/messager/model"
)

type UserService struct {
	m model.MessageService
	s model.UserStatusService
}

type UserServiceConfig struct {
	MessageService    model.MessageService
	UserStatusService model.UserStatusService
}

func NewUserService(c *UserServiceConfig) model.UserService {
	return &UserService{
		m: c.MessageService,
		s: c.UserStatusService,
	}
}

func (us *UserService) Listner(user *model.User) {
	for {
		var message model.UserMessage
		user.Conn.ReadJSON(&message)
		switch message.Type {
		case model.TypeMessage:
			if message.Message.From == user.UserName {
				err := us.m.Publish(&message.Message)
				if err != nil {
					log.Println(err)
				}
			}
		case model.TypeStatus:
			if message.Status.UserName == user.UserName {
				err := us.s.Publish(&message.Status)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}

func (us *UserService) Writer(user *model.User) {
	userMessageBufferSize := 10
	userMessageChan := make(chan model.UserMessage, userMessageBufferSize)
	go func() {
		for {
			message, err := us.m.Listern(user)
			if err != nil {
				log.Println(err)
				continue
			}
			userMessageChan <- model.UserMessage{Message: *message}
		}
	}()
	go func() {
		pubsub := us.s.Listern(user)
		for {
			msg, err := pubsub.ReceiveMessage(context.TODO())
			if err != nil {
				log.Println(err)
				continue
			}
			var status model.UserStatus
			log.Println(msg.Payload)
			err = json.Unmarshal([]byte(msg.Payload), &status)
			if err != nil {
				log.Println(err)
				continue
			}
			userMessageChan <- model.UserMessage{Status: status}
		}
	}()

	for {
		userMessage := <-userMessageChan
		err := user.Conn.WriteJSON(userMessage)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func (us *UserService) Validate(token string) (*model.User, error) {
	return &model.User{UserName: token}, nil
}
