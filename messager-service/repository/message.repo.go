package repository

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/kabi175/chat-app-go/messager/domain"
)

type RedisMessageRepo struct {
	db *redis.Client
}

func NewResidMessageRepo(db *redis.Client) domain.MessageRepo {
	return &RedisMessageRepo{db: db}
}

func (r *RedisMessageRepo) Producer(message *domain.Message, ctx context.Context) error {
	value, err := message.String()
	if err != nil {
		return err
	}
	r.db.RPush(ctx, string(message.To), value)
	return nil
}

func (r *RedisMessageRepo) Consumer(user *domain.User, ctx context.Context) (<-chan domain.Message, context.Context, error) {
	ch := make(chan domain.Message, 5)
	childCtx, cancel := context.WithCancel(context.Background())

	go func(userID uint, ctx context.Context) {

		defer func() {
			cancel()
			close(ch)
		}()

		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			result, err := r.db.BLPop(ctx, 0, string(userID)).Result()
			if err != nil {
				return
			}
			message := domain.Message{}
			err = json.Unmarshal([]byte(result[0]), &message)
			if err != nil {
				return
			}
			ch <- message
		}
	}(user.ID, ctx)

	return ch, childCtx, nil
}
