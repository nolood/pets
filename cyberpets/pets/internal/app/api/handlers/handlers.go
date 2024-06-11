package handlers

import (
	"cyberpets/pets/internal/app/api/handlers/auth"
	"cyberpets/pets/internal/app/api/handlers/farm"
	"cyberpets/pets/internal/app/api/handlers/incubator"
	"cyberpets/pets/internal/app/api/handlers/user"
	"cyberpets/pets/internal/app/api/services"

	"go.uber.org/zap"
)

type Handlers struct {
	User      user.Handler
	Auth      auth.Handler
	Farm      farm.Handler
	Incubator incubator.Handler
}

func New(log *zap.Logger, servs *services.Services, token string) *Handlers {
	return &Handlers{
		User:      user.New(log, servs.User),
		Auth:      auth.New(log, servs.Auth, token),
		Farm:      farm.New(log, servs.Farm),
		Incubator: incubator.New(log, servs.Incubator),
	}
}
