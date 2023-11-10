package service

import (
	"context"
	"github.com/blazee5/finance-tracker/internal/domain"
	"github.com/blazee5/finance-tracker/internal/models"
	storage "github.com/blazee5/finance-tracker/internal/storage/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	Storage *Storage
}

type Storage struct {
	Db *mongo.Client
	UserDAO
	TransactionDAO
}

//go:generate mockery --name UserDAO
type UserDAO interface {
	Create(ctx context.Context, user models.User) (string, error)
	GetUser(ctx context.Context, email, password string) (models.User, error)
	GetUserById(ctx context.Context, id string) (models.ShortUser, error)
}

//go:generate mockery --name TransactionDAO
type TransactionDAO interface {
	Create(ctx context.Context, user models.ShortUser, transaction domain.Transaction) (string, error)
	GetTransactions(ctx context.Context, userID string) ([]models.Transaction, error)
	GetTransaction(ctx context.Context, id string) (models.Transaction, error)
	Update(ctx context.Context, id string, transaction domain.Transaction) (int, error)
	Delete(ctx context.Context, id string) error
	GetAnalyze(ctx context.Context, id string) (models.Analyze, error)
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		Storage: &Storage{
			Db:             storage.Db,
			UserDAO:        storage.UserDAO,
			TransactionDAO: storage.TransactionDAO,
		},
	}
}
