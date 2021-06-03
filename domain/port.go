package domain

type Service interface {
	SendMessage(UserId, Message)
}
