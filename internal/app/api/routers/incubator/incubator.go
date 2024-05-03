package incubator

import (
	"github.com/go-chi/chi/v5"
	"pets/internal/app/api/handlers/incubator"
)

func New(handler incubator.Handler) *chi.Mux {
	r := chi.NewRouter()

	return r
}
