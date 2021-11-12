package repository

import (
	"github.com/kabi175/chat-app-go/messager/domain"
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
	db.AutoMigrate(&domain.User{})
	return &PostgresUserRepo{db: db}, nil
}

func (u *PostgresUserRepo) Create(user *domain.User) error {
	return nil
}
func (u *PostgresUserRepo) GetByID(id uint) (*domain.User, error) {
	panic("not implemented")
}
func (u *PostgresUserRepo) GetByEmail(name string) (*domain.User, error) {
	panic("not implemented")
}
func (u *PostgresUserRepo) DeleteByID(id uint) error {
	panic("not implemented")
}
func (u *PostgresUserRepo) DeleteByEmail(name string) error {
	panic("not implemented")
}
