package main

import (
	"cyberpets/logger"
	"cyberpets/sso/internal/app"
	"cyberpets/sso/internal/config"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	log := logger.New(cfg.Env)

	application := app.New(log, cfg.GRPC)

	go application.GRPCSrv.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCSrv.Stop()

	log.Info("Application stopped")
}
