package farm

import (
	"context"
	"fmt"
	"pets/internal/domain/models"
	"pets/internal/repositories/farm"
)

type Service interface {
	Get(ctx context.Context) (models.Farm, error)
}

type farmService struct {
	repo farm.Repository
}

func New(repo farm.Repository) Service {
	return &farmService{repo: repo}
}

func (s *farmService) Get(ctx context.Context) (models.Farm, error) {
	const op = "service.farm.get"

	isFarm, err := s.repo.Get(ctx)
	if err != nil {
		return models.Farm{}, fmt.Errorf("%s: %w", op, err)
	}

	if isFarm.ID != 0 {
		return isFarm, nil
	}

	newFarm, err := s.repo.Init(ctx)
	if err != nil {
		return models.Farm{}, fmt.Errorf("%s: %w", op, err)
	}

	return newFarm, nil

}
