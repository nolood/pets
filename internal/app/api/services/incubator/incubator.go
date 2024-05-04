package incubator

import (
	"context"
	"fmt"
	"log"
	"pets/internal/domain/models"
	"pets/internal/repositories/incubator"
)

type Service interface {
	Get(ctx context.Context) (models.Incubator, error)
}

type incubatorService struct {
	repo incubator.Repository
}

func New(repo incubator.Repository) Service {
	return &incubatorService{repo: repo}
}

func (s *incubatorService) Get(ctx context.Context) (models.Incubator, error) {
	const op = "service.incubator.get"

	log.Println(op)

	incubatorInst, err := s.repo.Get(ctx)
	if err != nil {
		return models.Incubator{}, fmt.Errorf("%s: %w", op, err)
	}

	if incubatorInst.ID != 0 {
		return incubatorInst, nil
	}

	newIncubator, err := s.repo.Init(ctx)
	if err != nil {
		return models.Incubator{}, fmt.Errorf("%s: %w", op, err)
	}

	return newIncubator, nil
}
