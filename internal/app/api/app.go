package api

import (
	"errors"
	"go.uber.org/zap"
	"net/http"
	"pets/internal/app/api/handlers"
	"pets/internal/app/api/repositories"
	"pets/internal/app/api/routers"
	"pets/internal/app/api/services"
	"pets/internal/storage/postgres"
	"strconv"
)

type Api struct {
	log    *zap.Logger
	srv    *http.Server
	router *routers.Routers
	port   int
}

func New(log *zap.Logger, port int, storage *postgres.Storage) *Api {
	repos := repositories.New(log, storage)
	servs := services.New(log, repos)
	hands := handlers.New(log, servs)
	router := routers.New(hands)
	return &Api{
		log:    log,
		port:   port,
		router: router,
	}
}

func (a *Api) MustRun() {
	router := a.router.Router

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
