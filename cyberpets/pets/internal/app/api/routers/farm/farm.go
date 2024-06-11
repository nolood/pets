package farm

import (
	"cyberpets/pets/internal/app/api/handlers/farm"

	"github.com/go-chi/chi/v5"
)

func New(handlers farm.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", handlers.Get)
	r.Put("/slot/{slotId}/pet/{petId}", handlers.SetPet)

	return r
}
