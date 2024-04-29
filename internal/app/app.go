package app

import (
	"go.uber.org/zap"
	"pets/internal/app/api"
	"pets/internal/config"
	"pets/internal/handlers"
	"pets/internal/repositories"
	"pets/internal/routers"
	"pets/internal/services"
	"pets/internal/storage/postgres"
)

type App struct {
	Api *api.Api
}

func New(log *zap.Logger, port int, storageCfg config.Storage) *App {
	storage := postgres.New(storageCfg)
	repos := repositories.New(storage)
	servs := services.New(repos)
	hands := handlers.New(servs)
	router := routers.New(hands)
	apiSrv := api.New(log, port, router)

	return &App{
		Api: apiSrv,
	}
}
