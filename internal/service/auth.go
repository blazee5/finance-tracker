package service

import (
	"context"
	"github.com/blazee5/finance-tracker/internal/domain"
	"github.com/blazee5/finance-tracker/internal/repository"
	"github.com/blazee5/finance-tracker/lib/auth"
	"go.uber.org/zap"
)

type AuthService struct {
	log  *zap.SugaredLogger
	repo *repository.Repository
}

func NewAuthService(log *zap.SugaredLogger, repo *repository.Repository) *AuthService {
	return &AuthService{log: log, repo: repo}
}

func (s *AuthService) CreateUser(ctx context.Context, user domain.SignUpRequest) (string, error) {
	user.Password = auth.GenerateHashPassword(user.Password)

	id, err := s.repo.User.Create(ctx, user)

	if err != nil {
		return "", nil
	}

	return auth.GenerateToken(id)
}

func (s *AuthService) GenerateToken(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.User.GetUser(ctx, email, auth.GenerateHashPassword(password))

	if err != nil {
		return "", err
	}

	return auth.GenerateToken(user.ID.Hex())
}
