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

func (s *TransactionService) GetTransactions(ctx context.Context, id, category string) ([]models.Transaction, error) {
	return s.repo.Transaction.GetTransactions(ctx, id, category)
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
	cachedTransaction, err := s.repo.TransactionRedis.GetByIdCtx(ctx, id)

	if err != nil {
		s.log.Infof("error while get transaction from redis: %v", err)
	}

	if cachedTransaction != nil {
		return *cachedTransaction, nil
	}

	transaction, err := s.repo.Transaction.GetTransaction(ctx, id)

	if err != nil {
		return models.Transaction{}, err
	}

	if err := s.repo.TransactionRedis.SetTransactionCtx(ctx, id, 3600, transaction); err != nil {
		s.log.Infof("error while save transaction in redis: %v", err)
	}

	return transaction, nil
}

func (s *TransactionService) UpdateTransaction(ctx context.Context, id string, transaction domain.Transaction) error {
	_, err := s.repo.Transaction.Update(ctx, id, transaction)

	if err != nil {
		return err
	}

	if err := s.repo.TransactionRedis.DeleteTransactionCtx(ctx, id); err != nil {
		return err
	}

	return nil
}

func (s *TransactionService) DeleteTransaction(ctx context.Context, id string) error {
	return s.repo.Transaction.Delete(ctx, id)
}

func (s *TransactionService) AnalyzeTransactions(ctx context.Context, id string) (models.Analyze, error) {
	return s.repo.Transaction.GetAnalyze(ctx, id)
}
