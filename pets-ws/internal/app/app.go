package app

import (
	"cyberpets/pets-ws/internal/app/ws"
	"cyberpets/pets-ws/internal/app/ws/handlers"
	"cyberpets/pets-ws/internal/services"
	"go.uber.org/zap"
)

type App struct {
	Ws *ws.Ws
}

func New(log *zap.Logger, port int) *App {
	service := services.New(log)
	hands := handlers.New(service.Router)
	wsSrv := ws.New(log, port, hands)

	return &App{
		Ws: wsSrv,
	}
}
