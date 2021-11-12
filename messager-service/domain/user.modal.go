package domain

import (
	"encoding/json"

	"github.com/kabi175/chat-app-go/messager/domain/apperrors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primarykey"`
	DisplayName string `json:"displayName" gorm:"not null"`
	Password    string `json:"password" gorm:"not null"`
	Email       string `json:"email" gorm:"unique;not null"`
	CreatedAt   int64  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   int64  `json:"updatedAt" gorm:"autoUpdateTime"`
}

func (u *User) String() (string, error) {
	str, err := json.Marshal(u)
	if err != nil {
		return "", apperrors.NewInternalServerError(err.Error())
	}
	return string(str), nil
}
