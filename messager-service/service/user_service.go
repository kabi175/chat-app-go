package service

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/kabi175/chat-app-go/messager/domain"
)

type UserService struct {
	m domain.MessageService
	s domain.UserStatusService
}

type UserServiceConfig struct {
	MessageService    domain.MessageService
	UserStatusService domain.UserStatusService
}

func NewUserService(c *UserServiceConfig) domain.UserService {
	return &UserService{
		m: c.MessageService,
		s: c.UserStatusService,
	}
}

func (us *UserService) Listner(user *domain.User) {
	var userMessage domain.UserMessage
	for {
		select {
		case <-user.Wait:
			return
		default:
			err := user.Conn.ReadJSON(&userMessage)
			if err != nil {
				if err == websocket.ErrCloseSent {
					close(user.Wait)
					return
				}
			}
			switch userMessage.Type {
			case domain.TypeMessage:
				userMessage.Message.RecordTime()
				userMessage = *domain.NewUserMessage(&userMessage.Message)
				if err = us.m.Publish(&userMessage); err != nil {
					log.Println(err)
					continue
				}
				ackMessage := domain.NewAckReceived(userMessage.Message.MessageID)
				userMessage = *domain.NewUserMessage(&ackMessage)
				if err = us.m.Publish(&userMessage); err != nil {
					log.Println(err)
				}
			case domain.TypeStatus:
				if err = us.s.Publish(&userMessage.Status); err != nil {
					log.Println(err)
				}
			case domain.TypeAck:
				userMessage = *domain.NewUserMessage(&userMessage.Ack)
				if err = us.m.Publish(&userMessage); err != nil {
					log.Println(err)
				}
			}
		}
	}
}

func (us *UserService) Writer(user *domain.User) {
	messageChan := make(chan domain.UserMessage, 5)
	statusChan := make(chan domain.UserStatus, 1)
	go us.m.Listern(user, messageChan)
	go us.s.Listern(user, statusChan)
	for {
		select {
		case <-user.Wait:
			return
		case message := <-messageChan:
			err := user.Conn.WriteJSON(message)
			if err != nil {
				if err == websocket.ErrCloseSent {
					close(user.Wait)
					return
				}
			}
		case status := <-statusChan:
			err := user.Conn.WriteJSON(domain.NewUserMessage(&status))
			if err != nil {
				if err == websocket.ErrCloseSent {
					close(user.Wait)
					return
				}
			}
		}
	}
}

func (us *UserService) Validate(token string) (*domain.User, error) {
	return domain.NewUser(token), nil
}
