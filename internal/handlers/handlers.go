package handlers

import (
	"pets/internal/handlers/user"
	"pets/internal/services"
)

type Handlers struct {
	User *user.Handler
}

func New(servs *services.Services) *Handlers {
	return &Handlers{
		User: user.New(servs.User),
	}
}
