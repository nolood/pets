package handlers

import (
	"cyberpets/pets-ws/internal/app/ws/handlers/websocket"
	"cyberpets/pets-ws/internal/services/router"
)

type Handlers struct {
	Server *websockethandler.Server
}

func New(service router.Service) *Handlers {
	server := websockethandler.New(service)

	return &Handlers{
		Server: server,
	}
}
