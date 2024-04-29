package user

import (
	"context"
	"pets/internal/repositories/user"
)

type Service struct {
	Repo user.Repository
}

func New(repo user.Repository) *Service {
	return &Service{Repo: repo}
}

func (s *Service) Create(ctx context.Context, entity interface{}) error {
	return s.Repo.Create(ctx, entity)
}
