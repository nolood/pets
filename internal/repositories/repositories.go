package repositories

import (
	"go.uber.org/zap"
	"pets/internal/repositories/farm"
	"pets/internal/repositories/incubator"
	"pets/internal/repositories/pet"
	"pets/internal/repositories/user"
	"pets/internal/storage/postgres"
)

type Repositories struct {
	Pet       pet.Repository
	User      user.Repository
	Farm      farm.Repository
	Incubator incubator.Repository
}

func New(log *zap.Logger, storage *postgres.Storage) *Repositories {
	return &Repositories{
		Pet:       pet.New(storage),
		User:      user.New(log, storage),
		Farm:      farm.New(storage),
		Incubator: incubator.New(storage),
	}
}
