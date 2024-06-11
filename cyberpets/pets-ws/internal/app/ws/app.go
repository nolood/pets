package ws

import (
	"cyberpets/pets-ws/internal/app/ws/handlers"
	"errors"
	"net/http"
	"strconv"

	"go.uber.org/zap"
	"golang.org/x/net/websocket"
)

type Ws struct {
	log   *zap.Logger
	srv   *http.Server
	port  int
	hands *handlers.Handlers
}

func New(log *zap.Logger, port int, hands *handlers.Handlers) *Ws {
	return &Ws{
		log:   log,
		port:  port,
		hands: hands,
	}
}

func (w *Ws) MustRun() {
	srv := &http.Server{Addr: ":" + strconv.Itoa(w.port), Handler: websocket.Handler(w.hands.Server.Handler)}

	w.srv = srv
	w.log.Info("\nServer started on port - " + strconv.Itoa(w.port))
	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		w.log.Fatal("ListenAndServe(): ", zap.Error(err))
	}
}

func (w *Ws) Stop() {
	err := w.srv.Close()
	if err != nil {
		w.log.Fatal("Server Shutdown:", zap.Error(err))
	}
}
