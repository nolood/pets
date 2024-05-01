package user

import (
	"context"
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	"pets/internal/domain/models"
	"pets/internal/storage/postgres"
)

type Repository interface {
	CreateOrUpdate(ctx context.Context, entity models.User) (models.User, error)
}

type farmRepo struct {
	db  *sql.DB
	log *zap.Logger
}

func New(log *zap.Logger, storage *postgres.Storage) Repository {
	return &farmRepo{db: storage.Db, log: log}
}

func (r *farmRepo) CreateOrUpdate(ctx context.Context, entity models.User) (models.User, error) {
	const op = "repo.user.create"

	stmt, err := r.db.Prepare(
		"INSERT INTO users (tg_id, username, lastname, firstname, language_code, is_premium, photo_url) VALUES ($1, $2, $3, $4, $5, $6, $7) ON CONFLICT (tg_id) DO UPDATE SET username = excluded.username, lastname = excluded.lastname, firstname = excluded.firstname, language_code = excluded.language_code, is_premium = excluded.is_premium, photo_url = excluded.photo_url RETURNING id, tg_id, username, lastname, firstname, language_code, is_premium, photo_url;")
	if err != nil {
		return entity, fmt.Errorf("%s: %w", op, err)
	}

	var user models.User

	row := stmt.QueryRowContext(ctx, entity.TgID, entity.Username, entity.LastName, entity.FirstName, entity.LanguageCode, entity.IsPremium, entity.PhotoUrl)

	err = row.Scan(&user.ID, &user.TgID, &user.Username, &user.LastName, &user.FirstName, &user.LanguageCode, &user.IsPremium, &user.PhotoUrl)
	if err != nil {
		return entity, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}
