package repository

import (
	"errors"

	"github.com/kabi175/chat-app-go/messager/domain"
	"github.com/kabi175/chat-app-go/messager/domain/apperrors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresUserRepo struct {
	db *gorm.DB
}

func DB() (*gorm.DB, error) {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
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
	result := u.db.Where(&domain.User{ID: id}).Delete(&domain.User{})
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
