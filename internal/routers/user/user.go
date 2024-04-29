package user

import (
	"github.com/go-chi/chi/v5"
	"pets/internal/handlers/user"
)

func New(handler *user.Handler) *chi.Mux {
	r := chi.NewRouter()

	return r
}
