package repository

import (
	"errors"
)

func (s *service) Validate(tokenString string) (string, error) {
	return s.jwt.Verify(tokenString)
}

func (s *service) LogIn(userId, passwd string) (string, error) {
	pass, err := s.dataSource.GetPass(userId)
	if err != nil {
		return "", err
	}
	if pass != passwd {
		return "", errors.New("Wrong password")
	}
	token, err := s.jwt.NewToken(userId)
	return token, err
}

func (s *service) SignUp(userId, password string) (string, error) {
	err := s.dataSource.AddUser(userId, password)
	if err != nil {
		return "", err
	}
	return s.LogIn(userId, password)
}
