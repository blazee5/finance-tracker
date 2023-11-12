package service

import (
	"context"
	"github.com/blazee5/finance-tracker/internal/domain"
	"github.com/blazee5/finance-tracker/internal/models"
	"github.com/blazee5/finance-tracker/internal/repository"
	"go.uber.org/zap"
)

type UserService struct {
	log  *zap.SugaredLogger
	repo *repository.Repository
}

func NewUserService(log *zap.SugaredLogger, repo *repository.Repository) *UserService {
	return &UserService{log: log, repo: repo}
}

func (s *UserService) GetUserById(ctx context.Context, id string) (models.ShortUser, error) {
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

func (s *UserService) Update(ctx context.Context, id string, input domain.UpdateUserRequest) error {
	err := s.repo.User.UpdateUser(ctx, id, input)

	if err != nil {
		return err
	}

	if err := s.repo.UserRedis.DeleteUserCtx(ctx, id); err != nil {
		s.log.Infof("error while save user in redis: %v", err)
	}

	return nil
}

func (s *UserService) UploadAvatar(ctx context.Context, id string, file string) error {
	err := s.repo.User.UploadAvatar(ctx, id, file)

	if err != nil {
		return err
	}

	err = s.repo.UserRedis.DeleteUserCtx(ctx, id)

	if err != nil {
		s.log.Infof("error while delete user from redis: %v", err)
	}

	return nil
}
