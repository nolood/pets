package services

import (
	"pets/internal/repositories"
	"pets/internal/services/user"
)

type Services struct {
	User *user.Service
}

func New(repos *repositories.Repositories) *Services {
	return &Services{
		User: user.New(repos.User),
	}
}
