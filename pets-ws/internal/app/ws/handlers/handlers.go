package handlers

import "cyberpets/pets-ws/internal/app/ws/handlers/websocket"

type Handlers struct {
	Server *websockethandler.Server
}

func New() *Handlers {
	server := websockethandler.New()

	return &Handlers{
		Server: server,
	}
}
