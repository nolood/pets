package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"pets/internal/handlers"
	"pets/internal/routers/user"
	"time"
)

type Routers struct {
	Router *chi.Mux
}

func New(hands *handlers.Handlers) *Routers {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))

	userRts := user.New(hands.User)

	r.Mount("/user", userRts)

	return &Routers{Router: r}
}
