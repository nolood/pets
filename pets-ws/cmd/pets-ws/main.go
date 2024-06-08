package main

import (
	"cyberpets/logger"
	"cyberpets/pets-ws/internal/app"
	"cyberpets/pets-ws/internal/config"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	log := logger.New(cfg.Env)

	log.Info("starting pets-ws")

	application := app.New(log, cfg.Port)

	go application.Ws.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	application.Ws.Stop()

	log.Info("Application stopped")

	//server := NewServer()
	//http.Handle("/", websocket.Handler(server.handleWs))
	//err := http.ListenAndServe(":5001", nil)
	//if err != nil {
	//	panic(err)
	//}
}
