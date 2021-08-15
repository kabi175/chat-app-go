package model

import "github.com/go-redis/redis/v8"

type UserService interface {
	Listner(*User)
	Writer(*User)
	Validate(string) (*User, error)
}

type MessageService interface {
	Publish(*Message) error
	Listern(*User) (*Message, error)
}

type UserStatusService interface {
	Publish(*UserStatus) error
	Listern(*User) *redis.PubSub
}

type MessageRepository interface {
	Publish(*Message) error
	Listern(*User) (*Message, error)
}

type UserStatusRepository interface {
	Publish(*UserStatus) error
	Listern(*User) *redis.PubSub
}
