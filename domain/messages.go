package domain

type Message struct {
	Text       string
	SenderId   UserId
	ReceiverId UserId
}
