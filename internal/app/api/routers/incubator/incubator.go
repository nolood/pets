package incubator

import (
	"github.com/go-chi/chi/v5"
	"pets/internal/app/api/handlers/incubator"
)

func New(handler incubator.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", handler.Get)
	r.Post("/set", handler.SetEgg)
	r.Post("/open/{eggId}", handler.OpenEgg)
	r.Delete("/remove", handler.RemoveEgg)

	return r
}
