package service

import (
	"log"

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

// Read message from socket connection
func (us *UserService) Listner(user *domain.User) {
	var userMessage domain.UserMessage
	for {
		// Check for Listner Shutdown
		select {
		case <-user.Wait:
			return
		default:
		}

		err := user.Conn.ReadJSON(&userMessage)
		if err != nil {
			log.Println(err)
			close(user.Wait)
			return
		}

		switch userMessage.Type {

		case domain.TypeMessage:
			if userMessage.Message.From != user.UserName {
				continue
			}

			userMessage.Message.RecordTime()

			userMessage = *domain.NewUserMessage(&userMessage.Message)
			if err = us.m.Publish(&userMessage); err != nil {
				log.Println(err)
				continue
			}
			ackMessage := domain.NewAckReceived(userMessage.Message)
			userMessage = *domain.NewUserMessage(ackMessage)
			if err = us.m.Publish(&userMessage); err != nil {
				log.Println(err)
			}

		case domain.TypeStatus:
			if userMessage.Status.UserName != user.UserName {
				continue
			}
			// Publish status
			if err = us.s.Publish(&userMessage.Status); err != nil {
				log.Println(err)
			}

		case domain.TypeAck:
			// Block Unauthorized message
			if userMessage.Ack.From != user.UserName {
				continue
			}
			// format Ack message
			userMessage = *domain.NewUserMessage(&userMessage.Ack)
			// Publish Ack message
			if err = us.m.Publish(&userMessage); err != nil {
				log.Println(err)
			}
		}
	}
}

// Writer method writes message incoming message to socket
func (us *UserService) Writer(user *domain.User) {
	// communication chan for message & status  services
	messageChan := make(chan domain.UserMessage, 5)
	statusChan := make(chan domain.UserStatus, 5)

	go us.m.Listern(user, messageChan)
	go us.s.Listern(user, statusChan)

	for {
		select {
		case <-user.Wait:
			return
		case message := <-messageChan:
			err := user.Conn.WriteJSON(message)
			if err != nil {
				log.Println(err)
				close(user.Wait)
				return
			}
		case status := <-statusChan:
			err := user.Conn.WriteJSON(domain.NewUserMessage(&status))
			if err != nil {
				log.Println(err)
				close(user.Wait)
				return
			}
		}
	}
}

// Decode's  JWT token
func (us *UserService) Validate(token string) (*domain.User, error) {
	return domain.NewUser(token), nil
}
