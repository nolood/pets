package user

import (
	"context"
	"go.uber.org/zap"
	"pets/internal/app/api/repositories/user"
	"pets/internal/domain/models"
)

type Service interface {
	Create(ctx context.Context, entity models.User) error
}

type userService struct {
	repo user.Repository
	log  *zap.Logger
}

func New(log *zap.Logger, repo user.Repository) Service {
	return &userService{repo: repo, log: log}
}

func (s *userService) Create(ctx context.Context, entity models.User) error {
	s.log.Info("SERVICE: create user")
	return s.repo.Create(ctx, entity)
}
