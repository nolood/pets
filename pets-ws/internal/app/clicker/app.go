package clicker

import (
	"cyberpets/pets-ws/internal/app/clicker/handlers"
	"cyberpets/pets-ws/internal/config"
	"cyberpets/pets-ws/internal/domain/models"

	"go.uber.org/zap"
)

type App struct {
	Mode  int
	Hands handlers.ClickerHandlers
}

func New(log *zap.Logger, cfg *config.Config) *App {

	hands := handlers.New(log)

	return &App{
		Mode:  models.ClickerMode,
		Hands: hands,
	}
}
