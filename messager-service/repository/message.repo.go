package repository

import (
	"context"
	"encoding/json"
	"strconv"

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
	r.db.RPush(ctx, strconv.Itoa(int(message.To)), value)
	return nil
}

func (r *RedisMessageRepo) Consumer(user *domain.User, ctx context.Context) (<-chan domain.MessageChan, error) {
	ch := make(chan domain.MessageChan, 5)

	handleError := func(err error) {
		ch <- domain.MessageChan{Error: err}
	}
	go func(userID uint, ctx context.Context) {

		defer func() {
			close(ch)
		}()

		message := domain.Message{}

		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			result, err := r.db.BLPop(ctx, 0, strconv.FormatUint(uint64(userID), 10)).Result()
			if err != nil {
				handleError(err)
				return
			}
			err = json.Unmarshal([]byte(result[0]), &message)
			if err != nil {
				handleError(err)
				return
			}
			ch <- domain.MessageChan{Message: message}
		}
	}(user.ID, ctx)

	return ch, nil
}
