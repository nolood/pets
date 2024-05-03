package handlers

import (
	"go.uber.org/zap"
	"pets/internal/app/api/handlers/auth"
	"pets/internal/app/api/handlers/farm"
	"pets/internal/app/api/handlers/incubator"
	"pets/internal/app/api/handlers/user"
	"pets/internal/app/api/services"
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
