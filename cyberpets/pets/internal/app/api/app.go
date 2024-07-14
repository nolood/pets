package api

import (
	"context"
	"cyberpets/pets/internal/app/api/handlers"
	"cyberpets/pets/internal/app/api/routers"
	"cyberpets/pets/internal/app/api/services"
	ssoclient "cyberpets/pets/internal/clients/sso/grpc"
	"cyberpets/pets/internal/config"
	"cyberpets/pets/internal/repositories"
	"cyberpets/pets/internal/storage/postgres"
	"errors"
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

type Api struct {
	log    *zap.Logger
	srv    *http.Server
	router *routers.Routers
	port   int
}

func New(log *zap.Logger, port int, storage *postgres.Storage, cfg *config.Config) *Api {
	sso := ssoclient.New(context.Background(), cfg.Clients.SSO.Address, cfg.Clients.SSO.Timeout, cfg.Clients.SSO.RetriesCount)
	repos := repositories.New(log, storage)
	servs := services.New(log, repos, cfg, sso)
	hands := handlers.New(log, servs, cfg.Telegram.Token)
	router := routers.New(hands, cfg, sso)

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
