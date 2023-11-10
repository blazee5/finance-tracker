package service

import (
	"context"
	"github.com/blazee5/finance-tracker/internal/domain"
	"github.com/blazee5/finance-tracker/internal/models"
	"github.com/blazee5/finance-tracker/internal/repository"
	"go.uber.org/zap"
)

type TransactionService struct {
	log  *zap.SugaredLogger
	repo *repository.Repository
}

func NewTransactionService(log *zap.SugaredLogger, repo *repository.Repository) *TransactionService {
	return &TransactionService{log: log, repo: repo}
}

func (s *TransactionService) GetTransactions(ctx context.Context, id string) ([]models.Transaction, error) {
	return s.repo.Transaction.GetTransactions(ctx, id)
}

func (s *TransactionService) CreateTransaction(ctx context.Context, userId string, transaction domain.Transaction) (string, error) {
	user, err := s.repo.User.GetUserById(ctx, userId)

	if err != nil {
		return "", err
	}

	id, err := s.repo.Transaction.Create(ctx, user, transaction)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *TransactionService) GetTransaction(ctx context.Context, id string) (models.Transaction, error) {
	return s.repo.Transaction.GetTransaction(ctx, id)
}

func (s *TransactionService) UpdateTransaction(ctx context.Context, id string, transaction domain.Transaction) (int, error) {
	return s.repo.Transaction.Update(ctx, id, transaction)
}

func (s *TransactionService) DeleteTransaction(ctx context.Context, id string) error {
	return s.repo.Transaction.Delete(ctx, id)
}

func (s *TransactionService) AnalyzeTransactions(ctx context.Context, id string) (models.Analyze, error) {
	return s.repo.Transaction.GetAnalyze(ctx, id)
}
