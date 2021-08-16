package domain

type MessageRepository interface {
	Publish(*UserMessage) error
	Listern(*User, chan UserMessage)
}

type UserStatusRepository interface {
	Publish(*UserStatus) error
	Listern(*User, chan UserStatus)
}
