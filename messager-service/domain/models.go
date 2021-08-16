package domain

import (
	"encoding/json"
)

const (
	TypeMessage = iota
	TypeStatus
	TypeAck
)

type UserStatus struct {
	UserName string `json:"username"`
	Status   string `json:"status"`
}

func (us *UserStatus) Bytes() ([]byte, error) {
	return json.Marshal(us)
}

func NewUserStatus(obj interface{}) (*UserStatus, error) {
	switch o := obj.(type) {
	case string:
		var status UserStatus
		err := json.Unmarshal([]byte(o), &status)
		return &status, err
	}
	return nil, nil
}
