package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"pets/internal/routers/farm"
	"pets/internal/storage/postgres"
	"time"
)

type Routers struct {
	Router *chi.Mux
}

func New(storage *postgres.Storage) *Routers {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))

	farmRts := farm.New(storage)

	r.Mount("/farm", farmRts)

	return &Routers{Router: r}
}
