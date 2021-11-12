package repository

import (
	"context"

	"github.com/kabi175/chat-app-go/messager/domain"
)

type RedisMessageRepo struct{}

func NewResidMessageRepo() domain.MessageRepo {
	return &RedisMessageRepo{}
}

func (r *RedisMessageRepo) Producer(message *domain.Message, ctx context.Context) error {
	panic("not implemented")
}

func (r *RedisMessageRepo) Consumer(user *domain.User, ctx context.Context) (<-chan domain.Message, error) {
	panic("not implemented")
}
