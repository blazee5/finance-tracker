package redis

import (
	"context"
	"encoding/json"
	"github.com/blazee5/finance-tracker/internal/models"
	"github.com/redis/go-redis/v9"
	"time"
)

type UserRedisRepository struct {
	rdb *redis.Client
}

func NewUserRedisRepository(rdb *redis.Client) *UserRedisRepository {
	return &UserRedisRepository{rdb: rdb}
}

func (repo *UserRedisRepository) GetByIdCtx(ctx context.Context, key string) (*models.ShortUser, error) {
	var user *models.ShortUser

	res, err := repo.rdb.Get(ctx, "user:"+key).Bytes()

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRedisRepository) SetUserCtx(ctx context.Context, key string, seconds int, user models.ShortUser) error {
	userBytes, err := json.Marshal(&user)

	if err != nil {
		return err
	}

	err = repo.rdb.Set(ctx, "user:"+key, userBytes, time.Second*time.Duration(seconds)).Err()

	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRedisRepository) DeleteUserCtx(ctx context.Context, key string) error {
	if err := repo.rdb.Del(ctx, "user:"+key).Err(); err != nil {
		return err
	}

	return nil
}
