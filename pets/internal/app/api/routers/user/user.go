package user

import (
	"cyberpets/pets/internal/app/api/handlers/user"
	"github.com/go-chi/chi/v5"
)

func New(handler user.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", handler.Create)

	return r
}
