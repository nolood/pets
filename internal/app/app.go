package app

import (
	"go.uber.org/zap"
	"pets/internal/app/api"
	"pets/internal/config"
	"pets/internal/routers"
	"pets/internal/storage/postgres"
)

type App struct {
	Api *api.Api
}

func New(log *zap.Logger, port int, storageCfg config.Storage) *App {
	storage := postgres.New(storageCfg)
	apiRts := routers.New(storage)
	apiSrv := api.New(log, port, apiRts)

	return &App{
		Api: apiSrv,
	}
}
