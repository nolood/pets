package repositories

import (
	"go.uber.org/zap"
	"pets/internal/repositories/farm"
	"pets/internal/repositories/incubator"
	"pets/internal/repositories/user"
	"pets/internal/storage/postgres"
)

type Repositories struct {
	User      user.Repository
	Farm      farm.Repository
	Incubator incubator.Repository
}

func New(log *zap.Logger, storage *postgres.Storage) *Repositories {
	return &Repositories{
		User:      user.New(log, storage),
		Farm:      farm.New(storage),
		Incubator: incubator.New(storage),
	}
}
