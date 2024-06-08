package incubator

import (
	"cyberpets/pets/internal/app/api/handlers/incubator"
	"github.com/go-chi/chi/v5"
)

func New(handler incubator.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", handler.Get)
	r.Delete("/clear", handler.Clear)
	r.Post("/set/{eggId}", handler.SetEgg)
	r.Post("/open/{eggId}", handler.OpenEgg)

	return r
}
