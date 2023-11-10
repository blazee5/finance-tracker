package service

import (
	"context"
	"github.com/blazee5/finance-tracker/internal/models"
	"github.com/blazee5/finance-tracker/lib/auth"
)

func (s *Service) CreateUser(ctx context.Context, user models.User) (string, error) {
	user.Password = auth.GenerateHashPassword(user.Password)

	id, err := s.Storage.UserDAO.Create(ctx, user)

	if err != nil {
		return "", nil
	}

	return auth.GenerateToken(id)
}

func (s *Service) GenerateToken(ctx context.Context, email, password string) (string, error) {
	user, err := s.Storage.UserDAO.GetUser(ctx, email, auth.GenerateHashPassword(password))

	if err != nil {
		return "", err
	}

	return auth.GenerateToken(user.ID.Hex())
}

func (s *Service) GetUserById(ctx context.Context, id string) (models.ShortUser, error) {
	return s.Storage.GetUserById(ctx, id)
}
