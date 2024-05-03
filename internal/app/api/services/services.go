package services

import (
	"go.uber.org/zap"
	"pets/internal/app/api/services/auth"
	"pets/internal/app/api/services/farm"
	"pets/internal/app/api/services/incubator"
	"pets/internal/app/api/services/user"
	"pets/internal/config"
	"pets/internal/repositories"
)

type Services struct {
	User      user.Service
	Auth      auth.Service
	Farm      farm.Service
	Incubator incubator.Service
}

func New(log *zap.Logger, repos *repositories.Repositories, cfg *config.Config) *Services {
	return &Services{
		User:      user.New(log, repos.User),
		Auth:      auth.New(log, repos.User, cfg),
		Farm:      farm.New(repos.Farm),
		Incubator: incubator.New(repos.Incubator),
	}
}
