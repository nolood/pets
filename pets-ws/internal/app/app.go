package app

import (
	"cyberpets/pets-ws/internal/app/ws"
	"cyberpets/pets-ws/internal/app/ws/handlers"
	"go.uber.org/zap"
)

type App struct {
	Ws *ws.Ws
}

func New(log *zap.Logger, port int) *App {
	hands := handlers.New()
	wsSrv := ws.New(log, port, hands)

	return &App{
		Ws: wsSrv,
	}
}
