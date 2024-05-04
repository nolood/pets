package incubator

import (
	"context"
	"fmt"
	"pets/internal/domain/models"
	"pets/internal/repositories"
)

type Service interface {
	Get(ctx context.Context) (models.Incubator, error)
	OpenEgg(ctx context.Context, eggID uint64) (models.Pet, error)
}

type incubatorService struct {
	repos *repositories.Repositories
}

func New(repos *repositories.Repositories) Service {
	return &incubatorService{repos: repos}
}

func (s *incubatorService) OpenEgg(ctx context.Context, eggID uint64) (models.Pet, error) {
	const op = "service.incubator.openEgg"

	egg, err := s.repos.Incubator.RemoveEgg(ctx, eggID)
	if err != nil {
		return models.Pet{}, fmt.Errorf("%s: %w", op, err)
	}

	pet, err := s.repos.Pet.AddRandomPetToUser(ctx, egg.Rarity)
	if err != nil {
		return models.Pet{}, fmt.Errorf("%s: %w", op, err)
	}

	return pet, nil
}

func (s *incubatorService) Get(ctx context.Context) (models.Incubator, error) {
	const op = "service.incubator.get"

	incubatorInst, err := s.repos.Incubator.Get(ctx)
	if err != nil {
		return models.Incubator{}, fmt.Errorf("%s: %w", op, err)
	}

	if incubatorInst.ID != 0 {
		return incubatorInst, nil
	}

	newIncubator, err := s.repos.Incubator.Init(ctx)
	if err != nil {
		return models.Incubator{}, fmt.Errorf("%s: %w", op, err)
	}

	return newIncubator, nil
}
