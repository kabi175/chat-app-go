package service

import (
	"github.com/kabi175/chat-app-go/messager/domain"
	"github.com/kabi175/chat-app-go/messager/domain/apperrors"
)

type DefaultUserService struct {
	userRepo     domain.UserRepo
	tokenService domain.TokenService
}

func NewDefaultUserService(userRepo domain.UserRepo, tokenService domain.TokenService) domain.UserService {
	return &DefaultUserService{userRepo: userRepo, tokenService: tokenService}
}

func (s *DefaultUserService) LogIn(user *domain.User) (string, error) {
	emailMatchedUser, err := s.userRepo.GetByEmail(user.Email)
	if err != nil {
		return "", err
	}
	if emailMatchedUser == nil || emailMatchedUser.Password != user.Password {
		return "", apperrors.NewNotFoundError("User email and password not found")
	}
	token, err := s.tokenService.GenerateToken(emailMatchedUser)
	return token, err
}

func (s *DefaultUserService) SignUp(user *domain.User) (*domain.User, error) {
	user.ID = 0
	err := s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}
	user, err = s.userRepo.GetByEmail(user.Email)
	return user, err
}

func (s *DefaultUserService) GetByID(id uint) (*domain.User, error) {
	user, err := s.userRepo.GetByID(id)
	return user, err
}

func (s *DefaultUserService) GetByEmail(userEmail string) (*domain.User, error) {
	user, err := s.userRepo.GetByEmail(userEmail)
	return user, err
}

func (s *DefaultUserService) Delete(user *domain.User) error {
	err := s.userRepo.DeleteByID(user.ID)
	return err
}
