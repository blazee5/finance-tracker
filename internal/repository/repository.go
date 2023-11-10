package repository

import (
	"context"
	"github.com/blazee5/finance-tracker/internal/config"
	"github.com/blazee5/finance-tracker/internal/domain"
	"github.com/blazee5/finance-tracker/internal/models"
	"github.com/blazee5/finance-tracker/internal/repository/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	db          *mongo.Client
	User        UserRepository
	Transaction TransactionRepository
}

type UserRepository interface {
	Create(ctx context.Context, user models.User) (string, error)
	GetUser(ctx context.Context, email, password string) (models.User, error)
	GetUserById(ctx context.Context, id string) (models.ShortUser, error)
}

type TransactionRepository interface {
	Create(ctx context.Context, user models.ShortUser, transaction domain.Transaction) (string, error)
	GetTransactions(ctx context.Context, userID string) ([]models.Transaction, error)
	GetTransaction(ctx context.Context, id string) (models.Transaction, error)
	Update(ctx context.Context, id string, transaction domain.Transaction) (int, error)
	Delete(ctx context.Context, id string) error
	GetAnalyze(ctx context.Context, id string) (models.Analyze, error)
}

func NewRepository(cfg *config.Config, db *mongo.Client) *Repository {
	return &Repository{
		db:          db,
		User:        mongodb.NewUserRepository(cfg, db),
		Transaction: mongodb.NewTransactionRepository(cfg, db),
	}
}
