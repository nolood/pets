package auth

import (
	"context"
	"go.uber.org/zap"
	"pets/internal/domain/models"
	"pets/internal/repositories/user"
)

type Service interface {
	Login(ctx context.Context, entity models.User) (models.User, error)
	Register(ctx context.Context, entity models.User) (models.User, error)
}

type authService struct {
	repo user.Repository
	log  *zap.Logger
}

func New(log *zap.Logger, repo user.Repository) Service {
	return &authService{repo: repo, log: log}
}

func (s *authService) Login(ctx context.Context, entity models.User) (models.User, error) {

	return models.User{}, nil
}

func (s *authService) Register(ctx context.Context, entity models.User) (models.User, error) {

	return s.repo.Create(ctx, entity)
}
