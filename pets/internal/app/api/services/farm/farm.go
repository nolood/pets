package farm

import (
	"context"
	"fmt"
	"pets/internal/domain/models"
	"pets/internal/repositories/farm"
)

type Service interface {
	Get(ctx context.Context) (models.Farm, error)
	SetPet(ctx context.Context, petID uint64, slotID uint64) (models.Farm, error)
	RemovePet(ctx context.Context) (models.Farm, error)
}

type farmService struct {
	repo farm.Repository
}

func New(repo farm.Repository) Service {
	return &farmService{repo: repo}
}

func (s *farmService) RemovePet(ctx context.Context) (models.Farm, error) {
	// TODO: implement me
	return models.Farm{}, nil
}

func (s *farmService) SetPet(ctx context.Context, petID uint64, slotID uint64) (models.Farm, error) {
	const op = "service.farm.setPet"

	farmInst, err := s.repo.SetPet(ctx, petID, slotID)
	if err != nil {
		return models.Farm{}, fmt.Errorf("%s: %w", op, err)
	}
	return farmInst, nil
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
