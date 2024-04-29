package farm

import (
	"github.com/go-chi/chi/v5"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	//r.Get("/", farmService.Get)

	return r
}
