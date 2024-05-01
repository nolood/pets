package main

import (
	"os"
	"os/signal"
	"pets/internal/app"
	"pets/internal/config"
	"pets/internal/lib/logger"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	log := logger.New("local")

	application := app.New(log, cfg.Port, cfg)

	go application.Api.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	application.Api.Stop()

	log.Info("Application stopped")
}
