package domain

import "context"

type MessageRepo interface {
	Producer(*Message, context.Context) error
	Consumer(*User, context.Context) (<-chan MessageChan, error)
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
