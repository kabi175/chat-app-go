package domain

type Message struct {
	Text       string `json:"Text"`
	SenderId   UserId `json:"SenderId"`
	ReceiverId RoomId `json:"ReceiverId"`
}
