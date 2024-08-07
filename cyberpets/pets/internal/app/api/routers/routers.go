package routers

import (
	"cyberpets/pets/internal/app/api/handlers"
	authmiddleware "cyberpets/pets/internal/app/api/middlewares/auth"
	"cyberpets/pets/internal/app/api/routers/auth"
	"cyberpets/pets/internal/app/api/routers/farm"
	"cyberpets/pets/internal/app/api/routers/incubator"
	"cyberpets/pets/internal/app/api/routers/user"
	ssoclient "cyberpets/pets/internal/clients/sso/grpc"
	"cyberpets/pets/internal/config"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Routers struct {
	Router *chi.Mux
}

func New(hands *handlers.Handlers, cfg *config.Config, sso *ssoclient.Client) *Routers {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))

	userRts := user.New(hands.User)
	authRts := auth.New(hands.Auth)
	farmRts := farm.New(hands.Farm)
	incubatorRts := incubator.New(hands.Incubator)

	r.Group(func(r chi.Router) {
		r.Mount("/auth", authRts)
	})

	r.Group(func(r chi.Router) {
		r.Use(authmiddleware.Middleware(cfg.Env, sso))
		r.Mount("/user", userRts)
		r.Mount("/farm", farmRts)
		r.Mount("/incubator", incubatorRts)
	})

	fs := http.StripPrefix("/static", http.FileServer(http.Dir(cfg.Static)))
	r.Handle("/static/*", fs)

	return &Routers{Router: r}
}
