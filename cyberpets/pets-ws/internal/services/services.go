package services

import (
	"cyberpets/pets-ws/internal/app/clicker"
	"cyberpets/pets-ws/internal/services/router"

	"go.uber.org/zap"
)

type Services struct {
	Router router.Service
}

func New(log *zap.Logger, clicker *clicker.App) *Services {
	return &Services{
		Router: router.New(clicker, log),
	}
}
