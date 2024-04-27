package farm

import (
	"log"
	"net/http"
	"pets/internal/storage/postgres"
	"pets/internal/storage/postgres/farm"
)

type Service struct {
	Repo *farm.Repository
}

func New(storage *postgres.Storage) *Service {
	farmRepo := farm.New(storage.Db)
	return &Service{Repo: farmRepo}
}

func (s *Service) Get(w http.ResponseWriter, r *http.Request) {
	log.Print("FARM GET")

	w.WriteHeader(http.StatusOK)
}
