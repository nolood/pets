package app

import "cyberpets/pets-ws/internal/app/ws"

type App struct {
	ws *ws.Ws
}

func New() *App {
	return &App{}
}
