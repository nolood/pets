package user

import (
	"context"
	"go.uber.org/zap"
	"pets/internal/domain/models"
	"pets/internal/repositories/user"
)

type Service interface {
	CreateOrUpdate(ctx context.Context, entity models.User) (models.User, error)
}

type userService struct {
	repo user.Repository
	log  *zap.Logger
}

func New(log *zap.Logger, repo user.Repository) Service {
	return &userService{repo: repo, log: log}
}

func (s *userService) CreateOrUpdate(ctx context.Context, entity models.User) (models.User, error) {
	return s.repo.CreateOrUpdate(ctx, entity)
}
