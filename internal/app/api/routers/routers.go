package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"pets/internal/app/api/handlers"
	authmiddleware "pets/internal/app/api/middlewares/auth"
	"pets/internal/app/api/routers/auth"
	"pets/internal/app/api/routers/user"
	"pets/internal/config"
	"time"
)

type Routers struct {
	Router *chi.Mux
}

func New(hands *handlers.Handlers, cfg *config.Config) *Routers {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))

	userRts := user.New(hands.User)
	authRts := auth.New(hands.Auth)

	r.Group(func(r chi.Router) {
		r.Use(authmiddleware.Middleware(cfg.Secret))
		r.Mount("/auth", authRts)
	})

	r.Group(func(r chi.Router) {
		r.Mount("/user", userRts)
	})

	return &Routers{Router: r}
}
