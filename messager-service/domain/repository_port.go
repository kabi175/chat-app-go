package domain

type MessageRepository interface {
	Publish(*UserMessage) error
	Listern(*User, chan UserMessage)
}

type UserStatusRepository interface {
	Publish(*UserStatus) error
	Listern(*User, chan UserStatus)
}

type UserRepo interface {
	Create(*User) error
	GetByID(uint) (*User, error)
	GetByEmail(string) (*User, error)
	DeleteByID(uint) error
	DeleteByEmail(string) error
}
