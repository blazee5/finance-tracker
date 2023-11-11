package redis

import (
	"context"
	"encoding/json"
	"github.com/blazee5/finance-tracker/internal/models"
	"github.com/redis/go-redis/v9"
	"time"
)

type TransactionRedisRepository struct {
	rdb *redis.Client
}

func NewTransactionRedisRepository(rdb *redis.Client) *TransactionRedisRepository {
	return &TransactionRedisRepository{rdb: rdb}
}

func (repo *TransactionRedisRepository) GetByIdCtx(ctx context.Context, key string) (*models.Transaction, error) {
	var transaction *models.Transaction

	res, err := repo.rdb.Get(ctx, "transaction:"+key).Bytes()

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &transaction)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (repo *TransactionRedisRepository) SetTransactionCtx(ctx context.Context, key string, seconds int, transaction models.Transaction) error {
	transactionBytes, err := json.Marshal(&transaction)

	if err != nil {
		return err
	}

	err = repo.rdb.Set(ctx, "transaction:"+key, transactionBytes, time.Second*time.Duration(seconds)).Err()

	if err != nil {
		return err
	}

	return nil
}

func (repo *TransactionRedisRepository) DeleteTransactionCtx(ctx context.Context, key string) error {
	if err := repo.rdb.Del(ctx, "transaction:"+key).Err(); err != nil {
		return err
	}

	return nil
}
