package user

import (
	"context"
	"database/sql"
	"go.uber.org/zap"
	"pets/internal/domain/models"
	"pets/internal/storage/postgres"
)

type Repository interface {
	Create(ctx context.Context, entity models.User) error
}

type farmRepo struct {
	db  *sql.DB
	log *zap.Logger
}

func New(log *zap.Logger, storage *postgres.Storage) Repository {
	return &farmRepo{db: storage.Db, log: log}
}

func (r *farmRepo) Create(ctx context.Context, entity models.User) error {
	r.log.Info("REPOSITORY: create user", zap.Any("entity", entity))
	return nil
}
