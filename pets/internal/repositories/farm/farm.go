package farm

import (
	"context"
	"cyberpets/pets/internal/domain/models"
	"cyberpets/pets/internal/storage/postgres"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Init(ctx context.Context) (models.Farm, error)
	Get(ctx context.Context) (models.Farm, error)
	SetPet(ctx context.Context, petID uint64, slotID uint64) (models.Farm, error)
	RemovePet(ctx context.Context) (models.Farm, error)
	BuySlot(ctx context.Context) (models.Farm, error)
}

type farmRepo struct {
	db *sqlx.DB
}

func New(storage *postgres.Storage) Repository {
	return &farmRepo{db: storage.Db}
}

func (r *farmRepo) SetPet(ctx context.Context, petID uint64, slotID uint64) (models.Farm, error) {
	const op = "repository.farm.setPet"

	userId := ctx.Value("user_id")

	stmt, err := r.db.Prepare("WITH user_farm AS (SELECT * FROM farms WHERE user_id = $1), user_pet AS (SELECT * FROM userspets WHERE id = $2 AND user_id = $1) UPDATE slots SET pet_id = user_pet.id FROM user_farm, user_pet WHERE slots.farm_id = user_farm.id AND slots.id = $3 AND slots.is_available = true;")
	if err != nil {
		return models.Farm{}, fmt.Errorf("%s: %w", op, err)
	}

	stmt.QueryRowContext(ctx, userId, petID, slotID)

	farm, err := r.Get(ctx)
	if err != nil {
		return models.Farm{}, fmt.Errorf("%s: %w", op, err)
	}

	return farm, nil
}

func (r *farmRepo) RemovePet(ctx context.Context) (models.Farm, error) {
	//TODO implement me
	panic("implement me")
}

func (r *farmRepo) BuySlot(ctx context.Context) (models.Farm, error) {
	return models.Farm{}, errors.New("not implemented")
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

	stmt, err := r.db.Prepare("SELECT s.id, s.farm_id, s.charge, s.index, s.is_available, s.price, up.id, up.pet_id, up.user_id, up.level, p.id, p.image, p.rarity FROM slots s LEFT JOIN userspets up ON s.pet_id = up.id LEFT JOIN pets p ON up.pet_id = p.id WHERE s.farm_id = $1;")
	if err != nil {
		return slots, fmt.Errorf("%s: %w", op, err)
	}

	rows, err := stmt.QueryContext(ctx, farmID)
	if err != nil {
		return slots, fmt.Errorf("%s: %w", op, err)
	}

	for rows.Next() {
		var slot models.Slot
		var userPet models.UserPet
		var pet models.Pet

		err = rows.Scan(&slot.ID, &slot.FarmID, &slot.Charge, &slot.Index, &slot.IsAvailable, &slot.Price, &userPet.ID, &userPet.PetID, &userPet.UserID, &userPet.Level, &pet.ID, &pet.Image, &pet.Rarity)
		if err != nil {
			return slots, fmt.Errorf("%s: %w", op, err)
		}

		userPet.Pet = &pet

		if userPet.ID != nil {
			slot.Pet = &userPet
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
		stmt, err := r.db.Prepare("INSERT INTO slots (farm_id, index, is_available, price) VALUES ($1, $2, $3, $4) RETURNING id, farm_id, index, charge, is_available, price")
		if err != nil {
			return slots, fmt.Errorf("%s: %w", op, err)
		}

		isAvailable := false
		price := 5000 * (i + 1)

		if i == 0 {
			isAvailable = true
		}

		row := stmt.QueryRowContext(ctx, farmID, i, isAvailable, price)

		var slot models.Slot

		err = row.Scan(&slot.ID, &slot.FarmID, &slot.Index, &slot.Charge, &slot.IsAvailable, &slot.Price)
		if err != nil {
			return slots, fmt.Errorf("%s: %w", op, err)
		}

		slots = append(slots, slot)

	}
	return slots, nil
}
