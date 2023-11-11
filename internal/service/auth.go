package service

import (
	"context"
	"github.com/blazee5/finance-tracker/internal/domain"
	"github.com/blazee5/finance-tracker/internal/models"
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

func (s *AuthService) GetUserById(ctx context.Context, id string) (models.ShortUser, error) {
	cachedUser, err := s.repo.UserRedis.GetByIdCtx(ctx, id)

	if err != nil {
		s.log.Infof("error while get user from redis: %v", err)
	}

	if cachedUser != nil {
		return *cachedUser, nil
	}

	user, err := s.repo.User.GetUserById(ctx, id)

	if err != nil {
		return models.ShortUser{}, err
	}

	if err := s.repo.UserRedis.SetUserCtx(ctx, id, 3600, user); err != nil {
		s.log.Infof("error while save user in redis: %v", err)
	}

	return user, nil
}
