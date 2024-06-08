package pet

import (
	"context"
	"cyberpets/pets/internal/domain/models"
	"cyberpets/pets/internal/storage/postgres"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type Repository interface {
	AddRandomPetToUser(ctx context.Context, rarity *uint) (models.Pet, error)
}

type petRepo struct {
	db *sqlx.DB
}

func New(storage *postgres.Storage) Repository {
	return &petRepo{db: storage.Db}
}

func (r *petRepo) AddRandomPetToUser(ctx context.Context, rarity *uint) (models.Pet, error) {
	const op = "repository.pet.addRandomPet"

	userId := ctx.Value("user_id")
	log.Println("rarity", rarity, "userId", userId)

	stmt, err := r.db.Prepare("WITH random_pet AS (   SELECT id, image, rarity FROM pets WHERE rarity = $1 ORDER BY RANDOM() LIMIT 1) INSERT INTO userspets (user_id, pet_id) SELECT $2, id FROM random_pet RETURNING id, (SELECT rarity FROM random_pet), (SELECT image FROM random_pet)")
	if err != nil {
		return models.Pet{}, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, rarity, userId)

	var pet models.Pet

	err = row.Scan(&pet.ID, &pet.Rarity, &pet.Image)
	if err != nil {
		return models.Pet{}, fmt.Errorf("%s: %w", op, err)
	}

	return pet, nil
}
