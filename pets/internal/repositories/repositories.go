package repositories

import (
	"cyberpets/pets/internal/repositories/farm"
	"cyberpets/pets/internal/repositories/incubator"
	"cyberpets/pets/internal/repositories/pet"
	"cyberpets/pets/internal/repositories/user"
	"cyberpets/pets/internal/storage/postgres"
	"go.uber.org/zap"
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
