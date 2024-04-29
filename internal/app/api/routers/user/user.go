package user

import (
	"github.com/go-chi/chi/v5"
	"pets/internal/app/api/handlers/user"
)

func New(handler user.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", handler.Create)

	return r
}
