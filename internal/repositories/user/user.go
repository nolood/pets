package user

import (
	"context"
	"database/sql"
	"pets/internal/repositories/repository"
	"pets/internal/storage/postgres"
)

type Repository interface {
	repository.Repository
}

type farmRepo struct {
	db *sql.DB
}

func New(storage *postgres.Storage) Repository {
	return &farmRepo{db: storage.Db}
}

func (r *farmRepo) Create(ctx context.Context, entity interface{}) error {

	return nil
}

func (r *farmRepo) Delete(ctx context.Context, id int) error {

	return nil
}

func (r *farmRepo) Update(ctx context.Context, id int, entity interface{}) error {

	return nil
}

func (r *farmRepo) Get(ctx context.Context, id int) (interface{}, error) {

	return nil, nil
}
