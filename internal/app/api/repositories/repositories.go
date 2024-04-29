package repositories

import (
	"go.uber.org/zap"
	"pets/internal/app/api/repositories/user"
	"pets/internal/storage/postgres"
)

type Repositories struct {
	User user.Repository
}

func New(log *zap.Logger, storage *postgres.Storage) *Repositories {
	return &Repositories{
		User: user.New(log, storage),
	}
}
