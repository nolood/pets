package incubator

import (
	"context"
	"fmt"
	"pets/internal/domain/models"
	"pets/internal/repositories"
)

type Service interface {
	Get(ctx context.Context) (models.Incubator, error)
	Clear(ctx context.Context) (models.Incubator, error)
	SetEgg(ctx context.Context, eggID uint64) (models.Incubator, error)
	OpenEgg(ctx context.Context, eggID uint64) (models.Pet, error)
}

type incubatorService struct {
	repos *repositories.Repositories
}

func New(repos *repositories.Repositories) Service {
	return &incubatorService{repos: repos}
}

// SetEgg - eggID this is id of UserEgg, not egg
func (s *incubatorService) SetEgg(ctx context.Context, eggID uint64) (models.Incubator, error) {
	const op = "service.incubator.setEgg"

	userEgg, err := s.repos.Incubator.GetUserEgg(ctx, eggID)
	if err != nil {
		return models.Incubator{}, fmt.Errorf("%s: %w", op, err)
	}

	if userEgg.Egg == nil {
		return models.Incubator{}, fmt.Errorf("%s: egg-not-found", op)
	}

	incubator, err := s.repos.Incubator.SetEgg(ctx, userEgg.Egg.ID)
	if err != nil {
		return models.Incubator{}, fmt.Errorf("%s: %w", op, err)
	}

	return incubator, nil
}

func (s *incubatorService) Clear(ctx context.Context) (models.Incubator, error) {
	const op = "service.incubator.clear"

	incubator, err := s.repos.Incubator.Clear(ctx)
	if err != nil {
		return models.Incubator{}, fmt.Errorf("%s: %w", op, err)
	}

	return incubator, nil
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
