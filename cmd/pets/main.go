package main

import (
	"go.uber.org/zap"
	"pets/internal/config"
	"pets/internal/lib/logger"
)

func main() {

	cfg := config.MustLoad()

	log := logger.New("local")

	log.Info("config", zap.Any("config", cfg))

	// TODO: ...
}
