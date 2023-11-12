package service

import (
	"context"
	"github.com/blazee5/finance-tracker/internal/domain"
	"github.com/blazee5/finance-tracker/internal/models"
	"github.com/blazee5/finance-tracker/internal/repository"
	"go.uber.org/zap"
)

type Auth interface {
	CreateUser(ctx context.Context, user domain.SignUpRequest) (string, error)
	GenerateToken(ctx context.Context, email, password string) (string, error)
}

type User interface {
	Update(ctx context.Context, id string, input domain.UpdateUserRequest) error
	GetUserById(ctx context.Context, id string) (models.ShortUser, error)
	UploadAvatar(ctx context.Context, id string, file string) error
}

type Transaction interface {
	GetTransactions(ctx context.Context, id string) ([]models.Transaction, error)
	CreateTransaction(ctx context.Context, userId string, transaction domain.Transaction) (string, error)
	GetTransaction(ctx context.Context, id string) (models.Transaction, error)
	UpdateTransaction(ctx context.Context, id string, transaction domain.Transaction) error
	DeleteTransaction(ctx context.Context, id string) error
	AnalyzeTransactions(ctx context.Context, id string) (models.Analyze, error)
}

type Service struct {
	log         *zap.SugaredLogger
	repo        *repository.Repository
	Auth        Auth
	User        User
	Transaction Transaction
}

func NewService(log *zap.SugaredLogger, repo *repository.Repository) *Service {
	return &Service{
		Auth:        NewAuthService(log, repo),
		User:        NewUserService(log, repo),
		Transaction: NewTransactionService(log, repo),
	}
}
