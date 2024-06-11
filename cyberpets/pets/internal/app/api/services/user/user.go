package user

import (
	"context"
	"cyberpets/pets/internal/domain/models"
	"cyberpets/pets/internal/repositories/user"

	"go.uber.org/zap"
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
