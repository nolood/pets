package auth

import (
	"cyberpets/pets/internal/app/api/handlers/auth"

	"github.com/go-chi/chi/v5"
)

func New(handler auth.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/{user}", handler.Validate)

	return r
}
