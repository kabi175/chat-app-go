package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primarykey"`
	DisplayName string `json:"displayName" gorm:"not null"`
	Password    string `json:"password" gorm:"not null"`
	Email       string `json:"email" gorm:"unique;not null"`
	CreatedAt   int64  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   int64  `json:"updatedAt" gorm:"autoUpdateTime"`
}
