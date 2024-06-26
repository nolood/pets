package app

import (
	"cyberpets/pets-ws/internal/app/clicker"
	"cyberpets/pets-ws/internal/app/ws"
	"cyberpets/pets-ws/internal/app/ws/handlers"
	"cyberpets/pets-ws/internal/config"
	"cyberpets/pets-ws/internal/services"

	"go.uber.org/zap"
)

type App struct {
	Ws *ws.Ws
}

func New(log *zap.Logger, cfg *config.Config) *App {
	clicker := clicker.New(log, cfg)

	service := services.New(log, clicker)
	hands := handlers.New(log, service, cfg)
	wsSrv := ws.New(log, cfg.Port, hands)

	return &App{
		Ws: wsSrv,
	}
}
