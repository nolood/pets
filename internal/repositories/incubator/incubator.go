package incubator

import (
	"context"
	"database/sql"
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

	stmt, err := r.db.Prepare("SELECT id, user_id, egg_id FROM incubators WHERE user_id = $1")
	if err != nil {
		return models.Incubator{}, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, userId)

	var incubator models.Incubator

	err = row.Scan(&incubator.ID, &incubator.UserID, &incubator.Egg.ID)
	if err != nil {
		return models.Incubator{}, fmt.Errorf("%s: %w", op, err)
	}

	// TODO: get useregg with egg
	//stmt, err = r.db.Prepare("SELECT id, user_id, name, description, image FROM user_eggs WHERE id = $1")

	return models.Incubator{}, nil
}

func (r *incubatorRepo) Init(ctx context.Context) (models.Incubator, error) {
	const op = "repository.incubator.init"

	userId := ctx.Value("user_id")

	egg, err := r.createDefaultEgg(ctx)
	if err != nil {
		return models.Incubator{}, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := r.db.Prepare("INSERT INTO incubators (user_id, egg_id) VALUES ($1) RETURNING id, user_id")
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

	stmt, err := r.db.Prepare("INSERT INTO userseggs (user_id, hatch_time, hatch_start, hatch_end) VALUES ($1, $2, $3, $4) RETURNING id, user_id, hatch_time, hatch_start, hatch_end")
	if err != nil {
		return models.UserEgg{}, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, userId, eggID, 0, nil, nil)

	var egg models.UserEgg

	err = row.Scan(&egg.ID, &egg.UserID, &egg.HatchTime, &egg.HatchStart, &egg.HatchEnd)
	if err != nil {
		return models.UserEgg{}, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err = r.db.Prepare("SELECT id, rarity, image FROM eggs")
	if err != nil {
		return models.UserEgg{}, fmt.Errorf("%s: %w", op, err)
	}

	row = stmt.QueryRowContext(ctx)

	err = row.Scan(&egg.Egg.ID, &egg.Egg.Rarity, &egg.Egg.Image)

	return egg, nil
}
