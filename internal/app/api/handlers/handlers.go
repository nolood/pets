package handlers

import (
	"go.uber.org/zap"
	"pets/internal/app/api/handlers/user"
	"pets/internal/app/api/services"
)

type Handlers struct {
	User user.Handler
}

func New(log *zap.Logger, servs *services.Services) *Handlers {
	return &Handlers{
		User: user.New(log, servs.User),
	}
}
