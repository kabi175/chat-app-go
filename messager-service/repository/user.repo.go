package repository

import (
	"errors"

	"github.com/kabi175/chat-app-go/messager/domain"
	"github.com/kabi175/chat-app-go/messager/domain/apperrors"
	"gorm.io/gorm"
)

type PostgresUserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) (domain.UserRepo, error) {
	err := db.AutoMigrate(&domain.User{})
	if err != nil {
		return nil, apperrors.NewInternalServerError(err.Error())
	}
	return &PostgresUserRepo{db: db}, nil
}

func (u *PostgresUserRepo) Create(user *domain.User) error {
	tx := u.db.Create(user)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrInvalidField) {
			return apperrors.NewConflictError("user already exists")
		}
		return apperrors.NewInternalServerError(tx.Error.Error())
	}
	return nil
}

func (u *PostgresUserRepo) GetByID(id uint) (*domain.User, error) {
	user := &domain.User{}
	u.db.Where(&domain.User{ID: id}).First(user)
	return user, nil
}

func (u *PostgresUserRepo) GetByEmail(email string) (*domain.User, error) {
	user := &domain.User{}
	u.db.Where(&domain.User{Email: email}).First(&user)
	return user, nil
}

func (u *PostgresUserRepo) DeleteByID(id uint) error {
	result := u.db.Delete(&domain.User{ID: id})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return apperrors.NewNotFoundError("user not found")
		}
		return apperrors.NewInternalServerError(result.Error.Error())
	}
	return nil
}

func (u *PostgresUserRepo) DeleteByEmail(email string) error {
	result := u.db.Where(&domain.User{Email: email}).Delete(&domain.User{})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return apperrors.NewNotFoundError("user not found")
		}
		return apperrors.NewInternalServerError(result.Error.Error())
	}
	return nil
}
