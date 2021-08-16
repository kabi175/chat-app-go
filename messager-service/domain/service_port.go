package domain

type UserService interface {
	Listner(*User)
	Writer(*User)
	Validate(string) (*User, error)
}

type MessageService interface {
	Publish(*UserMessage) error
	Listern(*User, chan UserMessage)
}

type UserStatusService interface {
	Publish(*UserStatus) error
	Listern(*User, chan UserStatus)
}
