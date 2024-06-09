package handlers

import (
	websockethandler "cyberpets/pets-ws/internal/app/ws/handlers/websocket"
	"cyberpets/pets-ws/internal/config"
	"cyberpets/pets-ws/internal/services"

	"go.uber.org/zap"
)

type Handlers struct {
	Server *websockethandler.Server
}

func New(log *zap.Logger, service *services.Services, cfg *config.Config) *Handlers {
	server := websockethandler.New(log, service.Router, cfg)

	return &Handlers{
		Server: server,
	}
}
