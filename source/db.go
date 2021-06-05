package source

import "errors"

var userDB map[string]string

func init() {
	userDB = make(map[string]string)
	userDB["default"] = "secret"
}

func (source) AddUser(userId string, password string) error {
	if _, ok := userDB[userId]; ok {
		return errors.New("User already exist")
	}
	userDB[userId] = password
	return nil
}

func (source) RemoveUser(userId string) error {
	if _, ok := userDB[userId]; !ok {
		return errors.New("User does not exist")
	}
	delete(userDB, userId)
	return nil
}

func (source) UpdateUser(userId string, password string) error {
	if _, ok := userDB[userId]; !ok {
		return errors.New("User does not exist")
	}
	userDB[userId] = password
	return nil
}

func (source) GetPass(userId string) (string, error) {
	if _, ok := userDB[userId]; !ok {
		return "", errors.New("User does not exist")
	}
	return userDB[userId], nil
}
