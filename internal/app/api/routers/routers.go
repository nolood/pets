package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"pets/internal/app/api/handlers"
	"pets/internal/app/api/routers/auth"
	"pets/internal/app/api/routers/user"
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
	authRts := auth.New(hands.Auth)

	r.Mount("/user", userRts)
	r.Mount("/auth", authRts)

	return &Routers{Router: r}
}
