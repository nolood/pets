package postgres

import (
	"cyberpets/pets/internal/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type Storage struct {
	Db *sqlx.DB
}

func New(cfg config.Storage) *Storage {

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s port=%d sslmode=disable", cfg.User, cfg.Dbname, cfg.Password, cfg.Port)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Panic(err)
	}

	return &Storage{Db: db}
}
