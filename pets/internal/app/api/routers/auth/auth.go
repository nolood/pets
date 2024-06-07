package auth

import (
	"github.com/go-chi/chi/v5"
	"pets/internal/app/api/handlers/auth"
)

func New(handler auth.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/{user}", handler.Validate)

	return r
}
