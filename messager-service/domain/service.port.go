package domain

type UserService interface {
	LogIn(*User) (string, error)
	SignUp(*User) (*User, error)
	GetByID(uint) (*User, error)
	GetByEmail(string) (*User, error)
	Delete(*User) error
}

type TokenService interface {
	GenerateToken(*User) (string, error)
	DecodeToken(string) (*User, error)
}

type MessageService interface {
	PostMessage(*Message) error
	GetMessage() error
}

type UserStatusService interface {
	Publish(*UserStatus) error
	Listern(*User, chan UserStatus)
}
