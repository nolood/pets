package main

import (
	"cyberpets/logger"
	"cyberpets/sso/internal/config"
)

func main() {

	cfg := config.MustLoad()

	log := logger.New(cfg.Env)

	log.Info("sso")
	log.Error("error")
	log.Debug("debug")
	log.Warn("warn")
	log.Fatal("fatal")
	log.Panic("panic")

}
