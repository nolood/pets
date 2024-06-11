package services

import (
	"cyberpets/pets/internal/app/api/services/auth"
	"cyberpets/pets/internal/app/api/services/farm"
	"cyberpets/pets/internal/app/api/services/incubator"
	"cyberpets/pets/internal/app/api/services/user"
	"cyberpets/pets/internal/config"
	"cyberpets/pets/internal/repositories"

	"go.uber.org/zap"
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
		Incubator: incubator.New(repos),
	}
}
