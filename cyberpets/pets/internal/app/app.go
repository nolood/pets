package app

import (
	"cyberpets/pets/internal/app/api"
	"cyberpets/pets/internal/config"
	"cyberpets/pets/internal/storage/postgres"

	"go.uber.org/zap"
)

type App struct {
	Api *api.Api
}

func New(log *zap.Logger, port int, cfg *config.Config) *App {
	storage := postgres.New(cfg.Storage)

	apiSrv := api.New(log, port, storage, cfg)

	return &App{
		Api: apiSrv,
	}
}
