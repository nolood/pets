package farm

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"pets/internal/domain/models"
	"pets/internal/storage/postgres"
)

type Repository interface {
	Init(ctx context.Context) (models.Farm, error)
	Get(ctx context.Context) (models.Farm, error)
}

type farmRepo struct {
	db *sql.DB
}

func New(storage *postgres.Storage) Repository {
	return &farmRepo{db: storage.Db}
}

func (r *farmRepo) Init(ctx context.Context) (models.Farm, error) {
	const op = "repository.farm.init"

	userId := ctx.Value("user_id")

	stmt, err := r.db.Prepare("INSERT INTO farms (user_id) VALUES ($1) RETURNING id, user_id")
	if err != nil {
		return models.Farm{}, fmt.Errorf("%s: %w", op, err)
	}

	var farm models.Farm

	row := stmt.QueryRowContext(ctx, userId)

	err = row.Scan(&farm.ID, &farm.UserID)
	if err != nil {
		return models.Farm{}, fmt.Errorf("%s: %w", op, err)
	}

	slots, err := r.createDefaultSlots(ctx, farm.ID)
	if err != nil {
		return models.Farm{}, fmt.Errorf("%s: %w", op, err)
	}

	farm.Slots = slots

	return farm, nil
}

func (r *farmRepo) Get(ctx context.Context) (models.Farm, error) {
	const op = "repository.farm.get"

	userId := ctx.Value("user_id")

	var farmInst models.Farm

	stmt, err := r.db.Prepare("SELECT id, user_id FROM farms WHERE user_id = $1")
	if err != nil {
		return farmInst, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, userId)

	err = row.Scan(&farmInst.ID, &farmInst.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return farmInst, nil
		}

		return farmInst, fmt.Errorf("%s: %w", op, err)
	}

	slots, err := r.getFarmSlots(ctx, farmInst.ID)
	if err != nil {
		return farmInst, fmt.Errorf("%s: %w", op, err)
	}

	farmInst.Slots = slots

	return farmInst, nil

}

func (r *farmRepo) getFarmSlots(ctx context.Context, farmID uint64) ([]models.Slot, error) {
	const op = "repository.farm.getSlotsByFarm"

	slots := make([]models.Slot, 0, 6)

	stmt, err := r.db.Prepare("SELECT id, farm_id, pet_id, charge, index FROM slots WHERE farm_id = $1")
	if err != nil {
		return slots, fmt.Errorf("%s: %w", op, err)
	}

	rows, err := stmt.QueryContext(ctx, farmID)
	if err != nil {
		return slots, fmt.Errorf("%s: %w", op, err)
	}

	for rows.Next() {
		var slot models.Slot
		err = rows.Scan(&slot.ID, &slot.FarmID, &slot.PetID, &slot.Charge, &slot.Index)
		if err != nil {
			return slots, fmt.Errorf("%s: %w", op, err)
		}
		slots = append(slots, slot)
	}
	if err = rows.Err(); err != nil {
		return slots, fmt.Errorf("%s: %w", op, err)
	}

	return slots, nil
}

// TODO: подумать могут ли быть ошибки во время создания слотов и создадутся не все слоты

func (r *farmRepo) createDefaultSlots(ctx context.Context, farmID uint64) ([]models.Slot, error) {
	const op = "repository.farm.createDefaultSlots"

	slots := make([]models.Slot, 0, 6)

	for i := 0; i < 6; i++ {
		stmt, err := r.db.Prepare("INSERT INTO slots (farm_id, index) VALUES ($1, $2) RETURNING id, farm_id, index, charge")
		if err != nil {
			return slots, fmt.Errorf("%s: %w", op, err)
		}
		row := stmt.QueryRowContext(ctx, farmID, i)

		var slot models.Slot

		err = row.Scan(&slot.ID, &slot.FarmID, &slot.Index, &slot.Charge)
		if err != nil {
			return slots, fmt.Errorf("%s: %w", op, err)
		}

		slots = append(slots, slot)

	}
	return slots, nil
}
