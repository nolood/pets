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
	RemoveEgg(ctx context.Context, eggID uint64) (models.Egg, error)
}

type incubatorRepo struct {
	db *sql.DB
}

func New(storage *postgres.Storage) Repository {
	return &incubatorRepo{db: storage.Db}
}

func (r *incubatorRepo) RemoveEgg(ctx context.Context, eggID uint64) (models.Egg, error) {
	const op = "repository.incubator.removeEgg"

	userId := ctx.Value("user_id")

	stmt, err := r.db.Prepare("WITH deleted_rows AS (DELETE FROM userseggs WHERE user_id = $1 AND egg_id = $2 RETURNING *), egg_data AS (SELECT id, rarity, image FROM eggs WHERE id = $2) UPDATE incubators SET egg_id = NULL FROM deleted_rows, egg_data WHERE incubators.egg_id = deleted_rows.id RETURNING egg_data.id AS egg_id, egg_data.image AS egg_image, egg_data.rarity AS egg_rarity")
	if err != nil {

		return models.Egg{}, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, userId, eggID)

	var egg models.Egg

	err = row.Scan(&egg.ID, &egg.Image, &egg.Rarity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Egg{}, fmt.Errorf("egg-not-found")
		}

		return models.Egg{}, fmt.Errorf("%s: %w", op, err)
	}

	return egg, nil
}

func (r *incubatorRepo) Get(ctx context.Context) (models.Incubator, error) {
	const op = "repository.incubator.get"

	userId := ctx.Value("user_id")

	stmt, err := r.db.Prepare("SELECT incubators.id AS incubator_id, incubators.user_id AS user_id, userseggs.id AS useregg_id, userseggs.hatch_time AS hatch_time, userseggs.hatch_start AS hatch_start, userseggs.hatch_end AS hatch_end, eggs.id AS egg_id, eggs.rarity AS egg_rarity, eggs.image AS egg_image FROM incubators LEFT JOIN UsersEggs AS userseggs ON incubators.egg_id = userseggs.id LEFT JOIN eggs ON userseggs.egg_id = eggs.id WHERE incubators.user_id = $1;")
	if err != nil {
		return models.Incubator{}, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, userId)

	var incubator models.Incubator
	var userEgg models.UserEgg
	var egg models.Egg

	err = row.Scan(&incubator.ID, &incubator.UserID, &userEgg.ID, &userEgg.HatchTime, &userEgg.HatchStart, &userEgg.HatchEnd, &egg.ID, &egg.Rarity, &egg.Image)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Incubator{}, nil
		}

		return models.Incubator{}, fmt.Errorf("%s: %w", op, err)
	}

	if incubator.ID != 0 && userEgg.ID != nil {
		userEgg.Egg = &egg
		incubator.Egg = &userEgg
	}

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
