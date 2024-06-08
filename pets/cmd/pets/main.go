package main

import (
	"cyberpets/logger"
	"cyberpets/pets/internal/app"
	"cyberpets/pets/internal/config"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	log := logger.New(cfg.Env)

	application := app.New(log, cfg.Port, cfg)

	go application.Api.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	application.Api.Stop()

	log.Info("Application stopped")
}
