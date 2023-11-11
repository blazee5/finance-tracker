package repository

import (
	"context"
	"github.com/blazee5/finance-tracker/internal/config"
	"github.com/blazee5/finance-tracker/internal/domain"
	"github.com/blazee5/finance-tracker/internal/models"
	"github.com/blazee5/finance-tracker/internal/repository/mongodb"
	redisRepo "github.com/blazee5/finance-tracker/internal/repository/redis"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	db               *mongo.Client
	rdb              *redis.Client
	User             UserRepository
	Transaction      TransactionRepository
	UserRedis        UserRedisRepository
	TransactionRedis TransactionRedisRepository
}

type UserRedisRepository interface {
	GetByIdCtx(ctx context.Context, key string) (*models.ShortUser, error)
	SetUserCtx(ctx context.Context, key string, seconds int, user models.ShortUser) error
	DeleteUserCtx(ctx context.Context, key string) error
}

type TransactionRedisRepository interface {
	GetByIdCtx(ctx context.Context, key string) (*models.Transaction, error)
	SetTransactionCtx(ctx context.Context, key string, seconds int, transaction models.Transaction) error
	DeleteTransactionCtx(ctx context.Context, key string) error
}

type UserRepository interface {
	Create(ctx context.Context, input domain.SignUpRequest) (string, error)
	GetUser(ctx context.Context, email, password string) (models.User, error)
	GetUserById(ctx context.Context, id string) (models.ShortUser, error)
	UpdateUser(ctx context.Context, id string, input domain.UpdateUserRequest) error
}

type TransactionRepository interface {
	Create(ctx context.Context, user models.ShortUser, transaction domain.Transaction) (string, error)
	GetTransactions(ctx context.Context, userID string) ([]models.Transaction, error)
	GetTransaction(ctx context.Context, id string) (models.Transaction, error)
	Update(ctx context.Context, id string, transaction domain.Transaction) (int, error)
	Delete(ctx context.Context, id string) error
	GetAnalyze(ctx context.Context, id string) (models.Analyze, error)
}

func NewRepository(cfg *config.Config, db *mongo.Client, rdb *redis.Client) *Repository {
	return &Repository{
		db:               db,
		rdb:              rdb,
		User:             mongodb.NewUserRepository(cfg, db),
		Transaction:      mongodb.NewTransactionRepository(cfg, db),
		UserRedis:        redisRepo.NewUserRedisRepository(rdb),
		TransactionRedis: redisRepo.NewTransactionRedisRepository(rdb),
	}
}
