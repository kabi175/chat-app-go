package user

var users map[UserId]*User

func Ref(id UserId) *User {
	return users[id]
}
