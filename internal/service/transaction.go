package service

import (
	"context"
	"github.com/blazee5/finance-tracker/internal/models"
)

func (s *Service) GetTransactions(id string) ([]models.Transaction, error) {
	return s.Storage.TransactionDAO.GetTransactions(context.Background(), id)
}

func (s *Service) CreateTransaction(transaction models.Transaction) (interface{}, error) {
	return s.Storage.TransactionDAO.Create(context.Background(), transaction)
}

func (s *Service) GetTransaction(id string) (models.Transaction, error) {
	return s.Storage.TransactionDAO.GetTransaction(context.Background(), id)
}

func (s *Service) UpdateTransaction(transaction models.Transaction) (int64, error) {
	return s.Storage.TransactionDAO.Update(context.Background(), transaction)
}

func (s *Service) DeleteTransaction(id string) (int64, error) {
	return s.Storage.TransactionDAO.Delete(context.Background(), id)
}
