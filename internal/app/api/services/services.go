package services

import (
	"go.uber.org/zap"
	"pets/internal/app/api/repositories"
	"pets/internal/app/api/services/user"
)

type Services struct {
	User user.Service
}

func New(log *zap.Logger, repos *repositories.Repositories) *Services {
	return &Services{
		User: user.New(log, repos.User),
	}
}
