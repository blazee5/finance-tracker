package service

import storage "github.com/blazee5/finance-tracker/internal/storage/mongodb"

type Service struct {
	Storage *storage.Storage
}

func NewService(storage *storage.Storage) *Service {
	return &Service{Storage: storage}
}
