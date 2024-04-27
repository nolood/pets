package farm

import (
	"github.com/go-chi/chi/v5"
	"pets/internal/services/farm"
	"pets/internal/storage/postgres"
)

func New(storage *postgres.Storage) *chi.Mux {
	r := chi.NewRouter()

	farmService := farm.New(storage)

	r.Get("/", farmService.Get)

	return r
}
