package api

import (
	"errors"
	"go.uber.org/zap"
	"net/http"
	"pets/internal/routers"
	"strconv"
)

type Api struct {
	log  *zap.Logger
	srv  *http.Server
	rts  *routers.Routers
	port int
}

func New(log *zap.Logger, port int, rts *routers.Routers) *Api {
	return &Api{
		log:  log,
		port: port,
		rts:  rts,
	}
}

func (a *Api) MustRun() {
	router := a.rts.Router

	srv := &http.Server{Addr: ":" + strconv.Itoa(a.port), Handler: router}

	a.srv = srv

	a.log.Info("\nServer started on port - " + strconv.Itoa(a.port))
	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		a.log.Fatal("ListenAndServe(): ", zap.Error(err))
	}
}

func (a *Api) Stop() {
	err := a.srv.Close()
	if err != nil {
		a.log.Fatal("Server Shutdown:", zap.Error(err))
	}
}
