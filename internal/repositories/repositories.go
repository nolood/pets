package repositories

import (
	"pets/internal/repositories/user"
	"pets/internal/storage/postgres"
)

type Repositories struct {
	User user.Repository
}

func New(storage *postgres.Storage) *Repositories {
	return &Repositories{
		User: user.New(storage),
	}
}
