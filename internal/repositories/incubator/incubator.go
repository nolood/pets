package incubator

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"pets/internal/domain/models"
	"pets/internal/storage/postgres"
)

type Repository interface {
	Init(ctx context.Context) (models.Incubator, error)
	Get(ctx context.Context) (models.Incubator, error)
}

type incubatorRepo struct {
	db *sql.DB
}

func New(storage *postgres.Storage) Repository {
	return &incubatorRepo{db: storage.Db}
}

func (r *incubatorRepo) Get(ctx context.Context) (models.Incubator, error) {
	const op = "repository.incubator.get"

	userId := ctx.Value("user_id")

	stmt, err := r.db.Prepare("SELECT inc.id, inc.user_id, ue.id, ue.user_id, ue.hatch_time, ue.hatch_start, ue.hatch_end, e.id, e.rarity, e.image FROM incubators inc JOIN userseggs ue on inc.egg_id = ue.id JOIN eggs e ON ue.egg_id = e.id WHERE inc.user_id = $1")
	if err != nil {
		return models.Incubator{}, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, userId)

	var incubator models.Incubator
	var userEgg models.UserEgg
	var egg models.Egg

	err = row.Scan(&incubator.ID, &incubator.UserID, &userEgg.ID, &userEgg.UserID, &userEgg.HatchTime, &userEgg.HatchStart, &userEgg.HatchEnd, &egg.ID, &egg.Rarity, &egg.Image)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Incubator{}, nil
		}

		return models.Incubator{}, fmt.Errorf("%s: %w", op, err)
	}

	userEgg.Egg = egg
	incubator.Egg = &userEgg

	return incubator, nil
}

func (r *incubatorRepo) Init(ctx context.Context) (models.Incubator, error) {
	const op = "repository.incubator.init"

	userId := ctx.Value("user_id")

	egg, err := r.createDefaultEgg(ctx)
	if err != nil {
		return models.Incubator{}, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := r.db.Prepare("INSERT INTO incubators (user_id, egg_id) VALUES ($1, $2) RETURNING id, user_id")
	if err != nil {
		return models.Incubator{}, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, userId, egg.ID)

	var incubator models.Incubator

	incubator.Egg = &egg

	err = row.Scan(&incubator.ID, &incubator.UserID)
	if err != nil {
		return models.Incubator{}, fmt.Errorf("%s: %w", op, err)
	}

	return incubator, nil
}

func (r *incubatorRepo) createDefaultEgg(ctx context.Context) (models.UserEgg, error) {
	const op = "repository.incubator.createDefaultEgg"

	userId := ctx.Value("user_id")

	eggID := 1

	stmt, err := r.db.Prepare("INSERT INTO userseggs (user_id, egg_id, hatch_time, hatch_start, hatch_end) VALUES ($1, $2, $3, $4, $5) RETURNING id, user_id, hatch_time, hatch_start, hatch_end")
	if err != nil {
		return models.UserEgg{}, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, eggID, userId, 5, nil, nil)

	var egg models.UserEgg

	err = row.Scan(&egg.ID, &egg.UserID, &egg.HatchTime, &egg.HatchStart, &egg.HatchEnd)
	if err != nil {
		return models.UserEgg{}, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err = r.db.Prepare("SELECT id, rarity, image FROM eggs WHERE id = $1")
	if err != nil {
		return models.UserEgg{}, fmt.Errorf("%s: %w", op, err)
	}

	row = stmt.QueryRowContext(ctx, eggID)

	err = row.Scan(&egg.Egg.ID, &egg.Egg.Rarity, &egg.Egg.Image)

	return egg, nil
}
