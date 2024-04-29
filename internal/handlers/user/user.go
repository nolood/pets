package user

import (
	"pets/internal/services/user"
)

type Handler struct {
	service *user.Service
}

func New(service *user.Service) *Handler {
	return &Handler{
		service: service,
	}
}
