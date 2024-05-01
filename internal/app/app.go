package app

import (
	"go.uber.org/zap"
	"pets/internal/app/api"
	"pets/internal/config"
	"pets/internal/storage/postgres"
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
