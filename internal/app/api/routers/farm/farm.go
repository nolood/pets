package farm

import (
	"github.com/go-chi/chi/v5"
	"pets/internal/app/api/handlers/farm"
)

func New(handlers farm.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", handlers.Get)
	r.Put("/slot/{slotId}/pet/{petId}", handlers.SetPet)

	return r
}
